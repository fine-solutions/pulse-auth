package storage

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
)

var (
	ErrUserNotExist = errors.New("user not exist")
)

type Storage interface {
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, user *UserDB) error
	GetUserByUsername(ctx context.Context, username string) (*UserDB, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*UserDB, error)
	GetUsers(ctx context.Context) ([]*UserDB, error)
	DeleteUser(ctx context.Context, user *UserDB) error
}

type User struct {
	Username     string    `db:"username" json:"username"`
	PasswordHash string    `db:"password_hash" json:"-"`
	Email        string    `db:"email" json:"email"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

type UserDB struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Username     string    `db:"username" json:"username"`
	PasswordHash string    `db:"password_hash" json:"-"`
	Email        string    `db:"email" json:"email"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}
