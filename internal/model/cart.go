package model

import "github.com/google/uuid"

type CreateCartRequest struct {
	CustomerID uuid.UUID `json:"customer_id" validate:"required"`
	ProductID  uuid.UUID `json:"product_id" validate:"required"`
}
