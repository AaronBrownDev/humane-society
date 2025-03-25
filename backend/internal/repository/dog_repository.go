package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/models"
	"github.com/google/uuid"
)

// DogRepository CRUD Interface
type DogRepository interface {

	// Dog functions
	GetAllDogs() ([]models.Dog, error)
	GetAvailableDogs() ([]models.Dog, error)
	GetDogByID(dogID uuid.UUID) (*models.Dog, error)
	CreateDog(dog *models.Dog) error
	UpdateDog(dog *models.Dog) error
	DeleteDog(dogID uuid.UUID) error

	// Dog prescription functions
	GetDogPrescriptions(dogID uuid.UUID) ([]models.DogPrescription, error)
	AddDogPrescription(dogPrescription *models.DogPrescription) error
	UpdateDogPrescription(dogPrescription *models.DogPrescription) error
	RemoveDogPrescription(dogPrescriptionID uuid.UUID) error

	// Etc
	MarkAsAdopted(dogID uuid.UUID) error
	// GetMedicalHistory(dogID uuid.UUID) ([]models.Dog, error)
}

type SQLDogRepository struct {
	db *sql.DB
}

func NewDogRepository(db *sql.DB) DogRepository {
	return &SQLDogRepository{db}
}

// GetAllDogs returns a slice of all dogs
// returns error if query fails
func (r *SQLDogRepository) GetAllDogs() ([]models.Dog, error) {
	query := `SELECT DogID, Name, IntakeDate, EstimatedBirthDate, Breed, Sex, Color, CageNumber, IsAdopted FROM shelter.Dog`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dogs []models.Dog
	for rows.Next() {
		var dog models.Dog
		err = rows.Scan(
			&dog.DogID,
			&dog.Name,
			&dog.IntakeDate,
			&dog.EstimatedBirthDate,
			&dog.Breed,
			&dog.Sex,
			&dog.Color,
			&dog.CageNumber,
			&dog.IsAdopted,
		)
		if err != nil {
			return nil, err
		}

		dogs = append(dogs, dog)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return dogs, nil
}

// GetDogByID returns the Dog associated with the dogID
// returns error if query fails
func (r *SQLDogRepository) GetDogByID(dogID uuid.UUID) (*models.Dog, error) {
	query := `SELECT DogID, Name, IntakeDate, EstimatedBirthDate, Breed, Sex, Color, CageNumber, IsAdopted FROM shelter.Dog WHERE DogID = @p1` // Should eventually create sql procedure
	row := r.db.QueryRow(query, dogID)

	var dog models.Dog

	err := row.Scan(
		&dog.DogID,
		&dog.Name,
		&dog.IntakeDate,
		&dog.EstimatedBirthDate,
		&dog.Breed,
		&dog.Sex,
		&dog.Color,
		&dog.CageNumber,
		&dog.IsAdopted,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("dog not found")
		}
		return nil, err
	}

	return &dog, nil

}

// CreateDog inserts a new Dog record with the dog information
// returns error if query fails
func (r *SQLDogRepository) CreateDog(dog *models.Dog) error {
	query := `INSERT INTO shelter.Dog
				(DogID, Name, IntakeDate, EstimatedBirthDate, Breed, Sex, Color, CageNumber, IsAdopted)
				VALUES
				(@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9)` // create procedure eventually then insert here

	// creates new DogID if not given
	if dog.DogID == uuid.Nil {
		dog.DogID = uuid.New()
	}

	// executes query
	_, err := r.db.Exec(
		query,
		dog.DogID,
		dog.Name,
		dog.IntakeDate,
		dog.EstimatedBirthDate,
		dog.Breed,
		dog.Sex,
		dog.Color,
		dog.CageNumber,
		dog.IsAdopted,
	)
	if err != nil {
		return fmt.Errorf("error creating dog: %w", err)
	}

	return nil
}

// UpdateDog overwrites the Dog record that has the DogID in the given Dog struct with the information given
// returns error if query fails
func (r *SQLDogRepository) UpdateDog(dog *models.Dog) error {
	query := `UPDATE shelter.Dog 
				SET Name = @p1, IntakeDate = @p2, EstimatedBirthDate = @p3, Breed = @p4,
				    Sex = @p5, Color = @p6, CageNumber = @p7, IsAdopted = @p8
				WHERE DogID = @p9` // Might replace with procedure

	if dog.DogID == uuid.Nil {
		return errors.New("dog ID cannot be nil")
	}

	result, err := r.db.Exec(
		query,
		dog.Name,
		dog.IntakeDate,
		dog.EstimatedBirthDate,
		dog.Breed,
		dog.Sex,
		dog.Color,
		dog.CageNumber,
		dog.IsAdopted,
		dog.DogID,
	)
	if err != nil {
		return fmt.Errorf("error updating dog: %w", err)
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		if err != nil {
			return err
		}
		return errors.New("dog not found")
	}

	return nil
}

// DeleteDog deletes the dog record that has the dogID in the argument
// returns error if query fails
func (r *SQLDogRepository) DeleteDog(dogID uuid.UUID) error {
	query := `DELETE FROM shelter.Dog
              WHERE DogID = @p1` // Might replace with procedure

	if dogID == uuid.Nil {
		return errors.New("dog ID cannot be nil")
	}

	result, err := r.db.Exec(query, dogID)
	if err != nil {
		return fmt.Errorf("error deleting dog: %w", err)
	} else if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		if err != nil {
			return err
		}
		return errors.New("dog not found")
	}

	return nil
}

// GetAvailableDogs returns a slice of all dogs that are available to be adopted
// returns error if query fails
func (r *SQLDogRepository) GetAvailableDogs() ([]models.Dog, error) {
	query := `SELECT DogID, Name, IntakeDate, EstimatedBirthDate, Breed, Sex, Color, CageNumber, IsAdopted FROM shelter.AvailableDogs`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dogs []models.Dog
	for rows.Next() {
		var dog models.Dog

		err = rows.Scan(
			&dog.DogID,
			&dog.Name,
			&dog.IntakeDate,
			&dog.EstimatedBirthDate,
			&dog.Breed,
			&dog.Sex,
			&dog.Color,
			&dog.CageNumber,
			&dog.IsAdopted,
		)
		if err != nil {
			return nil, err
		}

		dogs = append(dogs, dog)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return dogs, nil
}

// Dog prescription functions

// GetDogPrescriptions returns a slice of DogPrescription that is associated with the dog argument
// returns error if query fails
func (r *SQLDogRepository) GetDogPrescriptions(dogID uuid.UUID) ([]models.DogPrescription, error) {
	query := `SELECT PrescriptionID, DogID, MedicineID, Dosage, Frequency, StartDate, EndDate, Notes, VetPrescriberID FROM medical.DogPrescription`

	rows, err := r.db.Query(query, dogID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prescriptions []models.DogPrescription
	for rows.Next() {
		var prescription models.DogPrescription
		err = rows.Scan(
			&prescription.PrescriptionID,
			&prescription.DogID,
			&prescription.MedicineID,
			&prescription.Dosage,
			&prescription.Frequency,
			&prescription.StartDate,
			&prescription.EndDate,
			&prescription.Notes,
			&prescription.VetPrescriberID,
		)
		if err != nil {
			return nil, err
		}

		prescriptions = append(prescriptions, prescription)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return prescriptions, nil
}

// AddDogPrescription returns error if query fails
func (r *SQLDogRepository) AddDogPrescription(dogPrescription *models.DogPrescription) error {

	return nil
}

// UpdateDogPrescription returns error if query fails
func (r *SQLDogRepository) UpdateDogPrescription(dogPrescription *models.DogPrescription) error {

	return nil
}

// RemoveDogPrescription returns error if query fails
func (r *SQLDogRepository) RemoveDogPrescription(dogPrescriptionID uuid.UUID) error {

	return nil
}

// Etc

// MarkAsAdopted returns error if query fails
func (r *SQLDogRepository) MarkAsAdopted(dogID uuid.UUID) error {

	return nil
}
