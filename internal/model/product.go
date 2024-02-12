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
	ID   uuid.UUID `json:"id" swaggertype:"string"`
	Name string    `json:"string"`
}

type GetListProductRequest struct {
	Name          string        `query:"name" json:"name"`
	CategoryID    uuid.NullUUID `query:"category_id" json:"category_id" swaggertype:"string"`
	Page          int           `query:"page" json:"page" validate:"min=1"`
	Limit         int           `query:"limit" json:"limit" validate:"min=1"`
	DisableOffset bool          `json:"-"`
}

type GetProductResponse struct {
	ID           uuid.UUID       `json:"id" swaggertype:"string"`
	Name         string          `json:"name"`
	Stok         int             `json:"stok"`
	Price        decimal.Decimal `json:"price" swaggertype:"string"`
	CategoryID   uuid.UUID       `json:"category_id" swaggertype:"string"`
	CategoryName string          `json:"category_name"`
	OwnerID      uuid.UUID       `json:"owner_id" swaggertype:"string"`
	OwnerName    string          `json:"owner_name"`
}
