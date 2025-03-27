package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// DogPrescription represents a medication prescription for a shelter dog
type DogPrescription struct {
	PrescriptionID  int       `json:"prescriptionId"`
	DogID           uuid.UUID `json:"dogId"`
	MedicineID      int       `json:"medicineId"`
	Dosage          float64   `json:"dosage"`
	Frequency       string    `json:"frequency"`
	StartDate       time.Time `json:"startDate"`
	EndDate         time.Time `json:"endDate"`
	Notes           string    `json:"notes"`
	VetPrescriberID uuid.UUID `json:"vetPrescriberId"`
}

// DogPrescriptionRepository defines the interface for dog prescription data access operations
type DogPrescriptionRepository interface {
	GetAll(ctx context.Context) ([]DogPrescription, error)
	GetByID(ctx context.Context, prescriptionID int) (*DogPrescription, error)
	GetByDogID(ctx context.Context, dogID uuid.UUID) ([]DogPrescription, error)
	Create(ctx context.Context, prescription *DogPrescription) error
	Update(ctx context.Context, prescription *DogPrescription) error
	Delete(ctx context.Context, prescriptionID int) error

	// Domain-specific operations
	GetActivePrescriptions(ctx context.Context, dogID uuid.UUID) ([]DogPrescription, error)
}
