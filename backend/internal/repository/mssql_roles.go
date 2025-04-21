package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
)

// mssqlRoleRepository implements domain.RoleRepository interface using SQL Server.
// It handles database operations for role entities, including core CRUD operations
// and role name lookups.
type mssqlRoleRepository struct {
	conn Connection
}

// NewMSSQLRole creates a new repository instance for role operations
// using the provided database connection.
func NewMSSQLRole(conn Connection) domain.RoleRepository {
	return &mssqlRoleRepository{conn: conn}
}

// fetch is a helper function to retrieve multiple role records based on the provided query.
// It handles row scanning, result collection, and proper resource cleanup.
// The function returns the matched roles or an error if the database operation fails.
func (r *mssqlRoleRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.Role, error) {
	rows, err := r.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []domain.Role
	for rows.Next() {
		var role domain.Role
		err = rows.Scan(
			&role.RoleID,
			&role.Name,
			&role.Description,
		)
		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

// GetAll returns all role records from the database.
// Returns a slice of all roles or an error if the database operation fails.
func (r *mssqlRoleRepository) GetAll(ctx context.Context) ([]domain.Role, error) {
	query := `SELECT RoleID, Name, Description 
              FROM auth.Role 
              ORDER BY RoleID`

	return r.fetch(ctx, query)
}

// GetByID returns a role with the specified ID.
// Returns a pointer to the role if found, domain.ErrNotFound if no matching role exists,
// or another error if the database operation fails.
func (r *mssqlRoleRepository) GetByID(ctx context.Context, id int) (*domain.Role, error) {
	query := `SELECT RoleID, Name, Description 
              FROM auth.Role 
              WHERE RoleID = @p1`

	var role domain.Role
	err := r.conn.QueryRowContext(ctx, query, id).Scan(
		&role.RoleID,
		&role.Name,
		&role.Description,
	)

	if err != nil {
		if errors.Is(err, errors.New("sql: no rows in result set")) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &role, nil
}

// Create inserts a new role record into the database.
// Returns the created role (with generated ID) or an error if the database operation fails.
// Returns domain.ErrConflict if a role with the same name already exists.
func (r *mssqlRoleRepository) Create(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	query := `INSERT INTO auth.Role (Name, Description)
              OUTPUT INSERTED.RoleID
              VALUES (@p1, @p2)`

	if role.Name == "" {
		return nil, domain.ErrInvalidInput
	}

	// Check if a role with the same name already exists
	existingRole, err := r.GetByName(ctx, role.Name)
	if err != nil && !errors.Is(err, domain.ErrNotFound) {
		return nil, err
	}
	if existingRole != nil {
		return nil, domain.ErrConflict
	}

	var roleID int
	err = r.conn.QueryRowContext(ctx, query, role.Name, role.Description).Scan(&roleID)
	if err != nil {
		return nil, fmt.Errorf("error creating role: %w", err)
	}

	// Update the role with the generated ID
	role.RoleID = roleID
	return role, nil
}

// Update modifies an existing role record in the database.
// Returns the updated role or an error if the database operation fails.
// Returns domain.ErrNotFound if no role with the given ID exists,
// or domain.ErrConflict if attempting to update to a name that already exists.
func (r *mssqlRoleRepository) Update(ctx context.Context, role *domain.Role) (*domain.Role, error) {
	query := `UPDATE auth.Role
              SET Name = @p1, Description = @p2
              WHERE RoleID = @p3`

	if role.RoleID <= 0 || role.Name == "" {
		return nil, domain.ErrInvalidInput
	}

	// Check if the role exists
	_, err := r.GetByID(ctx, role.RoleID)
	if err != nil {
		return nil, err
	}

	// Check if the new name conflicts with an existing role
	if role.Name != "" {
		existingRole, err := r.GetByName(ctx, role.Name)
		if err != nil && !errors.Is(err, domain.ErrNotFound) {
			return nil, err
		}
		if existingRole != nil && existingRole.RoleID != role.RoleID {
			return nil, domain.ErrConflict
		}
	}

	result, err := r.conn.ExecContext(
		ctx,
		query,
		role.Name,
		role.Description,
		role.RoleID,
	)
	if err != nil {
		return nil, fmt.Errorf("error updating role: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, domain.ErrNotFound
	}

	return role, nil
}

// Delete removes a role record from the database by its ID.
// Returns domain.ErrInvalidInput if the role ID is invalid,
// domain.ErrNotFound if no role with the given ID exists,
// or another error if the database operation fails.
func (r *mssqlRoleRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM auth.Role WHERE RoleID = @p1`

	if id <= 0 {
		return domain.ErrInvalidInput
	}

	result, err := r.conn.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting role: %w", err)
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
