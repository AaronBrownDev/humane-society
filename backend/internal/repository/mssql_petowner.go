package repository

import (
	"database/sql"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type PetOwnerRepository interface {
	GetAll() ([]domain.PetOwner, error)
	GetByID(PetOwnerID uuid.UUID) (*domain.PetOwner, error)
}

type SQLPetOwnerRepository struct {
	db *sql.DB
}

func NewPetOwnerRepository(db *sql.DB) *SQLPetOwnerRepository {
	return &SQLPetOwnerRepository{db: db}
}

func (r *SQLPetOwnerRepository) GetAll() ([]domain.PetOwner, error) {
	query := `SELECT PetOwnerID, VeterinarianID, HasSterilizedPets, HasVaccinatedPets, UsesVetHeartWormPrevention FROM people.PetOwner`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var petOwners []domain.PetOwner
	for rows.Next() {
		var petOwner domain.PetOwner
		err := rows.Scan(
			&petOwner.PetOwnerID,
			&petOwner.VeterinarianID,
			&petOwner.HasSterilizedPets,
			&petOwner.HasVaccinatedPets,
			&petOwner.UsesVetHeartWormPrevention,
		)
		if err != nil {
			return nil, err
		}
		petOwners = append(petOwners, petOwner)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return petOwners, nil
}
