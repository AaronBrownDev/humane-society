package domain

import (
	"context"
	"github.com/google/uuid"
)

// ItemCatalog represents a catalog item
type ItemCatalog struct {
	ItemID          uuid.UUID `json:"itemId"`
	Name            string    `json:"name"`
	Category        string    `json:"category"`
	Description     string    `json:"description"`
	MinimumQuantity int       `json:"minimumQuantity"`
	IsActive        bool      `json:"isActive"`
}

// ItemCatalogRepository defines the interface for item catalog data access operations
type ItemCatalogRepository interface {
	GetAll(ctx context.Context) ([]ItemCatalog, error)
	GetByID(ctx context.Context, itemID uuid.UUID) (*ItemCatalog, error)
	Create(ctx context.Context, item *ItemCatalog) error
	Update(ctx context.Context, item *ItemCatalog) error
	Delete(ctx context.Context, itemID uuid.UUID) error

	// Domain-specific operations
	GetByCategory(ctx context.Context, category string) ([]ItemCatalog, error)
	GetActive(ctx context.Context) ([]ItemCatalog, error)
	SetActiveStatus(ctx context.Context, itemID uuid.UUID, isActive bool) error
}
