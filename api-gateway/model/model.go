package model

import (
	"time"

	"github.com/google/uuid"
)

type RegisterUserInput struct {
	Name      string     `json:"name" binding:"required"`
	Role      string     `json:"role" binding:"required"`
	Email     string     `json:"email" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	CreatedBy *uuid.UUID `json:"created_by"`
}

type User struct {
	ID        uuid.UUID
	Name      string
	Role      string
	Email     string
	Password  string
	CreatedAt time.Time
	CreatedBy uuid.UUID
	UpdatedAt time.Time
	UpdatedBy uuid.UUID
}
