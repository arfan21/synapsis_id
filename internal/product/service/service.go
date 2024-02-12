package productsvc

import (
	"context"
	"fmt"

	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/arfan21/synapsis_id/internal/product"
	"github.com/arfan21/synapsis_id/pkg/constant"
	"github.com/arfan21/synapsis_id/pkg/validation"
	"github.com/google/uuid"
)

type Service struct {
	repo product.Repository
}

func New(repo product.Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) Create(ctx context.Context, req model.ProductCreateRequest) (err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("product.service.Create: failed to validate request : %w", err)
		return
	}

	customerIDUUID, err := uuid.Parse(req.CustomerID)
	if err != nil {
		err = fmt.Errorf("product.service.Create: failed to parse Customer ID : %w", err)
		return
	}

	categoryIDUUID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		err = fmt.Errorf("product.service.Create: failed to parse Category ID : %w", err)
		return
	}

	exist, err := s.repo.IsCategoryExist(ctx, req.CategoryID)
	if err != nil {
		err = fmt.Errorf("product.service.Create: failed to check category  id : %w", err)
		return
	}

	if !exist {
		return constant.ErrCategoryNotFound
	}

	data := entity.Product{
		CustomerID: customerIDUUID,
		CategoryID: categoryIDUUID,
		Name:       req.Name,
		Stok:       req.Stok,
		Price:      req.Price,
	}

	err = s.repo.Create(ctx, data)
	if err != nil {
		err = fmt.Errorf("product.service.Create: failed to create new product : %w", err)
		return
	}

	return nil
}

func (s Service) GetCategories(ctx context.Context) (res []model.GetCategoryResponse, err error) {
	results, err := s.repo.GetCategories(ctx)
	if err != nil {
		err = fmt.Errorf("product.service.GetCategories: failed to get categories : %w", err)
		return
	}

	if len(results) == 0 {
		res = make([]model.GetCategoryResponse, 0)
		return
	}

	res = make([]model.GetCategoryResponse, len(results))

	for i, result := range results {
		res[i].ID = result.ID
		res[i].Name = result.Name
	}

	return
}

func (s Service) GetProducts(ctx context.Context, req model.GetListProductRequest) (
	res []model.GetProductResponse,
	total int,
	err error,
) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("product.service.GetProducts: failed to validate request : %w", err)
		return
	}

	results, err := s.repo.GetProducts(ctx, req)
	if err != nil {
		err = fmt.Errorf("product.service.GetProducts: failed to get products from db : %w", err)
		return
	}

	if len(results) == 0 {
		res = make([]model.GetProductResponse, 0)
		return
	}

	res = make([]model.GetProductResponse, len(results))

	for i, result := range results {
		res[i].ID = result.ID
		res[i].Name = result.Name
		res[i].Stok = result.Stok
		res[i].Price = result.Price
		res[i].CategoryID = result.Category.ID
		res[i].CategoryName = result.Category.Name
		res[i].OwnerID = result.Customer.ID
		res[i].OwnerName = result.Customer.Fullname
	}

	total, err = s.repo.GetTotalProduct(ctx, req)
	if err != nil {
		err = fmt.Errorf("product.service.GetProducts: failed to get total product from db : %w", err)
		return
	}

	return
}
