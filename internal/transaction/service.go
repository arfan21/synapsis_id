package transaction

import (
	"context"

	"github.com/arfan21/synapsis_id/internal/model"
)

type Service interface {
	Checkout(ctx context.Context, req model.CreateTransactionRequest) (res model.CreateTransactionResponse, err error)
}
