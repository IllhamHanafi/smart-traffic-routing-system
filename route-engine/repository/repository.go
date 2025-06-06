package repository

import (
	"context"

	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/model"
	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/config"
)

type Driver string

const (
	SQLC Driver = "sqlc"
)

type Repository interface {
	GetActiveRoutingLogic(ctx context.Context) (ActiveRoutingLogicResult, error)
	GetCourierByCode(ctx context.Context, courierCode string) (model.Courier, error)
	InsertRoutingDecisionLog(ctx context.Context, input InsertRoutingDecisionLogParams) error
	CreateActiveRoutingDecision(ctx context.Context, input InsertNewRoutingDecisionParams) error
	GetRoutingDecisionLogs(ctx context.Context, input GetRoutingDecisionLogsRequest) ([]model.RoutingDecisionLog, error)
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
