package cart

import (
	"context"

	cartrepo "github.com/arfan21/synapsis_id/internal/cart/repository"
	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Begin(ctx context.Context) (tx pgx.Tx, err error)
	WithTx(tx pgx.Tx) *cartrepo.Repository

	Create(ctx context.Context, cart entity.Cart) (err error)
	GetByCustomerID(ctx context.Context, customerID string) (data []entity.Cart, err error)
	Delete(ctx context.Context, customerID, productID string) (err error)
}
