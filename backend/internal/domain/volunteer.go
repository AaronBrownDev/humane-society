package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// Volunteer represents a shelter volunteer
type Volunteer struct {
	VolunteerID           uuid.UUID `json:"volunteerId"`
	Person                Person    `json:"person"`
	VolunteerPosition     string    `json:"volunteerPosition"`
	StartDate             time.Time `json:"startDate"`
	EndDate               time.Time `json:"endDate"`
	EmergencyContactName  string    `json:"emergencyContactName"`
	EmergencyContactPhone string    `json:"emergencyContactPhone"`
	IsActive              bool      `json:"isActive"`
}

// VolunteerRepository defines the interface for volunteer data access operations
type VolunteerRepository interface {
	GetAll(ctx context.Context) ([]Volunteer, error)
	GetActive(ctx context.Context) ([]Volunteer, error)
	GetByID(ctx context.Context, volunteerID uuid.UUID) (*Volunteer, error)
	Create(ctx context.Context, volunteer *Volunteer) error
	Update(ctx context.Context, volunteer *Volunteer) error
	Delete(ctx context.Context, volunteerID uuid.UUID) error

	// Domain-specific operations
	SetActiveStatus(ctx context.Context, volunteerID uuid.UUID, isActive bool) error
	// TODO: Look into adding constraints for this
	// GetByPosition(ctx context.Context, position string) ([]Volunteer, error)
}
