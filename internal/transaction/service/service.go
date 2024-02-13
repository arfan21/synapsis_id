package transactionsvc

import (
	"context"
	"fmt"

	"github.com/arfan21/synapsis_id/internal/cart"
	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/arfan21/synapsis_id/internal/payment"
	"github.com/arfan21/synapsis_id/internal/product"
	"github.com/arfan21/synapsis_id/internal/transaction"
	"github.com/arfan21/synapsis_id/pkg/constant"
	"github.com/arfan21/synapsis_id/pkg/validation"
)

type Service struct {
	repo       transaction.Repository
	cartSvc    cart.Service
	paymentSvc payment.Service
	productSvc product.Service
}

func New(repo transaction.Repository, cartSvc cart.Service, paymentSvc payment.Service, productSvc product.Service) *Service {
	return &Service{repo: repo, cartSvc: cartSvc, paymentSvc: paymentSvc, productSvc: productSvc}
}

func (s Service) Checkout(ctx context.Context, req model.CreateTransactionRequest) (res model.CreateTransactionResponse, err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("transaction.service.Checkout: failed to validate request : %w", err)
		return
	}

	tx, err := s.repo.Begin(ctx)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
			return
		}

		err = tx.Commit(ctx)
		if err != nil {
			err = fmt.Errorf("transaction.service.Checkout: failed to commit transaction : %w", err)
			return
		}
	}()

	// check payment method
	_, err = s.paymentSvc.GetPaymentMethodByID(ctx, req.PaymentMethodID.String())
	if err != nil {
		err = fmt.Errorf("transaction.service.Checkout: failed to get payment method : %w", err)
		return
	}

	// get product from cart
	products, err := s.cartSvc.GetByCustomerID(ctx, req.CustomerID.String())
	if err != nil {
		err = fmt.Errorf("transaction.service.Checkout: failed to get product from cart : %w", err)
		return
	}

	if len(products) == 0 {
		err = fmt.Errorf("transaction.service.Checkout: cart is empty : %w", constant.ErrNoProductInCart)
		return
	}

	data := entity.Transaction{
		CustomerID:      req.CustomerID,
		PaymentMethodID: req.PaymentMethodID,
	}

	dataBatchUpdateStok := make([]model.UpdateStokRequest, len(products))

	// calculate total amount
	for i, product := range products {
		data.TotalAmount = data.TotalAmount.Add(product.ProductPrice)
		dataBatchUpdateStok[i] = model.UpdateStokRequest{
			ID:   product.ProductID,
			Stok: product.ProductStok - product.Qty,
		}

		if dataBatchUpdateStok[i].Stok < 0 {
			err = fmt.Errorf("transaction.service.Checkout: product stok is not enough : %w", constant.ErrProductStokNotEnough)
			return res, err
		}
	}

	id, err := s.repo.WithTx(tx).Create(ctx, data)
	if err != nil {
		return
	}

	dataDetail := make([]entity.TransactionDetail, len(products))

	for i, product := range products {
		dataDetail[i] = entity.TransactionDetail{
			TransactionID: id,
			ProductID:     product.ProductID,
			Qty:           product.Qty,
		}
	}

	err = s.repo.WithTx(tx).CreateDetail(ctx, dataDetail)
	if err != nil {
		err = fmt.Errorf("transaction.service.Checkout: failed to create transaction detail : %w", err)
		return
	}

	err = s.productSvc.WithTx(tx).BatchUpdateStok(ctx, dataBatchUpdateStok)
	if err != nil {
		err = fmt.Errorf("transaction.service.Checkout: failed to update product stok : %w", err)
		return
	}

	err = s.cartSvc.WithTx(tx).DeleteAll(ctx, req.CustomerID.String())
	if err != nil {
		err = fmt.Errorf("transaction.service.Checkout: failed to delete all product from cart : %w", err)
		return
	}

	res.TransactionID = id

	return
}

func (s Service) Pay(ctx context.Context, req model.TransactionPayRequest) (err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("transaction.service.Pay: failed to validate request : %w", err)
		return
	}

	data, err := s.repo.GetByID(ctx, req.TransactionID)
	if err != nil {
		err = fmt.Errorf("transaction.service.Pay: failed to get transaction by id : %w", err)
		return
	}

	if data.Status != entity.TransactionStatusWaitingPayment {
		err = fmt.Errorf("transaction.service.Pay: transaction status is not waiting payment : %w", constant.ErrTransactionAlreadyPaidOrFailed)
		return
	}

	if data.TotalAmount.Cmp(req.TotalAmount) != 0 {
		err = fmt.Errorf("transaction.service.Pay: amount is not enough : %w", constant.ErrPaymentNotEqualTotalAmount)
		return
	}

	tx, err := s.repo.Begin(ctx)
	if err != nil {
		err = fmt.Errorf("transaction.service.Pay: failed to begin transaction : %w", err)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
			return
		}

		err = tx.Commit(ctx)
		if err != nil {
			err = fmt.Errorf("transaction.service.Pay: failed to commit transaction : %w", err)
			return
		}
	}()

	err = s.paymentSvc.WithTx(tx).CreatePayment(ctx, model.CreatePaymentRequest{
		TransactionID: req.TransactionID,
	})
	if err != nil {
		err = fmt.Errorf("transaction.service.Pay: failed to create payment : %w", err)
		return
	}

	err = s.repo.WithTx(tx).UpdateStatus(ctx, req.TransactionID, entity.TransactionStatusProcessing)
	if err != nil {
		err = fmt.Errorf("transaction.service.Pay: failed to update transaction status : %w", err)
		return
	}

	return
}

func (s Service) GetByCustomerID(ctx context.Context, customerID string) (result []model.GetTransactionResponse, err error) {
	data, err := s.repo.GetByCustomerID(ctx, customerID)
	if err != nil {
		err = fmt.Errorf("transaction.service.GetByCustomerID: failed to get transaction by customer id : %w", err)
		return
	}

	if len(data) == 0 {
		result = make([]model.GetTransactionResponse, 0)
		return
	}

	result = make([]model.GetTransactionResponse, len(data))
	for i, v := range data {
		result[i] = model.GetTransactionResponse{
			ID:            v.ID,
			CustomerID:    v.CustomerID,
			Status:        string(v.Status),
			TotalAmount:   v.TotalAmount,
			CreatedAt:     v.CreatedAt,
			PaymentMethod: v.PaymentMethod.Name,
			UpdatedAt:     v.UpdatedAt,
		}

		for _, detail := range v.TransactionDetails {
			result[i].Details = append(result[i].Details, model.GetTransactionDetailResponse{
				ID:           detail.ID,
				TrasactionID: detail.TransactionID,
				ProductID:    detail.ProductID,
				ProductName:  detail.Product.Name,
				ProductPrice: detail.Product.Price,
				Qty:          detail.Qty,
			})
		}
	}

	return
}
