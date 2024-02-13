package product

import (
	"context"

	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/jackc/pgx/v5"
)

type Service interface {
	WithTx(tx pgx.Tx) Service

	Create(ctx context.Context, req model.ProductCreateRequest) (err error)
	GetCategories(ctx context.Context) (res []model.GetCategoryResponse, err error)
	GetProducts(ctx context.Context, req model.GetListProductRequest) (
		res []model.GetProductResponse,
		total int,
		err error,
	)

	IsProductExist(ctx context.Context, id string) (exist bool, err error)
	GetProductByID(ctx context.Context, id string) (res model.GetProductResponse, err error)
	BatchUpdateStok(ctx context.Context, req []model.UpdateStokRequest) (err error)
}
