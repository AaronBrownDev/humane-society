package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type UserRole struct {
	UserID   uuid.UUID `json:"user_id"`
	RoleID   int       `json:"role_id"`
	AssignAt time.Time `json:"assign_at"`
}

type UserRoleRepository interface {
	GetAll(ctx context.Context) ([]UserRole, error)
	GetByUserID(ctx context.Context, id uuid.UUID) (*UserRole, error)
	GetByRoleID(ctx context.Context, id int) (*UserRole, error)
	Create(ctx context.Context, userRole UserRole) (UserRole, error)
	Update(ctx context.Context, userRole UserRole) (UserRole, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
