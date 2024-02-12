package customer

import (
	"context"
	"time"

	customerrepo "github.com/arfan21/synapsis_id/internal/customer/repository"
	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Begin(ctx context.Context) (tx pgx.Tx, err error)
	WithTx(tx pgx.Tx) *customerrepo.Repository

	Create(ctx context.Context, data entity.Customer) (err error)
	GetByEmail(ctx context.Context, email string) (data entity.Customer, err error)
}

type RepositoryRedis interface {
	SetRefreshToken(ctx context.Context, token string, expireIn time.Duration, payload entity.CustomerRefreshToken) (err error)
	IsRefreshTokenExist(ctx context.Context, token string) (payload entity.CustomerRefreshToken, err error)
}
