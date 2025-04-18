package domain

import (
	"github.com/google/uuid"
	"time"
)

// UserAccount represents a user account for authentication
type UserAccount struct {
	UserID              uuid.UUID `json:"userId"`
	Email               string    `json:"email"`
	PasswordHash        string    `json:"-"` // Never expose in JSON responses
	LastLogin           time.Time `json:"lastLogin"`
	IsActive            bool      `json:"isActive"`
	FailedLoginAttempts int       `json:"-"`
	IsLocked            bool      `json:"isLocked"`
	LockoutEnd          time.Time `json:"lockoutEnd"`
	CreatedAt           time.Time `json:"createdAt"`
	PersonID            uuid.UUID `json:"personId,omitempty"`
	Roles               []Role    `json:"roles,omitempty"`
}

type UserAccountRepository interface {
}
