package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/models"
	"github.com/google/uuid"
)

// DogRepository defines the interface for dog and dog prescription data access operations.
// Provides methods for managing dogs and their medical prescriptions.
type DogRepository interface {

	// Dog CRUD functions
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
	RemoveDogPrescription(dogPrescriptionID int) error

	// Dog etc. functions
	MarkAsAdopted(dogID uuid.UUID) error
	// GetMedicalHistory(dogID uuid.UUID) ([]models.Dog, error)
}

// SQLDogRepository implements the DogRepository interface using SQL database access.
type SQLDogRepository struct {
	db *sql.DB
}

// NewDogRepository creates a new SQLDogRepository instance that implements the DogRepository interface.
func NewDogRepository(db *sql.DB) DogRepository {
	return &SQLDogRepository{db}
}

// GetAllDogs retrieves all dogs from the database.
// Returns a slice of Dog models or an error if the database operation fails.
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

// GetDogByID retrieves a specific dog by its unique identifier.
// Returns the dog if found or an error if the dog doesn't exist or if the query fails.
func (r *SQLDogRepository) GetDogByID(dogID uuid.UUID) (*models.Dog, error) {
	// TODO: Create SQL procedure for this query
	query := `SELECT DogID, Name, IntakeDate, EstimatedBirthDate, Breed, Sex, Color, CageNumber, IsAdopted FROM shelter.Dog WHERE DogID = @p1`
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

// CreateDog inserts a new dog record into the database.
// Generates a new UUID if none is provided in the dog model.
// Returns an error if the database operation fails.
func (r *SQLDogRepository) CreateDog(dog *models.Dog) error {
	// TODO: Create SQL procedure for this query
	query := `INSERT INTO shelter.Dog
				(DogID, Name, IntakeDate, EstimatedBirthDate, Breed, Sex, Color, CageNumber, IsAdopted)
				VALUES
				(@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9)`

	// Generate a new UUID if none is provided
	// TODO: Not sure if new UUID can cause errors have to look into
	if dog.DogID == uuid.Nil {
		dog.DogID = uuid.New()
	}

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

// UpdateDog modifies an existing dog record in the database.
// Returns an error if the dog isn't found or if the database operation fails.
func (r *SQLDogRepository) UpdateDog(dog *models.Dog) error {
	// TODO: Create SQL procedure for this query
	query := `UPDATE shelter.Dog 
				SET Name = @p1, IntakeDate = @p2, EstimatedBirthDate = @p3, Breed = @p4,
				    Sex = @p5, Color = @p6, CageNumber = @p7, IsAdopted = @p8
				WHERE DogID = @p9`

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

// DeleteDog removes a dog record from the database.
// Returns an error if the dog isn't found or if the database operation fails.
func (r *SQLDogRepository) DeleteDog(dogID uuid.UUID) error {
	// TODO: Create SQL procedure for this query
	query := `DELETE FROM shelter.Dog
              WHERE DogID = @p1`

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

// GetAvailableDogs retrieves all dogs that are available for adoption.
// Returns a slice of available dogs or an error if the database operation fails.
func (r *SQLDogRepository) GetAvailableDogs() ([]models.Dog, error) {
	// TODO: Create SQL procedure for this query
	query := `SELECT DogID, Name, IntakeDate, EstimatedBirthDate, Breed, Sex, Color, CageNumber FROM shelter.AvailableDogs`

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
		)
		if err != nil {
			return nil, err
		}

		// Available dogs will always be not adopted
		dog.IsAdopted = false

		dogs = append(dogs, dog)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return dogs, nil
}

// GetDogPrescriptions retrieves all prescription records for a specific dog.
// Returns a slice of DogPrescription models or an error if the database operation fails.
func (r *SQLDogRepository) GetDogPrescriptions(dogID uuid.UUID) ([]models.DogPrescription, error) {
	// TODO: Create SQL procedure for this query
	query := `SELECT PrescriptionID, DogID, MedicineID, Dosage, Frequency, StartDate, EndDate, Notes, VetPrescriberID FROM medical.DogPrescription
				WHERE DogID = @p1`

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

// AddDogPrescription creates a new prescription record for a dog.
// Returns an error if the insert operation fails or if no rows are affected.
func (r *SQLDogRepository) AddDogPrescription(dogPrescription *models.DogPrescription) error {
	// TODO: Create SQL procedure for this query
	query := `INSERT INTO medical.DogPrescription
				(DogID, MedicineID, Dosage, Frequency, StartDate, EndDate, Notes, VetPrescriberID)
				VALUES
				(@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8)`

	result, err := r.db.Exec(
		query,
		dogPrescription.DogID,
		dogPrescription.MedicineID,
		dogPrescription.Dosage,
		dogPrescription.Frequency,
		dogPrescription.StartDate,
		dogPrescription.EndDate,
		dogPrescription.Notes,
		dogPrescription.VetPrescriberID,
	)
	if err != nil {
		return fmt.Errorf("error adding dog prescription: %w", err)
	}

	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		if err != nil {
			return err
		}
		return errors.New("dog prescription insert failed")
	}

	return nil
}

// UpdateDogPrescription modifies an existing prescription record.
// Returns an error if the prescription isn't found or if the database operation fails.
func (r *SQLDogRepository) UpdateDogPrescription(dogPrescription *models.DogPrescription) error {
	// TODO: Create SQL procedure for this query
	query := `UPDATE medical.DogPrescription
				SET DogID = @p2, MedicineID = @p3, Dosage = @p4, Frequency = @p5, StartDate = @p6, EndDate = @p7, Notes = @p8, VetPrescriberID = @p9
				WHERE PrescriptionID = @p1`

	result, err := r.db.Exec(query,
		dogPrescription.PrescriptionID,
		dogPrescription.DogID,
		dogPrescription.MedicineID,
		dogPrescription.Dosage,
		dogPrescription.Frequency,
		dogPrescription.StartDate,
		dogPrescription.EndDate,
		dogPrescription.Notes,
		dogPrescription.VetPrescriberID,
	)
	if err != nil {
		return fmt.Errorf("error updating dog prescription: %w", err)
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		if err != nil {
			return err
		}
		return errors.New("dog prescription not found")
	}

	return nil
}

// RemoveDogPrescription deletes a prescription record from the database.
// Returns an error if the prescription isn't found or if the database operation fails.
func (r *SQLDogRepository) RemoveDogPrescription(dogPrescriptionID int) error {
	// TODO: Create SQL procedure for this query
	query := `DELETE FROM medical.DogPrescription
				WHERE PrescriptionID = @p1`

	result, err := r.db.Exec(query, dogPrescriptionID)
	if err != nil {
		return fmt.Errorf("error removing dog prescription: %w", err)
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		if err != nil {
			return err
		}
		return errors.New("dog prescription not found")
	}

	return nil
}

// MarkAsAdopted updates a dog's adoption status to adopted (IsAdopted = true).
// Returns an error if the dog isn't found or if the database operation fails.
func (r *SQLDogRepository) MarkAsAdopted(dogID uuid.UUID) error {
	// TODO: Create SQL procedure for this query
	query := `UPDATE shelter.Dog
				SET IsAdopted = 1
				WHERE DogID = @p1`

	result, err := r.db.Exec(query, dogID)
	if err != nil {
		return fmt.Errorf("error updating dog adoption status: %w", err)
	}
	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		if err != nil {
			return err
		}
		return errors.New("dog not found")
	}

	return nil
}
