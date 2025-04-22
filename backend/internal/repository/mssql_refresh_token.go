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

		// Convert UUIDs from SQL Server format to standard UUID format
		token.TokenID, err = SwapUUIDFormat(token.TokenID)
		if err != nil {
			return nil, fmt.Errorf("error converting TokenID UUID: %w", err)
		}

		token.UserID, err = SwapUUIDFormat(token.UserID)
		if err != nil {
			return nil, fmt.Errorf("error converting UserID UUID: %w", err)
		}

		if revokedAt.Valid {
			token.RevokedAt = revokedAt.Time
		}

		if replacedBy.Valid {
			replacedID, err := uuid.Parse(replacedBy.String)
			if err != nil {
				return nil, err
			}

			// Convert ReplacedByTokenID from SQL Server format to standard UUID format
			replacedID, err = SwapUUIDFormat(replacedID)
			if err != nil {
				return nil, fmt.Errorf("error converting ReplacedByTokenID UUID: %w", err)
			}

			token.ReplacedByTokenID = replacedID
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (r *mssqlRefreshTokenRepository) GetByTokenID(ctx context.Context, tokenID uuid.UUID) (*domain.RefreshToken, error) {
	// Convert tokenID to SQL Server format for query
	sqlTokenID, err := SwapUUIDFormat(tokenID)
	if err != nil {
		return nil, fmt.Errorf("error converting TokenID UUID: %w", err)
	}

	query := `SELECT TokenID, UserID, Expires, CreatedAt, RevokedAt, ReplacedByTokenID
              FROM auth.RefreshToken 
              WHERE TokenID = @p1`

	var token domain.RefreshToken
	var revokedAt sql.NullTime
	var replacedBy sql.NullString

	err = r.conn.QueryRowContext(ctx, query, sqlTokenID).Scan(
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

	// Convert UUIDs from SQL Server format to standard UUID format
	token.TokenID, err = SwapUUIDFormat(token.TokenID)
	if err != nil {
		return nil, fmt.Errorf("error converting TokenID UUID: %w", err)
	}

	token.UserID, err = SwapUUIDFormat(token.UserID)
	if err != nil {
		return nil, fmt.Errorf("error converting UserID UUID: %w", err)
	}

	if revokedAt.Valid {
		token.RevokedAt = revokedAt.Time
	}

	if replacedBy.Valid {
		replacedID, err := uuid.Parse(replacedBy.String)
		if err != nil {
			return nil, err
		}

		// Convert ReplacedByTokenID from SQL Server format to standard UUID format
		replacedID, err = SwapUUIDFormat(replacedID)
		if err != nil {
			return nil, fmt.Errorf("error converting ReplacedByTokenID UUID: %w", err)
		}

		token.ReplacedByTokenID = replacedID
	}

	return &token, nil
}

func (r *mssqlRefreshTokenRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*domain.RefreshToken, error) {
	// Convert userID to SQL Server format for query
	sqlUserID, err := SwapUUIDFormat(userID)
	if err != nil {
		return nil, fmt.Errorf("error converting UserID UUID: %w", err)
	}

	// Get the latest non-revoked token for this user
	query := `SELECT TokenID, UserID, Expires, CreatedAt, RevokedAt, ReplacedByTokenID
              FROM auth.RefreshToken 
              WHERE UserID = @p1 AND RevokedAt IS NULL
              ORDER BY CreatedAt DESC`

	var token domain.RefreshToken
	var revokedAt sql.NullTime
	var replacedBy sql.NullString

	err = r.conn.QueryRowContext(ctx, query, sqlUserID).Scan(
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

	// Convert UUIDs from SQL Server format to standard UUID format
	token.TokenID, err = SwapUUIDFormat(token.TokenID)
	if err != nil {
		return nil, fmt.Errorf("error converting TokenID UUID: %w", err)
	}

	token.UserID, err = SwapUUIDFormat(token.UserID)
	if err != nil {
		return nil, fmt.Errorf("error converting UserID UUID: %w", err)
	}

	if revokedAt.Valid {
		token.RevokedAt = revokedAt.Time
	}

	if replacedBy.Valid {
		replacedID, err := uuid.Parse(replacedBy.String)
		if err != nil {
			return nil, err
		}

		// Convert ReplacedByTokenID from SQL Server format to standard UUID format
		replacedID, err = SwapUUIDFormat(replacedID)
		if err != nil {
			return nil, fmt.Errorf("error converting ReplacedByTokenID UUID: %w", err)
		}

		token.ReplacedByTokenID = replacedID
	}

	return &token, nil
}

func (r *mssqlRefreshTokenRepository) Create(ctx context.Context, token *domain.RefreshToken) error {
	query := `INSERT INTO auth.RefreshToken (TokenID, UserID, Expires, CreatedAt)
              VALUES (@p1, @p2, @p3, @p4)`

	// Convert UUIDs to SQL Server format for insertion
	sqlTokenID, err := SwapUUIDFormat(token.TokenID)
	if err != nil {
		return fmt.Errorf("error converting TokenID UUID: %w", err)
	}

	sqlUserID, err := SwapUUIDFormat(token.UserID)
	if err != nil {
		return fmt.Errorf("error converting UserID UUID: %w", err)
	}

	_, err = r.conn.ExecContext(
		ctx,
		query,
		sqlTokenID, // Use converted TokenID
		sqlUserID,  // Use converted UserID
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

	// Convert UUIDs to SQL Server format for query
	sqlTokenID, err := SwapUUIDFormat(token.TokenID)
	if err != nil {
		return nil, fmt.Errorf("error converting TokenID UUID: %w", err)
	}

	var revokedAt sql.NullTime
	var replacedBy sql.NullString

	if !token.RevokedAt.IsZero() {
		revokedAt = sql.NullTime{Time: token.RevokedAt, Valid: true}
	}

	if token.ReplacedByTokenID != uuid.Nil {
		// Convert ReplacedByTokenID to SQL Server format
		sqlReplacedID, err := SwapUUIDFormat(token.ReplacedByTokenID)
		if err != nil {
			return nil, fmt.Errorf("error converting ReplacedByTokenID UUID: %w", err)
		}
		replacedBy = sql.NullString{String: sqlReplacedID.String(), Valid: true}
	}

	_, err = r.conn.ExecContext(
		ctx,
		query,
		revokedAt,
		replacedBy,
		sqlTokenID,
	)
	if err != nil {
		return nil, fmt.Errorf("error updating refresh token: %w", err)
	}

	return token, nil
}

func (r *mssqlRefreshTokenRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM auth.RefreshToken WHERE TokenID = @p1`

	// Convert TokenID to SQL Server format for query
	sqlTokenID, err := SwapUUIDFormat(id)
	if err != nil {
		return fmt.Errorf("error converting TokenID UUID: %w", err)
	}

	result, err := r.conn.ExecContext(ctx, query, sqlTokenID)
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
