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
	Dogs               domain.DogRepository
	DogPrescriptions   domain.DogPrescriptionRepository
	People             domain.PersonRepository
	Adopters           domain.AdopterRepository
	PetOwners          domain.PetOwnerRepository
	PetOwnerPets       domain.PetOwnerPetRepository
	PetSurrenderers    domain.PetSurrendererRepository
	Veterinarians      domain.VeterinarianRepository
	Volunteers         domain.VolunteerRepository
	VolunteerSchedules domain.VolunteerScheduleRepository
	Surrenders         domain.SurrenderFormRepository
	Adoptions          domain.AdoptionFormRepository
	VolunteerForms     domain.VolunteerFormRepository
	Medicines          domain.MedicineRepository
	Supplies           domain.SupplyRepository
	ItemCatalog        domain.ItemCatalogRepository
	UserAccounts       domain.UserAccountRepository
	UserRoles          domain.UserRoleRepository
	Roles              domain.RoleRepository
	RefreshTokens      domain.RefreshTokenRepository
}

// NewMSSQLStorage creates a new storage with all repositories initialized
func NewMSSQLStorage(db *sql.DB) *Storage {
	conn := db

	return &Storage{
		Dogs:               NewMSSQLDog(conn),
		DogPrescriptions:   NewMSSQLDogPrescription(conn),
		People:             NewMSSQLPerson(conn),
		Adopters:           NewMSSQLAdopter(conn),
		PetOwners:          NewMSSQLPetOwner(conn),
		PetOwnerPets:       NewMSSQLPetOwnerPet(conn),
		PetSurrenderers:    NewMSSQLSurrenderer(conn),
		Veterinarians:      NewMSSQLVeterinarian(conn),
		Volunteers:         NewMSSQLVolunteer(conn),
		VolunteerSchedules: NewMSSQLVolunteerSchedule(conn),
		Surrenders:         NewMSSQLSurrenderForm(conn),
		Adoptions:          NewMSSQLAdoptionForm(conn),
		VolunteerForms:     NewMSSQLVolunteerForm(conn),
		Medicines:          NewMSSQLMedicine(conn),
		Supplies:           NewMSSQLSupply(conn),
		ItemCatalog:        NewMSSQLItemCatalog(conn),
		UserAccounts:       NewMSSQLUserAccount(conn),
		UserRoles:          NewMSSQLUserRole(conn),
		Roles:              NewMSSQLRole(conn),
		RefreshTokens:      NewMSSQLRefreshToken(conn),
	}
}
