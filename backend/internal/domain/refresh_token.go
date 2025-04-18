package domain

import (
	"github.com/google/uuid"
	"time"
)

// RefreshToken represents a JWT refresh token
type RefreshToken struct {
	TokenID         uuid.UUID `json:"-"`
	UserID          uuid.UUID `json:"-"`
	Token           string    `json:"token"`
	Expires         time.Time `json:"expires"`
	CreatedAt       time.Time `json:"-"`
	RevokedAt       time.Time `json:"-"`
	ReplacedByToken string    `json:"-"`
}

type RefreshTokenRepository interface {
}
