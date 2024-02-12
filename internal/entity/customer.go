package entity

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID `json:"id"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Customer) TableName() string {
	return "customers"
}

type CustomerRefreshToken struct {
	ID    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}
