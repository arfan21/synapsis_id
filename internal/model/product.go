package model

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductCreateRequest struct {
	Name       string          `json:"name" validate:"required"`
	Stok       int             `json:"stok" validate:"required"`
	Price      decimal.Decimal `json:"price" validate:"required" swaggertype:"string"`
	CustomerID string          `json:"customer_id" validate:"required"`
	CategoryID string          `json:"category_id" validate:"required"`
}

type GetCategoryResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"string"`
}
