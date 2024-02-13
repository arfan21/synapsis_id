package entity

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	ProductID  uuid.UUID `json:"product_id"`
	Qty        int       `json:"qty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Product    Product   `json:"product"`
}

func (Cart) TableName() string {
	return "carts"
}
