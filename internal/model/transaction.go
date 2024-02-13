package model

import "github.com/google/uuid"

type CreateTransactionRequest struct {
	CustomerID      uuid.UUID `json:"customer_id" validate:"required"`
	PaymentMethodID uuid.UUID `json:"payment_method_id" validate:"required"`
}

type CreateTransactionResponse struct {
	TransactionID uuid.UUID `json:"transaction_id"`
}
