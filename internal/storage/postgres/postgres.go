package postgres

import (
	"context"

	"main/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func ConnectDB(cfg *config.Config, log *zap.Logger) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseUrl)
	if err != nil {
		log.Fatal("Unable to connect to database", zap.Error(err))
	}

	var dbVersion string
	err = pool.QueryRow(context.Background(), "SELECT version()").Scan(&dbVersion)
	if err != nil {
		log.Fatal("Failed to execute query", zap.Error(err))
	}

	log.Info("Connected to PostgreSQL", zap.String("version", dbVersion))

	return pool
}
