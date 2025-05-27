package repository

import "github.com/google/uuid"

type CreateUserInput struct {
	Name      string
	Role      string
	Email     string
	Password  string
	CreatedBy *uuid.UUID
	UpdatedBy *uuid.UUID
}
