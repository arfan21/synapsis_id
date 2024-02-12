package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateCartRequest struct {
	CustomerID uuid.UUID `json:"customer_id" validate:"required" swaggertype:"string"`
	ProductID  uuid.UUID `json:"product_id" validate:"required" swaggertype:"string"`
}

type GetCartResponse struct {
	ID           uuid.UUID       `json:"id" swaggertype:"string"`
	CustomerID   uuid.UUID       `json:"customer_id" swaggertype:"string"`
	ProductID    uuid.UUID       `json:"product_id" swaggertype:"string"`
	ProductName  string          `json:"product_name"`
	ProductStok  int             `json:"product_stok"`
	ProductPrice decimal.Decimal `json:"product_price" swaggertype:"string"`
	CreatedAt    time.Time       `json:"created_at"`
}
