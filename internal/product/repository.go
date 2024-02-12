package product

import (
	"context"

	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/arfan21/synapsis_id/internal/model"
	productrepo "github.com/arfan21/synapsis_id/internal/product/repository"
	"github.com/jackc/pgx/v5"
)

type Repository interface {
	Begin(ctx context.Context) (tx pgx.Tx, err error)
	WithTx(tx pgx.Tx) *productrepo.Repository

	Create(ctx context.Context, data entity.Product) (err error)
	IsCategoryExist(ctx context.Context, id string) (exist bool, err error)
	GetCategories(ctx context.Context) (result []entity.ProductCategory, err error)
	GetProducts(ctx context.Context, filter model.GetListProductRequest) (result []entity.Product, err error)
	GetTotalProduct(ctx context.Context, filter model.GetListProductRequest) (result int, err error)
	IsProductExist(ctx context.Context, id string) (exist bool, err error)
	GetProductByID(ctx context.Context, id string) (result entity.Product, err error)
}
