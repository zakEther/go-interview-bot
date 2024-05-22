package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/zakether/go-interview-bot/pkg/config"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPostgreSQLStorage(cfg config.StorageConfig) (*Postgres, error) {

	const op = "pkg/postgres/postgres.go"

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.User,
		cfg.Database,
		cfg.Password,
		"disable",
	)

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database: %s, %w", op, err)
	}

	return &Postgres{
		Pool: pool,
	}, nil

}
