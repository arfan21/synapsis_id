package cart

import (
	"context"

	"github.com/arfan21/synapsis_id/internal/model"
)

type Service interface {
	Create(ctx context.Context, req model.CreateCartRequest) (err error)
}
