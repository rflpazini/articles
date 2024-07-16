package database

import (
	"context"
	"fmt"

	"book-store/pkg/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPostgresDB(cfg *config.Config) (*pgxpool.Pool, error) {
	parseConfig, err := pgxpool.ParseConfig(cfg.DatabaseConfig.URI)
	if err != nil {
		return nil, fmt.Errorf("unable to parse DATABASE_URL: %v", err)
	}

	db, err := pgxpool.ConnectConfig(context.Background(), parseConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	return db, nil
}
