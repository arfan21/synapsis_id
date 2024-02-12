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
