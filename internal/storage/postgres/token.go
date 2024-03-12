package postgres

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"pulse-auth/internal/model"
	"pulse-auth/internal/utils"
	"time"
)

// GetCurrentUserToken get current user token from storage based on user ID.
func (s *Storage) GetCurrentUserToken(ctx context.Context, id model.UserID) (*model.Token, error) {
	sql, args, err := sq.Select(tokenFields...).
		From(TokenTable).
		Where(
			sq.Eq{
				fieldUserID:    id.String(),
				fieldDeletedAt: nil,
			},
			sq.LtOrEq{
				fieldAlivedAt: time.Now(),
			},
		).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, utils.WrapInternalError(fmt.Errorf("incorrect sql"))
	}

	var entity tokenEntity
	err = s.db.GetContext(ctx, &entity, sql, args...)
	if err != nil {
		return nil, utils.WrapSqlError(err)
	}

	return tokenEntityToModel(entity), nil
}

// CreateToken creates a token with the provided parameters and stores it in the database, returning the generated token or an error message.
func (s *Storage) CreateToken(ctx context.Context, params *model.TokenWithMetadata) (*model.Token, error) {
	err := params.Validate()
	if err != nil {
		return nil, utils.WrapValidationError(fmt.Errorf("params validate: %w", err))
	}

	now := time.Now().Truncate(time.Millisecond)
	sql, args, err := sq.Insert(TokenTable).
		Columns(tokenFields...).
		Values(params.TokenID, params.UserID, params.Token, now, params.AlivedAt).
		Suffix(returningToken).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, utils.WrapInternalError(fmt.Errorf("incorrect sql"))
	}

	var entity tokenEntity
	err = s.db.GetContext(ctx, &entity, sql, args...)
	if err != nil {
		return nil, utils.WrapSqlError(err)
	}

	return tokenEntityToModel(entity), nil
}

// RevokeToken revokes a token by updating the deleted_at field with the current timestamp in the database.
func (s *Storage) RevokeToken(ctx context.Context, params *model.Token) error {
	now := time.Now().Truncate(time.Millisecond)
	sql, args, err := sq.Update(TokenTable).
		Where(sq.Eq{
			fieldToken:     params.Token,
			fieldUserID:    params.UserID.String(),
			fieldDeletedAt: nil,
		}).
		Set(fieldDeletedAt, now).
		Suffix(returningToken).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return utils.WrapInternalError(fmt.Errorf("incorrect sql"))
	}

	result, err := s.db.ExecContext(ctx, sql, args...)

	if err != nil {
		return utils.WrapSqlError(err)
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return utils.WrapInternalError(err)
	}

	if rowsAffected == 0 {
		return utils.WrapNotFoundError(err, utils.NotFoundMessage)
	}

	return nil
}

// RefreshToken updates the token's value in the database based on the provided parameters, returning the updated token or an error message.
func (s *Storage) RefreshToken(ctx context.Context, params *model.TokenWithMetadata) (*model.Token, error) {
	err := params.Validate()
	if err != nil {
		return nil, fmt.Errorf("params validate: %w", err)
	}
	sql, args, err := sq.Update(TokenTable).
		Where(sq.Eq{
			fieldToken:     params.Token,
			fieldUserID:    params.UserID.String(),
			fieldDeletedAt: nil,
		}).
		Set(fieldAlivedAt, params.AlivedAt).
		Suffix(returningToken).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, utils.WrapInternalError(fmt.Errorf("incorrect sql"))
	}

	var entity tokenEntity
	err = s.db.GetContext(ctx, &entity, sql, args...)
	if err != nil {
		return nil, utils.WrapSqlError(err)
	}

	return tokenEntityToModel(entity), nil
}

type tokenEntity struct {
	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	Token     string    `db:"token"`
	CreatedAt time.Time `db:"created_at"`
	DeletedAt time.Time `db:"deleted_at"`
	AlivedAt  time.Time `db:"alived_at"`
}

// tokenEntityToModel converts a token entity to a token model and returns a pointer to it.
func tokenEntityToModel(entity tokenEntity) *model.Token {
	return &model.Token{
		UserID: model.UserID(entity.UserID),
		Token:  entity.Token,
	}
}
