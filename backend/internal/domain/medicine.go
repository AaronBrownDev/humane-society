package domain

import (
	"context"
)

// Medicine represents a medication used at the shelter
type Medicine struct {
	MedicineID   int    `json:"medicineId"`
	Name         string `json:"name"`
	Manufacturer string `json:"manufacturer"`
	Description  string `json:"description"`
	DosageUnit   string `json:"dosageUnit"`
}

// MedicineRepository defines the interface for medicine data access operations
type MedicineRepository interface {
	GetAll(ctx context.Context) ([]Medicine, error)
	GetByID(ctx context.Context, medicineID int) (*Medicine, error)
	Create(ctx context.Context, medicine *Medicine) error
	Update(ctx context.Context, medicine *Medicine) error
	Delete(ctx context.Context, medicineID int) error
}
