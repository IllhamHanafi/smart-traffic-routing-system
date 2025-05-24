package repository

import (
	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/config"
)

type Driver string

const (
	SQLC Driver = "sqlc"
)

type Repository interface {
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
