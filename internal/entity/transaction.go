package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type TransactionStatus string

const (
	TransactionStatusWaitingPayment TransactionStatus = "WAITING_PAYMENT"
	TransactionStatusProcessing     TransactionStatus = "PROCESSING"
	TransactionStatusCompleted      TransactionStatus = "COMPLETED"
	TransactionStatusFailed         TransactionStatus = "FAILED"
)

type Transaction struct {
	ID                 uuid.UUID           `json:"id"`
	CustomerID         uuid.UUID           `json:"customer_id"`
	PaymentMethodID    uuid.UUID           `json:"payment_method_id"`
	Status             TransactionStatus   `json:"status"`
	TotalAmount        decimal.Decimal     `json:"total_amount"`
	CreatedAt          time.Time           `json:"created_at"`
	UpdatedAt          time.Time           `json:"updated_at"`
	TransactionDetails []TransactionDetail `json:"transaction_details"`
	PaymentMethod      PaymentMethod       `json:"payment_method"`
}

func (Transaction) TableName() string {
	return "transactions"
}

type TransactionDetail struct {
	ID            uuid.UUID `json:"id"`
	TransactionID uuid.UUID `json:"transaction_id"`
	ProductID     uuid.UUID `json:"product_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Product       Product   `json:"product"`
	Qty           int       `json:"qty"`
}

func (TransactionDetail) TableName() string {
	return "transaction_details"
}
