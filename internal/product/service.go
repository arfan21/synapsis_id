package product

import (
	"context"

	"github.com/arfan21/synapsis_id/internal/model"
)

type Service interface {
	Create(ctx context.Context, req model.ProductCreateRequest) (err error)
	GetCategories(ctx context.Context) (res []model.GetCategoryResponse, err error)
	GetProducts(ctx context.Context, req model.GetListProductRequest) (
		res []model.GetProductResponse,
		total int,
		err error,
	)
}
