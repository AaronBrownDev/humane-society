package domain

import (
	"context"
	"github.com/google/uuid"
)

// PetOwner represents a person who owns pets
type PetOwner struct {
	PetOwnerID                 uuid.UUID `json:"petOwnerId"`
	PersonID                   uuid.UUID `json:"personId"`
	VeterinarianID             uuid.UUID `json:"veterinarianId"`
	HasSterilizedPets          bool      `json:"hasSterilizedPets"`
	HasVaccinatedPets          bool      `json:"hasVaccinatedPets"`
	UsesVetHeartWormPrevention bool      `json:"usesVetHeartWormPrevention"`
}

// PetOwnerRepository defines the interface for pet owner data access operations
type PetOwnerRepository interface {
	GetAll(ctx context.Context) ([]PetOwner, error)
	GetByID(ctx context.Context, petOwnerID uuid.UUID) (*PetOwner, error)
	Create(ctx context.Context, petOwner *PetOwner) error
	Update(ctx context.Context, petOwner *PetOwner) error
	Delete(ctx context.Context, petOwnerID uuid.UUID) error

	// Domain-specific operations
	GetByVeterinarianID(ctx context.Context, veterinarianID uuid.UUID) ([]PetOwner, error)
	GetWithPets(ctx context.Context, petOwnerID uuid.UUID) (*PetOwner, []PetOwnerPet, error)
}
