package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// UserAccount represents a user account for authentication
type UserAccount struct {
	UserID              uuid.UUID `json:"userId"`
	PasswordHash        string    `json:"-"` // Never expose in JSON responses
	LastLogin           time.Time `json:"lastLogin"`
	IsActive            bool      `json:"isActive"`
	FailedLoginAttempts int       `json:"-"`
	IsLocked            bool      `json:"isLocked"`
	LockoutEnd          time.Time `json:"lockoutEnd"`
	CreatedAt           time.Time `json:"createdAt"`
}

type UserAccountRepository interface {
	GetAll(ctx context.Context) ([]UserAccount, error)
	GetByID(ctx context.Context, id uuid.UUID) (*UserAccount, error)
	Create(ctx context.Context, userAccount UserAccount) (UserAccount, error)
	Update(ctx context.Context, userAccount UserAccount) (UserAccount, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
