package repository

import (
	"database/sql"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type SurrenderRepository interface {
	// Surrender Form CRUD
	GetAllSurrenderForms() ([]domain.SurrenderForm, error)
	GetSurrenderForm(formID int) (*domain.SurrenderForm, error)
	InsertSurrenderForm(form *domain.SurrenderForm) error
	UpdateSurrenderForm(form *domain.SurrenderForm) error
	DeleteSurrenderForm(formID int) error

	// Surrender CRUD
	// TODO: Need to rethink surrenderer model
	GetAllSurrenders() ([]domain.PetSurrenderer, error)
	GetSurrender(surrendererID uuid.UUID) (*domain.PetSurrenderer, error)
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
