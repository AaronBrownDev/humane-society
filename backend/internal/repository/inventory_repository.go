package repository

import (
	"database/sql"
	"github.com/AaronBrownDev/HumaneSociety/internal/models"
	"github.com/google/uuid"
)

type InventoryRepository interface {
	// ItemCatalog CRUD
	// TODO: Look into names. ItemCatalog makes it confusing
	GetEntireItemCatalog() ([]models.ItemCatalog, error)
	GetItemCatalog(itemID uuid.UUID) (*models.ItemCatalog, error)
	CreateItemCatalog(item *models.ItemCatalog) error
	UpdateItemCatalog(item *models.ItemCatalog) error
	DeleteItemCatalog(itemID uuid.UUID) error

	// Supply CRUD
	GetSupplies() ([]models.Supply, error)
	GetSupply(supplyID int) (*models.Supply, error)
	AddSupply(supply *models.Supply) error
	UpdateSupply(supply *models.Supply) error
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
