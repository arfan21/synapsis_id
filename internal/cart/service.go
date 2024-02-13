package cart

import (
	"context"

	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	WithTx(tx pgx.Tx) Service

	Create(ctx context.Context, req model.CreateCartRequest) (err error)
	GetByCustomerID(ctx context.Context, customerID string) (res []model.GetCartResponse, err error)
	Delete(ctx context.Context, customerID, productID string) (err error)
	DeleteAll(ctx context.Context, customerID string) (err error)
}
