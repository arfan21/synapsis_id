package cartsvc

import (
	"context"
	"fmt"

	"github.com/arfan21/synapsis_id/internal/cart"
	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/arfan21/synapsis_id/internal/product"
	"github.com/arfan21/synapsis_id/pkg/constant"
	"github.com/arfan21/synapsis_id/pkg/validation"
)

type Service struct {
	repo       cart.Repository
	productSvc product.Service
}

func New(repo cart.Repository, productSvc product.Service) *Service {
	return &Service{repo: repo, productSvc: productSvc}
}

func (s Service) Create(ctx context.Context, req model.CreateCartRequest) (err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("cart.service.Create: failed to validate request: %w", err)
		return
	}

	dataProduct, err := s.productSvc.GetProductByID(ctx, req.ProductID.String())
	if err != nil {
		err = fmt.Errorf("cart.service.Create: failed to check product id: %w", err)
		return
	}

	if dataProduct.OwnerID == req.CustomerID {
		err = fmt.Errorf("cart.service.Create: product id not found : %w", constant.ErrCannotAddOwnProductToCart)
		return
	}

	data := entity.Cart{
		CustomerID: req.CustomerID,
		ProductID:  req.ProductID,
	}

	err = s.repo.Create(ctx, data)
	if err != nil {
		err = fmt.Errorf("cart.service.Create: failed to create cart to db: %w", err)
		return
	}

	return
}

func (s Service) GetByCustomerID(ctx context.Context, customerID string) (res []model.GetCartResponse, err error) {
	results, err := s.repo.GetByCustomerID(ctx, customerID)
	if err != nil {
		err = fmt.Errorf("cart.service.GetByCustomerID: failed to get cart by customer id: %w", err)
		return
	}

	if len(results) == 0 {
		res = make([]model.GetCartResponse, 0)
		return
	}

	res = make([]model.GetCartResponse, len(results))

	for i, v := range results {
		res[i].ID = v.ID
		res[i].CustomerID = v.CustomerID
		res[i].ProductID = v.ProductID
		res[i].CreatedAt = v.CreatedAt
		res[i].ProductName = v.Product.Name
		res[i].ProductPrice = v.Product.Price
		res[i].ProductStok = v.Product.Stok
	}

	return
}
