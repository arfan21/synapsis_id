package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type CreateCartRequest struct {
	CustomerID uuid.UUID `json:"customer_id" validate:"required" swaggertype:"string"`
	ProductID  uuid.UUID `json:"product_id" validate:"required" swaggertype:"string"`
	Qty        int       `json:"qty" validate:"required,gte=1"`
}

type GetCartResponse struct {
	ID           uuid.UUID       `json:"id" swaggertype:"string"`
	CustomerID   uuid.UUID       `json:"customer_id" swaggertype:"string"`
	ProductID    uuid.UUID       `json:"product_id" swaggertype:"string"`
	ProductName  string          `json:"product_name"`
	ProductStok  int             `json:"product_stok"`
	ProductPrice decimal.Decimal `json:"product_price" swaggertype:"string"`
	CreatedAt    time.Time       `json:"created_at"`
	Qty          int             `json:"qty"`
}

type DeleteCartRequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required" swaggertype:"string"`
}
