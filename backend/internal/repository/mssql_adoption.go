package repository

import (
	"database/sql"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type AdoptionRepository struct {
	adoptionFormRepository interface {
		Repository[domain.AdoptionForm, int]
	}
	adopterRepository interface {
		Repository[domain.Adopter, uuid.UUID]
	}
	db *sql.DB
}

// TODO: Implement functions
