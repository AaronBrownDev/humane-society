package domain

import (
	"context"

	"github.com/google/uuid"
)

// Adopter represents a person who adopts animals
type Adopter struct {
	AdopterID          uuid.UUID `json:"adopterId"`
	Person             Person    `json:"person"`
	HasPetAllergies    bool      `json:"hasPetAllergies"`
	HasSurrenderedPets bool      `json:"hasSurrenderedPets"`
	HomeStatus         string    `json:"homeStatus"`
}

// AdopterRepository defines the interface for adopter data access operations
type AdopterRepository interface {
	GetAll(ctx context.Context) ([]Adopter, error)
	GetByID(ctx context.Context, adopterID uuid.UUID) (*Adopter, error)
	Create(ctx context.Context, adopter *Adopter) error
	Update(ctx context.Context, adopter *Adopter) error
	Delete(ctx context.Context, adopterID uuid.UUID) error
}
