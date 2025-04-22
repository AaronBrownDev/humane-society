// role_service.go
package services

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

// RoleService provides functionality for managing roles and permissions
type RoleService struct {
	roleRepo     domain.RoleRepository
	userRoleRepo domain.UserRoleRepository
}

// NewRoleService creates a new instance of RoleService
func NewRoleService(roleRepo domain.RoleRepository, userRoleRepo domain.UserRoleRepository) *RoleService {
	return &RoleService{
		roleRepo:     roleRepo,
		userRoleRepo: userRoleRepo,
	}
}

// GetUserRole returns the domain.Role assigned to a user
func (s *RoleService) GetUserRole(ctx context.Context, userID uuid.UUID) (*domain.Role, error) {
	// Get the user's role
	userRole, err := s.userRoleRepo.GetByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	// Get the role details
	role, err := s.roleRepo.GetByID(ctx, userRole.RoleID)
	if err != nil {
		return nil, err
	}

	return role, nil
}

// GetRoleName returns the name of the role assigned to a user
func (s *RoleService) GetRoleName(ctx context.Context, userID uuid.UUID) (string, error) {
	role, err := s.GetUserRole(ctx, userID)
	if err != nil {
		return "", err
	}
	return role.Name, nil
}

// AssignRoleToUser assigns a role to a user
func (s *RoleService) AssignRoleToUser(ctx context.Context, userID uuid.UUID, roleName string) error {
	// Find the role by name
	roles, err := s.roleRepo.GetAll(ctx)
	if err != nil {
		return err
	}

	var roleID int
	found := false
	for _, role := range roles {
		if role.Name == roleName {
			roleID = role.RoleID
			found = true
			break
		}
	}

	if !found {
		return domain.ErrNotFound
	}

	// Create the user role
	userRole := domain.UserRole{
		UserID: userID,
		RoleID: roleID,
	}

	_, err = s.userRoleRepo.Create(ctx, userRole)
	return err
}
