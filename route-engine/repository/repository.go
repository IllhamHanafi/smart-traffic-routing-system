package repository

import (
	"context"
	"route-engine/config"
)

type Driver string

const (
	SQLC Driver = "sqlc"
)

type Repository interface {
	GetActiveRoutingLogic(ctx context.Context) (ActiveRoutingLogicResult, error)
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
