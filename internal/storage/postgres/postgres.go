package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"pulse-auth/internal/config"
	"pulse-auth/internal/lib/e"
	"pulse-auth/internal/storage"
)

type Storage struct {
	db  *sqlx.DB
	log *slog.Logger
}

// New creates a new Storage instance connected to a PostgreSQL database using pgx driver.
func New(log *slog.Logger, cfg *config.Config) (*Storage, error) {
	dbSource := fmt.Sprintf("postgres://%s:%s@localhost:%d/%s", cfg.DB.Username, cfg.DB.Password, cfg.DB.Port, cfg.DB.Name)
	conn, err := sqlx.Connect("pgx", dbSource)
	if err != nil {
		return nil, e.Wrap("connect to pgx failed", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, e.Wrap("ping failed", err)
	}
	return &Storage{db: conn, log: log}, nil
}

// CreateUser creates a new user in the storage.
func (s *Storage) CreateUser(ctx context.Context, user *storage.User) error {
	q := `INSERT INTO users (username, password_hash, email, created_at) VALUES ($1, $2, $3, $4)`
	if _, err := s.db.ExecContext(ctx, q, user.Username, user.PasswordHash, user.Email, user.CreatedAt); err != nil {
		return e.Wrap("can't create user in storage", err)
	}

	s.log.Debug("create user", slog.String("username", user.Username))

	return nil
}

// UpdateUser updates user information in the storage.
func (s *Storage) UpdateUser(ctx context.Context, user *storage.UserDB) error {
	q := `UPDATE users SET username = $1, password_hash = $2, email = $3 WHERE username = $4`
	_, err := s.db.ExecContext(ctx, q, user.Username, user.PasswordHash, user.Email, user.Username)

	if err != nil {
		return e.Wrap("can't update user in storage", err)
	}

	s.log.Debug("update user", slog.String("username", user.Username))

	return nil
}

// GetUserByUsername returns user from storage by username.
func (s *Storage) GetUserByUsername(ctx context.Context, username string) (*storage.UserDB, error) {
	q := `SELECT * FROM users WHERE username = $1`
	var user storage.UserDB

	err := s.db.GetContext(ctx, &user, q, username)

	if err == sql.ErrNoRows {
		return nil, storage.ErrUserNotExist
	}

	if err != nil {
		return nil, e.Wrap("can't get user from storage by username", err)
	}

	return &user, nil
}

// GetUserByID returns user from storage by id.
func (s *Storage) GetUserByID(ctx context.Context, id uuid.UUID) (*storage.UserDB, error) {
	q := `SELECT * FROM users WHERE id = $1`

	var user storage.UserDB

	err := s.db.GetContext(ctx, &user, q, id)

	if err == sql.ErrNoRows {
		return nil, storage.ErrUserNotExist
	}

	if err != nil {
		return nil, e.Wrap("can't get user from storage by id", err)
	}

	return &user, nil
}

// GetUsers returns all users from storage.
func (s *Storage) GetUsers(ctx context.Context) ([]*storage.UserDB, error) {
	q := `SELECT * FROM users`

	users := []*storage.UserDB{}

	err := s.db.SelectContext(ctx, &users, q)
	if err != nil {
		return nil, e.Wrap("can't get all users from storage", err)
	}

	return users, nil
}

// DeleteUser deletes a user from the storage by username.
func (s *Storage) DeleteUser(ctx context.Context, user *storage.UserDB) error {
	q := `DELETE FROM users WHERE username = $1`

	_, err := s.db.ExecContext(ctx, q, user.Username)
	if err != nil {
		return e.Wrap("can't delete user from storage", err)
	}

	s.log.Debug("delete user", slog.String("username", user.Username))
	return nil
}
