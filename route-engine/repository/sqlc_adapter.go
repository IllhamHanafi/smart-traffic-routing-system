package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"route-engine/config"
	sqlc "route-engine/repository/sqlc"

	"github.com/jackc/pgx/v5"
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

func (s *SqlcRepository) Close() {
	s.conn.Close(context.Background())
}
