package repository

import (
	"database/sql"
	"github.com/AaronBrownDev/HumaneSociety/internal/models"
	"github.com/google/uuid"
)

type AdoptionRepository interface {
	// Adoption Form CRUD
	GetAllAdoptionForms() ([]models.AdoptionForm, error)
	GetAdoptionForm(formID int) (*models.AdoptionForm, error)
	InsertAdoptionForm(form *models.AdoptionForm) error
	UpdateAdoptionForm(form *models.AdoptionForm) error
	DeleteAdoptionForm(formID int) error

	// Adopter CRUD
	GetAllAdopters() ([]models.Adopter, error)
	GetAdopter(adopterID uuid.UUID) (*models.Adopter, error)
	InsertAdopter(adopter *models.Adopter) error
	UpdateAdopter(adopter *models.Adopter) error
	DeleteAdopter(adopterID uuid.UUID) error
}

type SQLAdoptionRepository struct {
	db *sql.DB
}

func NewSQLAdoptionRepository(db *sql.DB) AdoptionRepository {
	return &SQLAdoptionRepository{
		db: db,
	}
}

// TODO: Implement functions
