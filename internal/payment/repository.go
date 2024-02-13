package payment

import (
	"context"

	"github.com/arfan21/synapsis_id/internal/entity"
	paymentrepo "github.com/arfan21/synapsis_id/internal/payment/repository"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Begin(ctx context.Context) (tx pgx.Tx, err error)
	WithTx(tx pgx.Tx) *paymentrepo.Repository

	GetPaymentMethods(ctx context.Context) (result []entity.PaymentMethod, err error)
	GetPaymentMethodByID(ctx context.Context, id string) (result entity.PaymentMethod, err error)
}
