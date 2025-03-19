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
	GetAll() ([]models.Dog, error)
	GetByID(dogID uuid.UUID) (*models.Dog, error)
	Create(dog *models.Dog) error
	Update(dog *models.Dog) error
	Delete(dogID uuid.UUID) error
	GetAvailable() ([]models.Dog, error)
}

type SQLDogRepository struct {
	db *sql.DB
}

func NewDogRepository(db *sql.DB) DogRepository {
	return &SQLDogRepository{db}
}

func (r *SQLDogRepository) GetAll() ([]models.Dog, error) {
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

func (r *SQLDogRepository) GetByID(dogID uuid.UUID) (*models.Dog, error) {
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

func (r *SQLDogRepository) Create(dog *models.Dog) error {
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

func (r *SQLDogRepository) Update(dog *models.Dog) error {
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

func (r *SQLDogRepository) Delete(dogID uuid.UUID) error {
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

func (r *SQLDogRepository) GetAvailable() ([]models.Dog, error) {
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
