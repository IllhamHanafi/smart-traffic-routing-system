package model

import "github.com/google/uuid"

type Order struct {
	ID uuid.UUID `json:"id"`
}
