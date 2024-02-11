package customer

import (
	"context"

	customerrepo "github.com/arfan21/synapsis_id/internal/customer/repository"
	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Begin(ctx context.Context) (tx pgx.Tx, err error)
	WithTx(tx pgx.Tx) *customerrepo.Repository

	Create(ctx context.Context, data entity.Customer) (err error)
}
