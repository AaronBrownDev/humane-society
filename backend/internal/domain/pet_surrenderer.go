package domain

import (
	"context"
	"github.com/google/uuid"
)

// PetSurrenderer represents a person who surrenders pets
type PetSurrenderer struct {
	SurrendererID uuid.UUID `json:"surrendererId"`
}

// PetSurrendererRepository defines operations for pet surrenderer data access
type PetSurrendererRepository interface {
	GetAll(ctx context.Context) ([]PetSurrenderer, error)
	GetByID(ctx context.Context, surrendererID uuid.UUID) (*PetSurrenderer, error)
	Create(ctx context.Context, surrenderer *PetSurrenderer) error
	Update(ctx context.Context, surrenderer *PetSurrenderer) error
	Delete(ctx context.Context, surrendererID uuid.UUID) error

	// Domain-specific methods
	GetSurrenderHistory(ctx context.Context, surrendererID uuid.UUID) ([]SurrenderForm, error)
}
