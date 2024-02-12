package entity

import (
	"time"

	"github.com/google/uuid"
)

type PaymentMethod struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (PaymentMethod) TableName() string {
	return "payment_methods"
}
