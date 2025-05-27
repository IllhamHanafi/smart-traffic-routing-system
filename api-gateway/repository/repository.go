package repository

import (
	"context"

	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/model"
	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/config"
	"github.com/google/uuid"
)

type Driver string

const (
	SQLC Driver = "sqlc"
)

type Repository interface {
	CreateUser(ctx context.Context, input CreateUserInput) (uuid.UUID, error)
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	Close()
}

func New(cfg config.Database) Repository {
	switch cfg.Library {
	case string(SQLC):
		return NewSqlcRepository(cfg)
	default:
		return nil
	}
}
