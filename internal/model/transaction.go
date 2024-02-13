package model

import (
	"time"

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

type GetTransactionResponse struct {
	ID            uuid.UUID                      `json:"id" swaggertype:"string"`
	CustomerID    uuid.UUID                      `json:"customer_id" swaggertype:"string"`
	PaymentMethod string                         `json:"payment_method"`
	TotalAmount   decimal.Decimal                `json:"total_amount" swaggertype:"string"`
	Status        string                         `json:"status"`
	CreatedAt     time.Time                      `json:"created_at"`
	UpdatedAt     time.Time                      `json:"updated_at"`
	Details       []GetTransactionDetailResponse `json:"details"`
}

type GetTransactionDetailResponse struct {
	ID           uuid.UUID       `json:"id" swaggertype:"string"`
	TrasactionID uuid.UUID       `json:"transaction_id" swaggertype:"string"`
	ProductID    uuid.UUID       `json:"product_id" swaggertype:"string"`
	ProductName  string          `json:"product_name"`
	ProductPrice decimal.Decimal `json:"product_price" swaggertype:"string"`
}
