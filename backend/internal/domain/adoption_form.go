package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// AdoptionForm represents an adoption application
type AdoptionForm struct {
	AdoptionFormID         int       `json:"adoptionFormId"`
	AdopterID              uuid.UUID `json:"adopterId"`
	DogID                  uuid.UUID `json:"dogId"`
	SubmissionDate         time.Time `json:"submissionDate"`
	ProcessedByVolunteerID uuid.UUID `json:"processedByVolunteerId"`
	ProcessingDate         time.Time `json:"processingDate"`
	Status                 string    `json:"status"`
	RejectionReason        string    `json:"rejectionReason"`
}

// AdoptionFormRepository defines the interface for adoption form data access operations
type AdoptionFormRepository interface {
	GetAll(ctx context.Context) ([]AdoptionForm, error)
	GetPending(ctx context.Context) ([]AdoptionForm, error)
	GetByID(ctx context.Context, formID int) (*AdoptionForm, error)
	Create(ctx context.Context, form *AdoptionForm) error
	Update(ctx context.Context, form *AdoptionForm) error
	Delete(ctx context.Context, formID int) error

	// Domain-specific operations
	ApproveForm(ctx context.Context, formID int, volunteerID uuid.UUID) error
	RejectForm(ctx context.Context, formID int, volunteerID uuid.UUID, reason string) error
	CompleteForm(ctx context.Context, formID int, volunteerID uuid.UUID) error

	// TODO: Considering implementing
	// GetByAdopterID(ctx context.Context, adopterID uuid.UUID) ([]AdoptionForm, error)
	// GetByDogID(ctx context.Context, dogID uuid.UUID) ([]AdoptionForm, error)
}
