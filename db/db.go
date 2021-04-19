package db

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx" // driver
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

// Database
type Database struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

// Close - implements io.Closer
func (db *Database) Close() error {
	db.pool.Close()
	return nil
}

// Initialize - returns Database
//   # Example URL
//   postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10&pool_max_conn_lifetime=1h
func Initialize(ctx context.Context, databaseURL string, log *zap.Logger) (*Database, error) {
	const api = "Initialize"

	cfg, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		return nil, fmt.Errorf("%s.ParseConfig: %w", api, err)
	}

	// В этом режиме не будет подготовки и весь запрос пройдёт в одном сетевом вызове
	cfg.ConnConfig.PreferSimpleProtocol = true
	cfg.ConnConfig.RuntimeParams = map[string]string{"standard_conforming_strings": "on"}

	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("%s.ConnectConfig: %w", api, err)
	}

	if err = pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("%s.Ping: %w", api, err)
	}

	return &Database{
		pool:   pool,
		logger: log,
	}, nil
}
