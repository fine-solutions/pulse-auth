package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"pulse-auth/cmd/pulse/config"
	"pulse-auth/internal/storage"
)

type Storage struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewStorage(logger *zap.Logger, config config.StorageConfig) (*Storage, error) {
	postgresStorage, err := NewPGStorage(logger, config)
	if err != nil {
		return nil, fmt.Errorf("new pg storage: %w", err)
	}
	return &Storage{
		db:     postgresStorage.Conn,
		logger: logger,
	}, nil
}

func (s *Storage) User() storage.UserRepository {
	return s
}

func (s *Storage) Token() storage.TokenRepository {
	return s
}
