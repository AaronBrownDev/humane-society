package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
	"time"
)

// mssqlUserRoleRepository implements domain.UserRoleRepository interface using SQL Server.
// It handles database operations for user role assignments, enforcing that each user
// has exactly one role in the system.
type mssqlUserRoleRepository struct {
	conn Connection
}

// NewMSSQLUserRole creates a new repository instance for user role operations
// using the provided database connection.
func NewMSSQLUserRole(conn Connection) domain.UserRoleRepository {
	return &mssqlUserRoleRepository{conn: conn}
}

// fetch is a helper function to retrieve multiple user role records based on the provided query.
// It handles row scanning, result collection, and proper resource cleanup.
// The function returns the matched user roles or an error if the database operation fails.
func (r *mssqlUserRoleRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.UserRole, error) {
	rows, err := r.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userRoles []domain.UserRole
	for rows.Next() {
		var userRole domain.UserRole
		err = rows.Scan(
			&userRole.UserID,
			&userRole.RoleID,
			&userRole.AssignAt,
		)
		if err != nil {
			return nil, err
		}

		// Convert SQL Server UUID format to standard UUID format
		userRole.UserID, err = SwapUUIDFormat(userRole.UserID)
		if err != nil {
			return nil, fmt.Errorf("error converting UserID UUID: %w", err)
		}

		userRoles = append(userRoles, userRole)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return userRoles, nil
}

// GetAll returns all user role assignments from the database.
// Returns a slice of all user roles or an error if the database operation fails.
func (r *mssqlUserRoleRepository) GetAll(ctx context.Context) ([]domain.UserRole, error) {
	query := `SELECT UserID, RoleID, AssignedAt 
              FROM auth.UserRole`

	return r.fetch(ctx, query)
}

// GetByUserID retrieves the role assignment for a specific user.
// Returns a single UserRole instead of a slice since each user can only have one role.
// Returns domain.ErrInvalidInput if the user ID is nil,
// domain.ErrNotFound if no role is assigned to the user,
// or another error if the database operation fails.
func (r *mssqlUserRoleRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*domain.UserRole, error) {
	query := `SELECT UserID, RoleID, AssignedAt 
              FROM auth.UserRole 
              WHERE UserID = @p1`

	if userID == uuid.Nil {
		return nil, domain.ErrInvalidInput
	}

	var userRole domain.UserRole
	err := r.conn.QueryRowContext(ctx, query, userID).Scan(
		&userRole.UserID,
		&userRole.RoleID,
		&userRole.AssignAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	// Convert SQL Server UUID format to standard UUID format
	userRole.UserID, err = SwapUUIDFormat(userRole.UserID)
	if err != nil {
		return nil, fmt.Errorf("error converting UserID UUID: %w", err)
	}

	return &userRole, nil
}

// GetByRoleID retrieves all user assignments for a specific role.
// Returns domain.ErrInvalidInput if the role ID is invalid,
// domain.ErrNotFound if no users are assigned to the role,
// or another error if the database operation fails.
func (r *mssqlUserRoleRepository) GetByRoleID(ctx context.Context, roleID int) ([]domain.UserRole, error) {
	query := `SELECT UserID, RoleID, AssignedAt 
              FROM auth.UserRole 
              WHERE RoleID = @p1`

	if roleID <= 0 {
		return nil, domain.ErrInvalidInput
	}

	roles, err := r.fetch(ctx, query, roleID)
	if err != nil {
		return nil, err
	}

	if len(roles) == 0 {
		return nil, domain.ErrNotFound
	}

	return roles, nil
}

// Create inserts or updates a user role assignment in the database.
// Because each user can only have one role, this will replace any existing role.
// Sets the assignment time to the current time if not provided.
// Returns the created/updated user role or an error if the database operation fails.
func (r *mssqlUserRoleRepository) Create(ctx context.Context, userRole domain.UserRole) (domain.UserRole, error) {
	// First, check if the user already has a role
	existingRole, err := r.GetByUserID(ctx, userRole.UserID)

	// Set assignment time to current time if not provided
	if userRole.AssignAt.IsZero() {
		userRole.AssignAt = time.Now()
	}

	// Check for invalid inputs
	if userRole.UserID == uuid.Nil || userRole.RoleID <= 0 {
		return domain.UserRole{}, domain.ErrInvalidInput
	}

	// If user already has a role, update it
	if err == nil && existingRole != nil {
		return r.Update(ctx, userRole)
	}

	// If error is not "not found", return it
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		return domain.UserRole{}, err
	}

	// Otherwise, insert a new role assignment
	query := `INSERT INTO auth.UserRole
              (UserID, RoleID, AssignedAt)
              VALUES
              (@p1, @p2, @p3)`

	_, err = r.conn.ExecContext(
		ctx,
		query,
		userRole.UserID,
		userRole.RoleID,
		userRole.AssignAt,
	)
	if err != nil {
		return domain.UserRole{}, fmt.Errorf("error creating user role: %w", err)
	}

	return userRole, nil
}

// Update modifies an existing user role assignment in the database.
// This changes the user's role and updates the assignment date.
// Returns the updated user role or an error if the database operation fails.
func (r *mssqlUserRoleRepository) Update(ctx context.Context, userRole domain.UserRole) (domain.UserRole, error) {
	query := `UPDATE auth.UserRole
              SET RoleID = @p1, AssignedAt = @p2
              WHERE UserID = @p3`

	if userRole.UserID == uuid.Nil || userRole.RoleID <= 0 {
		return domain.UserRole{}, domain.ErrInvalidInput
	}

	// Ensure assignment time is set
	if userRole.AssignAt.IsZero() {
		userRole.AssignAt = time.Now()
	}

	result, err := r.conn.ExecContext(
		ctx,
		query,
		userRole.RoleID,
		userRole.AssignAt,
		userRole.UserID,
	)
	if err != nil {
		return domain.UserRole{}, fmt.Errorf("error updating user role: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return domain.UserRole{}, err
	}
	if rowsAffected == 0 {
		return domain.UserRole{}, domain.ErrNotFound
	}

	return userRole, nil
}

// Delete removes the role assignment for a specific user.
// Returns domain.ErrInvalidInput if the user ID is nil,
// domain.ErrNotFound if no assignment exists for the user,
// or another error if the database operation fails.
func (r *mssqlUserRoleRepository) Delete(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM auth.UserRole WHERE UserID = @p1`

	if userID == uuid.Nil {
		return domain.ErrInvalidInput
	}

	result, err := r.conn.ExecContext(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("error deleting user role: %w", err)
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
