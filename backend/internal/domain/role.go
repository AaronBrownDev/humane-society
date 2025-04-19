package domain

import (
	"context"
)

// Role represents a security role in the system
type Role struct {
	RoleID      int    `json:"roleId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type RoleRepository interface {
	GetAll(ctx context.Context) ([]Role, error)
	GetByID(ctx context.Context, id int) (*Role, error)
	Create(ctx context.Context, role *Role) (*Role, error)
	Update(ctx context.Context, role *Role) (*Role, error)
	Delete(ctx context.Context, id int) error
}
