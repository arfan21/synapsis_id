package model

import "github.com/google/uuid"

type GetPayemntMethodResponse struct {
	ID   uuid.UUID `json:"id" swaggertype:"string"`
	Name string    `json:"name"`
}
