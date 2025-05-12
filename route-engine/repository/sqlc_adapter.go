package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"route-engine/config"
	"route-engine/model"
	sqlc "route-engine/repository/sqlc"
	"time"

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

func (s *SqlcRepository) Close() {
	s.conn.Close(context.Background())
}
