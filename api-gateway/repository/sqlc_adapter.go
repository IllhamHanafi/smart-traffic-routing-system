package repository

import (
	"context"
	"fmt"
	"os"

	sqlc "github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/repository/sqlc"
	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/config"

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

func (s *SqlcRepository) Close() {
	s.conn.Close(context.Background())
}
