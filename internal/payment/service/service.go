package paymentsvc

import (
	"context"
	"fmt"

	"github.com/arfan21/synapsis_id/internal/entity"
	"github.com/arfan21/synapsis_id/internal/model"
	"github.com/arfan21/synapsis_id/internal/payment"
	"github.com/arfan21/synapsis_id/pkg/validation"
	"github.com/jackc/pgx/v5"
)

type Service struct {
	repo payment.Repository
}

func New(repo payment.Repository) *Service {
	return &Service{repo: repo}
}

func (s Service) WithTx(tx pgx.Tx) payment.Service {
	s.repo = s.repo.WithTx(tx)
	return &s
}

func (s Service) GetPaymentMethods(ctx context.Context) (res []model.GetPayemntMethodResponse, err error) {
	results, err := s.repo.GetPaymentMethods(ctx)
	if err != nil {
		err = fmt.Errorf("payment.service.GetPaymentMethods: failed to get payment methods: %w", err)
		return
	}

	if len(results) == 0 {
		res = make([]model.GetPayemntMethodResponse, 0)
		return
	}

	res = make([]model.GetPayemntMethodResponse, len(results))
	for i, v := range results {
		res[i].ID = v.ID
		res[i].Name = v.Name
	}

	return
}

func (s Service) GetPaymentMethodByID(ctx context.Context, id string) (res model.GetPayemntMethodResponse, err error) {
	result, err := s.repo.GetPaymentMethodByID(ctx, id)
	if err != nil {
		err = fmt.Errorf("payment.service.GetPaymentMethodByID: failed to get payment method by id: %w", err)
		return
	}

	res.ID = result.ID
	res.Name = result.Name
	return
}

func (s Service) CreatePayment(ctx context.Context, req model.CreatePaymentRequest) (err error) {
	err = validation.Validate(req)
	if err != nil {
		err = fmt.Errorf("payment.service.CreatePayment: failed to validate request: %w", err)
		return
	}

	data := entity.Payment{
		TransactionID: req.TransactionID,
	}

	err = s.repo.CreatePayment(ctx, data)
	if err != nil {
		err = fmt.Errorf("payment.service.CreatePayment: failed to create payment: %w", err)
		return
	}

	return
}
