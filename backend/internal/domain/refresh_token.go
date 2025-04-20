package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// RefreshToken represents a JWT refresh token
type RefreshToken struct {
	TokenID           uuid.UUID `json:"tokenID"`
	UserID            uuid.UUID `json:"-"`
	Expires           time.Time `json:"expires"`
	CreatedAt         time.Time `json:"-"`
	RevokedAt         time.Time `json:"-"`
	ReplacedByTokenID uuid.UUID `json:"-"`
}

type RefreshTokenRepository interface {
	GetAll(ctx context.Context) ([]RefreshToken, error)
	GetByTokenID(ctx context.Context, tokenID uuid.UUID) (*RefreshToken, error)
	GetByUserID(ctx context.Context, userID uuid.UUID) (*RefreshToken, error)
	Create(ctx context.Context, token *RefreshToken) error
	Update(ctx context.Context, token *RefreshToken) (*RefreshToken, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
