package customersvc

import (
	"context"
	"fmt"

	"github.com/arfan21/synapsis_id/internal/customer"
	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/arfan21/synapsis_id/pkg/validation"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repo customer.Repository
}

func New(repo customer.Repository) *Service {
	return &Service{repo: repo}
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
