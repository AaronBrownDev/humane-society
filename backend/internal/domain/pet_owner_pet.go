package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// PetOwnerPet represents a pet owned by a pet owner
type PetOwnerPet struct {
	PetID             int       `json:"petId"`
	PetOwnerID        uuid.UUID `json:"petOwnerId"`
	Name              string    `json:"name"`
	Type              string    `json:"type"`
	Breed             string    `json:"breed"`
	Sex               string    `json:"sex"`
	OwnershipDate     time.Time `json:"ownershipDate"`
	LivingEnvironment string    `json:"livingEnvironment"`
}

// PetOwnerPetRepository defines the interface for pet owner's pets data access operations
type PetOwnerPetRepository interface {
	GetAll(ctx context.Context) ([]PetOwnerPet, error)
	GetByID(ctx context.Context, petID int) (*PetOwnerPet, error)
	GetByPetOwnerID(ctx context.Context, petOwnerID uuid.UUID) ([]PetOwnerPet, error)
	Create(ctx context.Context, pet *PetOwnerPet) error
	Update(ctx context.Context, pet *PetOwnerPet) error
	Delete(ctx context.Context, petID int) error
}
