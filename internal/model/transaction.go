package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateTransactionRequest struct {
	CustomerID      uuid.UUID `json:"customer_id" validate:"required"`
	PaymentMethodID uuid.UUID `json:"payment_method_id" validate:"required"`
}

type CreateTransactionResponse struct {
	TransactionID uuid.UUID `json:"transaction_id"`
}

type TransactionPayRequest struct {
	TransactionID uuid.UUID       `json:"transaction_id" validate:"required"`
	TotalAmount   decimal.Decimal `json:"total_amount" validate:"required" swaggertype:"string"`
}
