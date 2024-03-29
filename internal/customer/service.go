package customer

import (
	"context"

	"github.com/arfan21/synapsis_id/internal/model"
)

type Service interface {
	Register(ctx context.Context, req model.CustomerRegisterRequest) (err error)
	Login(ctx context.Context, req model.CustomerLoginRequest) (res model.CustomerLoginResponse, err error)
	RefreshToken(ctx context.Context, req model.CustomerRefreshTokenRequest) (res model.CustomerLoginResponse, err error)
	Logout(ctx context.Context, req model.CustomerLogoutRequest) (err error)
}
