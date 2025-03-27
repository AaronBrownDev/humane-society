package domain

import (
	"context"
	"github.com/google/uuid"
)

// Veterinarian represents a veterinarian
type Veterinarian struct {
	VeterinarianID uuid.UUID `json:"veterinarianId"`
}

// VeterinarianRepository defines the interface for veterinarian data access operations
type VeterinarianRepository interface {
	GetAll(ctx context.Context) ([]Veterinarian, error)
	GetByID(ctx context.Context, veterinarianID uuid.UUID) (*Veterinarian, error)
	Create(ctx context.Context, veterinarian *Veterinarian) error
	Update(ctx context.Context, veterinarian *Veterinarian) error
	Delete(ctx context.Context, veterinarianID uuid.UUID) error
}
