package payment

import (
	"context"

	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	WithTx(tx pgx.Tx) Service

	GetPaymentMethods(ctx context.Context) (res []model.GetPayemntMethodResponse, err error)
	GetPaymentMethodByID(ctx context.Context, id string) (res model.GetPayemntMethodResponse, err error)
	CreatePayment(ctx context.Context, req model.CreatePaymentRequest) (err error)
}
