package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/model"
	sqlc "github.com/IllhamHanafi/smart-traffic-routing-system/route-engine/repository/sqlc"
	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/config"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type SqlcRepository struct {
	queries *sqlc.Queries
	conn    *pgx.Conn
}

func NewSqlcRepository(cfg config.Database) Repository {
	// pgx is optimized for postgresql
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DatabaseName)
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	queries := sqlc.New(conn)
	return &SqlcRepository{
		conn:    conn,
		queries: queries,
	}
}

func (s *SqlcRepository) GetActiveRoutingLogic(ctx context.Context) (ActiveRoutingLogicResult, error) {
	res, err := s.queries.GetActiveRoutingLogic(ctx)
	if err != nil {
		return ActiveRoutingLogicResult{}, err
	}

	allocationLogicMap := make(map[string]int)
	err = json.Unmarshal(res.AllocationLogic, &allocationLogicMap)
	if err != nil {
		return ActiveRoutingLogicResult{}, err
	}

	return ActiveRoutingLogicResult{
		ID:              res.ID,
		AllocationLogic: allocationLogicMap,
	}, nil
}

func (s *SqlcRepository) InsertRoutingDecisionLog(ctx context.Context, input InsertRoutingDecisionLogParams) error {
	reason := pgtype.Text{}
	if input.Reason != "" {
		reason = pgtype.Text{String: input.Reason, Valid: true}
	}
	createdAt := pgtype.Timestamp{
		Time:  time.Now(),
		Valid: true,
	}

	err := s.queries.InsertRoutingDecisionLog(ctx, sqlc.InsertRoutingDecisionLogParams{
		OrderID:           input.OrderID,
		CourierID:         input.CourierID,
		RoutingDecisionID: input.RoutingDecisionID,
		Status:            input.Status,
		Reason:            reason,
		CreatedBy:         input.CreatedBy,
		CreatedAt:         createdAt,
	})
	return err
}

func (s *SqlcRepository) GetCourierByCode(ctx context.Context, courierCode string) (model.Courier, error) {
	res, err := s.queries.GetCourierByCode(ctx, courierCode)
	if err != nil {
		return model.Courier{}, err
	}
	return model.Courier{
		ID:   res.ID,
		Code: res.Code,
		Name: res.Name,
	}, nil
}

func (s *SqlcRepository) CreateActiveRoutingDecision(ctx context.Context, input InsertNewRoutingDecisionParams) error {
	currentTime := time.Now()
	// convert map to json
	allocationLogic, err := json.Marshal(input.AllocationLogic)
	if err != nil {
		return err
	}

	tx, err := s.conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	qtx := s.queries.WithTx(tx)

	// inactivate current active routing decision
	err = qtx.InactivateCurrentActiveRoutingDecision(ctx, sqlc.InactivateCurrentActiveRoutingDecisionParams{
		UpdatedAt: pgtype.Timestamp{
			Time:  currentTime,
			Valid: true,
		},
		UpdatedBy: input.UserID,
	})
	if err != nil {
		return err
	}

	// insert new routing decision with status active
	if err := qtx.InsertActiveRoutingDecision(ctx, sqlc.InsertActiveRoutingDecisionParams{
		AllocationLogic: allocationLogic,
		CreatedAt: pgtype.Timestamp{
			Time:  currentTime,
			Valid: true,
		},
		CreatedBy: input.UserID,
		UpdatedAt: pgtype.Timestamp{
			Time:  currentTime,
			Valid: true,
		},
		UpdatedBy: input.UserID,
	}); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

func (s *SqlcRepository) GetRoutingDecisionLogs(ctx context.Context, input GetRoutingDecisionLogsRequest) ([]model.RoutingDecisionLog, error) {
	params := sqlc.GetRoutingDecisionLogsParams{
		Limit:  input.Limit,
		Offset: input.Offset,
	}
	if input.OrderID != nil {
		params.OrderID = pgtype.UUID{
			Bytes: *input.OrderID,
			Valid: true,
		}
	}
	if input.CourierID != nil {
		params.CourierID = pgtype.UUID{
			Bytes: *input.CourierID,
			Valid: true,
		}
	}
	if input.RoutingDecisionID != nil {
		params.RoutingDecisionID = pgtype.UUID{
			Bytes: *input.RoutingDecisionID,
			Valid: true,
		}
	}
	if input.Status != nil {
		params.Status = pgtype.Text{
			String: *input.Status,
			Valid:  true,
		}
	}

	res, err := s.queries.GetRoutingDecisionLogs(ctx, params)
	if err != nil {
		return []model.RoutingDecisionLog{}, err
	}
	var items []model.RoutingDecisionLog
	for _, i := range res {
		items = append(items, model.RoutingDecisionLog{
			ID:                i.ID,
			OrderID:           i.OrderID,
			CourierID:         i.CourierID,
			RoutingDecisionID: i.RoutingDecisionID,
			Status:            i.Status,
			Reason:            i.Reason.String,
			CreatedAt:         i.CreatedAt.Time,
			CreatedBy:         i.CreatedBy,
		})
	}
	return items, nil
}

func (s *SqlcRepository) Close() {
	s.conn.Close(context.Background())
}
