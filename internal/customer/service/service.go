package customersvc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/arfan21/synapsis_id/config"
	"github.com/arfan21/synapsis_id/internal/customer"
	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/arfan21/synapsis_id/pkg/constant"
	"github.com/arfan21/synapsis_id/pkg/validation"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo      customer.Repository
	repoRedis customer.RepositoryRedis
}

func New(repo customer.Repository, repoRedis customer.RepositoryRedis) *Service {
	return &Service{repo: repo, repoRedis: repoRedis}
}

func (s Service) Register(ctx context.Context, req model.CustomerRegisterRequest) (err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("customer.service.Register: failed to validate request: %w", err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		err = fmt.Errorf("customer.service.Register: failed to hash password: %w", err)
		return
	}

	data := entity.Customer{
		Fullname: req.Fullname,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	err = s.repo.Create(ctx, data)
	if err != nil {
		err = fmt.Errorf("customer.service.Register: failed to register customer: %w", err)
		return
	}

	return
}

func (s Service) Login(ctx context.Context, req model.CustomerLoginRequest) (res model.CustomerLoginResponse, err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("customer.service.Login: failed to validate request: %w", err)
		return
	}

	data, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		err = fmt.Errorf("customer.service.Login: failed to get customer by email: %w", err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(req.Password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			err = constant.ErrEmailOrPasswordInvalid
		}
		err = fmt.Errorf("customer.service.Login: failed to compare password: %w", err)
		return
	}

	accessTokenExpire := time.Duration(config.GetConfig().JWT.AccessTokenExpireIn) * time.Second

	accessToken, err := s.CreateJWTWithExpiry(
		data.ID.String(),
		data.Email,
		config.GetConfig().JWT.AccessTokenSecret,
		accessTokenExpire,
	)

	if err != nil {
		err = fmt.Errorf("customer.service.Login: failed to create access token: %w", err)
		return
	}

	refreshTokenExpire := time.Duration(config.GetConfig().JWT.RefreshTokenExpireIn) * time.Second

	refreshToken, err := s.CreateJWTWithExpiry(
		data.ID.String(),
		data.Email,
		config.GetConfig().JWT.RefreshTokenSecret,
		refreshTokenExpire,
	)

	if err != nil {
		err = fmt.Errorf("customer.service.Login: failed to create refresh token: %w", err)
		return
	}

	err = s.repoRedis.SetRefreshToken(ctx, refreshToken, refreshTokenExpire)
	if err != nil {
		err = fmt.Errorf("customer.service.Login: failed to set refresh token: %w", err)
		return
	}

	res = model.CustomerLoginResponse{
		AccessToken:           accessToken,
		ExpiresIn:             int(accessTokenExpire.Seconds()),
		TokenType:             "Bearer",
		RefreshToken:          refreshToken,
		ExpiresInRefreshToken: int(refreshTokenExpire.Seconds()),
	}

	return
}

func (s Service) CreateJWTWithExpiry(id, email, secret string, expiry time.Duration) (token string, err error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, model.JWTClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "Synapsis ID",
			Subject:   id,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	token, err = jwtToken.SignedString([]byte(secret))
	if err != nil {
		err = fmt.Errorf("usecase: failed to create jwt token: %w", err)
		return
	}

	return
}
