package model

import (
	"fmt"
	"time"

	"gopkg.in/validator.v2"
)

type UserID string

func (id UserID) String() string {
	return string(id)
}

type TokenID string

func (id TokenID) String() string {
	return string(id)
}

type User struct {
	UserID     UserID
	Username   string
	FirstName  string
	SecondName string
	Sex        string
	Birthdate  time.Time
	Biography  string
	City       string
}

type Token struct {
	TokenID TokenID
	UserID  UserID
	Token   string
}

type TokenWithMetadata struct {
	TokenID  string    `validator:"nonzero"`
	UserID   UserID    `validator:"nonzero"`
	Token    string    `validator:"nonzero"`
	AlivedAt time.Time `validator:"nonzero"`
}

func (t *TokenWithMetadata) Validate() error {
	if err := validator.Validate(t); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	return nil
}

type UserLogin struct {
	Username       string
	HashedPassword string
}

type UserRegister struct {
	ID             string `validator:"nonzero"`
	Username       string `validator:"nonzero"`
	HashedPassword string `validator:"nonzero"`
	FirstName      string
	SecondName     string
	Sex            string
	Birthdate      time.Time
	Biography      string
	City           string
}

func (u *UserRegister) Validate() error {
	if err := validator.Validate(u); err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	return nil
}
