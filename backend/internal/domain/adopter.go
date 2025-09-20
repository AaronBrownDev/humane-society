package domain

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

const (
	HomeStatusPending  = "Pending"
	HomeStatusApproved = "Approved"
	HomeStatusRejected = "Rejected"
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

var (
	ErrInvalidHomeStatus = errors.New("invalid home status")
)

func (a *Adopter) Validate() error {
	// Check if nested person is valid
	if err := a.Person.Validate(); err != nil {
		return err
	}

	// Set a.HomeStatus to Pending if empty
	if a.HomeStatus == "" {
		a.HomeStatus = HomeStatusPending
	}

	if !isValidHomeStatus(a.HomeStatus) {
		return ErrInvalidHomeStatus
	}

	return nil
}

func isValidHomeStatus(homeStatus string) bool {
	switch homeStatus {
	case HomeStatusPending, HomeStatusApproved, HomeStatusRejected:
		return true
	default:
		return false
	}
}
