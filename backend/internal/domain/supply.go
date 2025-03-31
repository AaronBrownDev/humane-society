package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// Supply represents an inventory item
type Supply struct {
	SupplyID        int       `json:"supplyId"`
	ItemID          uuid.UUID `json:"itemId"`
	Quantity        int       `json:"quantity"`
	StorageLocation string    `json:"storageLocation"`
	ExpirationDate  time.Time `json:"expirationDate"`
	BatchNumber     string    `json:"batchNumber"`
	AcquisitionDate time.Time `json:"acquisitionDate"`
}

// SupplyRepository defines the interface for supply data access operations
type SupplyRepository interface {
	GetAll(ctx context.Context) ([]Supply, error)
	GetByID(ctx context.Context, supplyID int) (*Supply, error)
	GetByItemID(ctx context.Context, itemID uuid.UUID) ([]Supply, error)
	Create(ctx context.Context, supply *Supply) error
	Update(ctx context.Context, supply *Supply) error
	Delete(ctx context.Context, supplyID int) error

	// Domain-specific operations
	UpdateQuantity(ctx context.Context, supplyID int, quantity int) error
	GetExpiringSoon(ctx context.Context, days int) ([]Supply, error)
	GetLowStock(ctx context.Context) ([]Supply, error)
}
