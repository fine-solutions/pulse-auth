package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"pulse-auth/cmd/pulse/config"
)

type PGStorage struct {
	Conn *sqlx.DB
}

// NewPGStorage creates a new PGStorage instance connected to a PostgreSQL database using pgx driver.
func NewPGStorage(log *zap.Logger, cfg config.StorageConfig) (*PGStorage, error) {
	dbSource := fmt.Sprintf("postgres://%s:%s@localhost:%d/%s", cfg.Username, cfg.Password, cfg.Port, cfg.Name)
	conn, err := sqlx.Connect("pgx", dbSource)

	if err != nil {
		return nil, fmt.Errorf("connect to pgx failed %w", err)
	}

	err = conn.Ping()

	if err != nil {
		return nil, fmt.Errorf("ping failed %w", err)
	}

	return &PGStorage{Conn: conn}, nil
}
