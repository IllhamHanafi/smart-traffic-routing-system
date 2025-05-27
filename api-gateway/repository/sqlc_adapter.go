package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/model"
	sqlc "github.com/IllhamHanafi/smart-traffic-routing-system/api-gateway/repository/sqlc"
	"github.com/IllhamHanafi/smart-traffic-routing-system/shared-libs/config"
	"github.com/google/uuid"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type SqlcRepository struct {
	queries *sqlc.Queries
	conn    *pgx.Conn
}

func NewSqlcRepository(cfg config.Database) Repository {
	ctx := context.Background()

	// pgx is optimized for postgresql
	databaseUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DatabaseName)
	conn, err := pgx.Connect(ctx, databaseUrl)
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

func (s *SqlcRepository) CreateUser(ctx context.Context, input CreateUserInput) (uuid.UUID, error) {
	now := time.Now()

	createdBy := uuid.Nil
	if input.CreatedBy != nil {
		createdBy = *input.CreatedBy
	}

	return s.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name:     input.Name,
		Role:     input.Role,
		Email:    input.Email,
		Password: input.Password,
		CreatedAt: pgtype.Timestamp{
			Time:  now,
			Valid: true,
		},
		CreatedBy: createdBy,
		UpdatedAt: pgtype.Timestamp{
			Time:  now,
			Valid: true,
		},
		UpdatedBy: createdBy,
	})

}

func (s *SqlcRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	res, err := s.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return model.User{}, err
	}

	return model.User{
		ID:       res.ID,
		Name:     res.Name,
		Role:     res.Role,
		Email:    res.Email,
		Password: res.Password,
	}, nil
}
