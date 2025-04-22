package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlRefreshTokenRepository struct {
	conn Connection
}

func NewMSSQLRefreshToken(conn Connection) domain.RefreshTokenRepository {
	return &mssqlRefreshTokenRepository{conn: conn}
}

func (r *mssqlRefreshTokenRepository) GetAll(ctx context.Context) ([]domain.RefreshToken, error) {
	query := `SELECT TokenID, UserID, Expires, CreatedAt, RevokedAt, ReplacedByTokenID
              FROM auth.RefreshToken`

	rows, err := r.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []domain.RefreshToken
	for rows.Next() {
		var token domain.RefreshToken
		var revokedAt sql.NullTime
		var replacedBy sql.NullString

		err = rows.Scan(
			&token.TokenID,
			&token.UserID,
			&token.Expires,
			&token.CreatedAt,
			&revokedAt,
			&replacedBy,
		)
		if err != nil {
			return nil, err
		}

		if revokedAt.Valid {
			token.RevokedAt = revokedAt.Time
		}

		if replacedBy.Valid {
			replacedID, err := uuid.Parse(replacedBy.String)
			if err != nil {
				return nil, err
			}
			token.ReplacedByTokenID = replacedID
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (r *mssqlRefreshTokenRepository) GetByTokenID(ctx context.Context, tokenID uuid.UUID) (*domain.RefreshToken, error) {
	// Debug logging
	fmt.Printf("GetByTokenID called with token ID: %s\n", tokenID.String())

	query := `SELECT TokenID, UserID, Expires, CreatedAt, RevokedAt, ReplacedByTokenID
              FROM auth.RefreshToken 
              WHERE TokenID = @p1`

	var token domain.RefreshToken
	var revokedAt sql.NullTime
	var replacedBy sql.NullString

	// Use tokenID directly without conversion
	err := r.conn.QueryRowContext(ctx, query, tokenID).Scan(
		&token.TokenID,
		&token.UserID,
		&token.Expires,
		&token.CreatedAt,
		&revokedAt,
		&replacedBy,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	if revokedAt.Valid {
		token.RevokedAt = revokedAt.Time
	}

	if replacedBy.Valid {
		replacedID, err := uuid.Parse(replacedBy.String)
		if err != nil {
			return nil, err
		}
		token.ReplacedByTokenID = replacedID
	}

	return &token, nil
}

func (r *mssqlRefreshTokenRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*domain.RefreshToken, error) {
	// Debug logging
	fmt.Printf("GetByUserID called with user ID: %s\n", userID.String())

	// Get the latest non-revoked token for this user
	query := `SELECT TokenID, UserID, Expires, CreatedAt, RevokedAt, ReplacedByTokenID
              FROM auth.RefreshToken 
              WHERE UserID = @p1 AND RevokedAt IS NULL
              ORDER BY CreatedAt DESC`

	var token domain.RefreshToken
	var revokedAt sql.NullTime
	var replacedBy sql.NullString

	// Use userID directly without conversion
	err := r.conn.QueryRowContext(ctx, query, userID).Scan(
		&token.TokenID,
		&token.UserID,
		&token.Expires,
		&token.CreatedAt,
		&revokedAt,
		&replacedBy,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	if revokedAt.Valid {
		token.RevokedAt = revokedAt.Time
	}

	if replacedBy.Valid {
		replacedID, err := uuid.Parse(replacedBy.String)
		if err != nil {
			return nil, err
		}
		token.ReplacedByTokenID = replacedID
	}

	return &token, nil
}

func (r *mssqlRefreshTokenRepository) Create(ctx context.Context, token *domain.RefreshToken) error {
	query := `INSERT INTO auth.RefreshToken (TokenID, UserID, Expires, CreatedAt)
              VALUES (@p1, @p2, @p3, @p4)`

	// Debug logging
	fmt.Printf("Creating refresh token - TokenID: %s, UserID: %s\n",
		token.TokenID.String(), token.UserID.String())

	// Use UUIDs directly without conversion
	_, err := r.conn.ExecContext(
		ctx,
		query,
		token.TokenID,
		token.UserID,
		token.Expires,
		token.CreatedAt,
	)
	if err != nil {
		return fmt.Errorf("error creating refresh token: %w", err)
	}

	return nil
}

func (r *mssqlRefreshTokenRepository) Update(ctx context.Context, token *domain.RefreshToken) (*domain.RefreshToken, error) {
	query := `UPDATE auth.RefreshToken
              SET RevokedAt = @p1, ReplacedByTokenID = @p2
              WHERE TokenID = @p3`

	// Debug logging
	fmt.Printf("Updating token ID: %s\n", token.TokenID.String())
	if token.ReplacedByTokenID != uuid.Nil {
		fmt.Printf("Replaced by token ID: %s\n", token.ReplacedByTokenID.String())
	}

	var revokedAt sql.NullTime
	var replacedBy sql.NullString

	if !token.RevokedAt.IsZero() {
		revokedAt = sql.NullTime{Time: token.RevokedAt, Valid: true}
	}

	if token.ReplacedByTokenID != uuid.Nil {
		replacedBy = sql.NullString{String: token.ReplacedByTokenID.String(), Valid: true}
	}

	// Use token ID directly without conversion
	_, err := r.conn.ExecContext(
		ctx,
		query,
		revokedAt,
		replacedBy,
		token.TokenID,
	)
	if err != nil {
		return nil, fmt.Errorf("error updating refresh token: %w", err)
	}

	return token, nil
}

func (r *mssqlRefreshTokenRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM auth.RefreshToken WHERE TokenID = @p1`

	// Debug logging
	fmt.Printf("Deleting token ID: %s\n", id.String())

	// Use ID directly without conversion
	result, err := r.conn.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting refresh token: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}
