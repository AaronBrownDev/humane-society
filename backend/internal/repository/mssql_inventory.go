package repository

import (
	"database/sql"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type InventoryRepository interface {
	// ItemCatalog CRUD
	// TODO: Look into names. ItemCatalog makes it confusing
	GetEntireItemCatalog() ([]domain.ItemCatalog, error)
	GetItemCatalog(itemID uuid.UUID) (*domain.ItemCatalog, error)
	CreateItemCatalog(item *domain.ItemCatalog) error
	UpdateItemCatalog(item *domain.ItemCatalog) error
	DeleteItemCatalog(itemID uuid.UUID) error

	// Supply CRUD
	GetSupplies() ([]domain.Supply, error)
	GetSupply(supplyID int) (*domain.Supply, error)
	AddSupply(supply *domain.Supply) error
	UpdateSupply(supply *domain.Supply) error
	DeleteSupply(supplyID int) error
}

type SQLInventoryRepository struct {
	db *sql.DB
}

func NewSQLInventoryRepository(db *sql.DB) InventoryRepository {
	return &SQLInventoryRepository{
		db: db,
	}
}

// TODO: implement functions
