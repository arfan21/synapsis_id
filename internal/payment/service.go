package payment

import (
	"context"

	"github.com/arfan21/synapsis_id/internal/model"
)

type Service interface {
	GetPaymentMethods(ctx context.Context) (res []model.GetPayemntMethodResponse, err error)
}
