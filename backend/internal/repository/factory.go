package repository

import (
	"database/sql"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"time"
)

// QueryTimeoutDuration default timeout for database operations
var QueryTimeoutDuration = time.Second * 5

// Storage is a factory for all repositories in the application
type Storage struct {
	Dogs           domain.DogRepository
	People         domain.PersonRepository
	Adopters       domain.AdopterRepository
	PetOwners      domain.PetOwnerRepository
	Veterinarians  domain.VeterinarianRepository
	Volunteers     domain.VolunteerRepository
	Surrenders     domain.SurrenderFormRepository
	Adoptions      domain.AdoptionFormRepository
	VolunteerForms domain.VolunteerFormRepository
	Medicines      domain.MedicineRepository
	Supplies       domain.SupplyRepository
	ItemCatalog    domain.ItemCatalogRepository
}

// NewMSSQLStorage creates a new storage with all repositories initialized
func NewMSSQLStorage(db *sql.DB) *Storage {
	conn := db

	return &Storage{
		Dogs:           NewMSSQLDog(conn),
		People:         NewMSSQLPerson(conn),
		Adopters:       NewMSSQLAdopter(conn),
		PetOwners:      NewMSSQLPetOwner(conn),
		Veterinarians:  NewMSSQLVeterinarian(conn),
		Volunteers:     NewMSSQLVolunteer(conn),
		Surrenders:     NewMSSQLSurrenderForm(conn),
		Adoptions:      NewMSSQLAdoptionForm(conn),
		VolunteerForms: NewMSSQLVolunteerForm(conn),
		Medicines:      NewMSSQLMedicine(conn),
		Supplies:       NewMSSQLSupply(conn),
		ItemCatalog:    NewMSSQLItemCatalog(conn),
	}
}
