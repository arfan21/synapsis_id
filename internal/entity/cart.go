package entity

import "github.com/google/uuid"

type Cart struct {
	ID         uuid.UUID `json:"id"`
	CustomerID uuid.UUID `json:"customer_id"`
	ProductID  uuid.UUID `json:"product_id"`
	CreatedAt  string    `json:"created_at"`
	UpdatedAt  string    `json:"updated_at"`
	Product    Product   `json:"product"`
}

func (Cart) TableName() string {
	return "carts"
}
