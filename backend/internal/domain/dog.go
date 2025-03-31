package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// Dog represents a dog in the shelter
type Dog struct {
	DogID              uuid.UUID `json:"dogId"`
	Name               string    `json:"name"`
	IntakeDate         time.Time `json:"intakeDate"`
	EstimatedBirthDate time.Time `json:"estimatedBirthDate"`
	Breed              string    `json:"breed"`
	Sex                string    `json:"sex"`
	Color              string    `json:"color"`
	CageNumber         int       `json:"cageNumber"`
	IsAdopted          bool      `json:"isAdopted"`
}

// DogRepository defines the interface for dog data access operations.
// Provides methods for managing dogs.
type DogRepository interface {
	// General CRUD operations
	GetAll(ctx context.Context) ([]Dog, error)
	GetAvailable(ctx context.Context) ([]Dog, error)
	GetByID(ctx context.Context, dogID uuid.UUID) (*Dog, error)
	Create(ctx context.Context, dog *Dog) error
	Update(ctx context.Context, dog *Dog) error
	Delete(ctx context.Context, dogID uuid.UUID) error

	// Domain-specific operations
	MarkAsAdopted(ctx context.Context, dogID uuid.UUID) error
}
