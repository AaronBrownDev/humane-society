package repository

// CRUD Interface
type AdoptionRepository interface {
	// Adoption Form CRUD
	GetAdoptionForm() error
	InsertAdoptionForm() error
	UpdateAdoptionForm() error
	DeleteAdoptionForm() error
}
