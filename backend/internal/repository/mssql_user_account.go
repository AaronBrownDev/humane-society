package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

// mssqlUserAccountRepository implements domain.UserAccountRepository interface using SQL Server.
// It handles database operations for user authentication accounts, including creation with role assignment,
// account locking, and nullable time fields.
type mssqlUserAccountRepository struct {
	conn Connection
}

// NewMSSQLUserAccount creates a new repository instance for user account operations
// using the provided database connection.
func NewMSSQLUserAccount(conn Connection) domain.UserAccountRepository {
	return &mssqlUserAccountRepository{conn: conn}
}

// GetAll retrieves all user accounts from the database.
// It handles nullable time fields for LastLogin and LockoutEnd.
// Returns a slice of all user accounts or an error if the database operation fails.
func (r *mssqlUserAccountRepository) GetAll(ctx context.Context) ([]domain.UserAccount, error) {
	query := `SELECT UserID, PasswordHash, LastLogin, IsActive, FailedLoginAttempts, IsLocked, LockoutEnd, CreatedAt
              FROM auth.UserAccount`

	rows, err := r.conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []domain.UserAccount
	for rows.Next() {
		var account domain.UserAccount
		var lastLogin, lockoutEnd sql.NullTime

		err = rows.Scan(
			&account.UserID,
			&account.PasswordHash,
			&lastLogin,
			&account.IsActive,
			&account.FailedLoginAttempts,
			&account.IsLocked,
			&lockoutEnd,
			&account.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if lastLogin.Valid {
			account.LastLogin = lastLogin.Time
		}

		if lockoutEnd.Valid {
			account.LockoutEnd = lockoutEnd.Time
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

// GetByID retrieves a specific user account by its UUID.
// It handles nullable time fields for LastLogin and LockoutEnd.
// Returns a pointer to the user account if found, domain.ErrNotFound if no matching account exists,
// or another error if the database operation fails.
func (r *mssqlUserAccountRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.UserAccount, error) {
	query := `SELECT UserID, PasswordHash, LastLogin, IsActive, FailedLoginAttempts, IsLocked, LockoutEnd, CreatedAt
              FROM auth.UserAccount 
              WHERE UserID = @p1`

	var account domain.UserAccount
	var lastLogin, lockoutEnd sql.NullTime

	err := r.conn.QueryRowContext(ctx, query, id).Scan(
		&account.UserID,
		&account.PasswordHash,
		&lastLogin,
		&account.IsActive,
		&account.FailedLoginAttempts,
		&account.IsLocked,
		&lockoutEnd,
		&account.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	if lastLogin.Valid {
		account.LastLogin = lastLogin.Time
	}

	if lockoutEnd.Valid {
		account.LockoutEnd = lockoutEnd.Time
	}

	return &account, nil
}

// Create inserts a new user account into the database using the auth.CreatePublicUser stored procedure.
// This procedure automatically assigns the Public role to the new user and handles proper transaction management.
// Returns the created user account or an error if the database operation fails.
func (r *mssqlUserAccountRepository) Create(ctx context.Context, userAccount domain.UserAccount) (domain.UserAccount, error) {
	query := `EXEC auth.CreatePublicUser @UserID = @p1, @PasswordHash = @p2`

	_, err := r.conn.ExecContext(ctx, query, userAccount.UserID, userAccount.PasswordHash)
	if err != nil {
		return domain.UserAccount{}, fmt.Errorf("error creating user account: %w", err)
	}

	return userAccount, nil
}

// Update modifies an existing user account in the database.
// It properly handles zero time values by converting them to SQL NULL values when appropriate.
// Returns the updated user account or an error if the database operation fails,
// such as when the account doesn't exist or database constraints are violated.
func (r *mssqlUserAccountRepository) Update(ctx context.Context, userAccount domain.UserAccount) (domain.UserAccount, error) {
	query := `UPDATE auth.UserAccount
              SET PasswordHash = @p1, 
                  LastLogin = @p2, 
                  IsActive = @p3, 
                  FailedLoginAttempts = @p4, 
                  IsLocked = @p5, 
                  LockoutEnd = @p6
              WHERE UserID = @p7`

	var lastLogin, lockoutEnd sql.NullTime

	if !userAccount.LastLogin.IsZero() {
		lastLogin = sql.NullTime{Time: userAccount.LastLogin, Valid: true}
	}

	if !userAccount.LockoutEnd.IsZero() {
		lockoutEnd = sql.NullTime{Time: userAccount.LockoutEnd, Valid: true}
	}

	_, err := r.conn.ExecContext(
		ctx,
		query,
		userAccount.PasswordHash,
		lastLogin,
		userAccount.IsActive,
		userAccount.FailedLoginAttempts,
		userAccount.IsLocked,
		lockoutEnd,
		userAccount.UserID,
	)
	if err != nil {
		return domain.UserAccount{}, fmt.Errorf("error updating user account: %w", err)
	}

	return userAccount, nil
}

// Delete removes a user account from the database by its UUID.
// Returns domain.ErrNotFound if no account with the given ID exists,
// or a wrapped database error if the deletion fails for other reasons,
// such as constraint violations.
func (r *mssqlUserAccountRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM auth.UserAccount WHERE UserID = @p1`

	result, err := r.conn.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting user account: %w", err)
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
