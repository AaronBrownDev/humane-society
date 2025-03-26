package repository

import (
	"database/sql"
	"github.com/AaronBrownDev/HumaneSociety/internal/models"
	"github.com/google/uuid"
)

type SurrenderRepository interface {
	// Surrender Form CRUD
	GetAllSurrenderForms() ([]models.SurrenderForm, error)
	GetSurrenderForm(formID int) (*models.SurrenderForm, error)
	InsertSurrenderForm(form *models.SurrenderForm) error
	UpdateSurrenderForm(form *models.SurrenderForm) error
	DeleteSurrenderForm(formID int) error

	// Surrender CRUD
	// TODO: Need to rethink surrenderer model
	GetAllSurrenders() ([]models.PetSurrenderer, error)
	GetSurrender(surrendererID uuid.UUID) (*models.PetSurrenderer, error)
	InsertSurrender() error
	UpdateSurrender() error
	DeleteSurrender() error
}

type SQLSurrenderRepository struct {
	db *sql.DB
}

func NewSQLSurrenderRepository(db *sql.DB) SurrenderRepository {
	return &SQLSurrenderRepository{
		db: db,
	}
}

// TODO: Implement functions
