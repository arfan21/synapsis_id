package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type ProductCategory struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (ProductCategory) TableName() string {
	return "product_categories"
}

type Product struct {
	ID         uuid.UUID       `json:"id"`
	CustomerID uuid.UUID       `json:"customer_id"`
	CategoryID uuid.UUID       `json:"category_id"`
	Name       string          `json:"name"`
	Stok       int             `json:"stok"`
	Price      decimal.Decimal `json:"price"`
	CreatedAt  time.Time       `json:"created_at"`
	UpdatedAt  time.Time       `json:"updated_at"`
	Category   ProductCategory `json:"category"`
	Customer   Customer        `json:"customer"`
}

func (Product) TableName() string {
	return "products"
}
