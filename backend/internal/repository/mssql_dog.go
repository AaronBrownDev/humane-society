package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlDogRepository struct {
	conn Connection
}

func NewMSSQLDog(conn Connection) domain.DogRepository {
	return &mssqlDogRepository{conn: conn}
}

// fetch is a helper function to retrieve multiple dog records
func (r *mssqlDogRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.Dog, error) {
	rows, err := r.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dogs []domain.Dog
	for rows.Next() {
		var dog domain.Dog
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

// GetAll retrieves all dogs from the database
func (r *mssqlDogRepository) GetAll(ctx context.Context) ([]domain.Dog, error) {
	query := `SELECT DogID, Name, IntakeDate, EstimatedBirthDate, Breed, Sex, Color, CageNumber, IsAdopted 
              FROM shelter.Dog`

	return r.fetch(ctx, query)
}

// GetAvailable retrieves all dogs that are available for adoption
func (r *mssqlDogRepository) GetAvailable(ctx context.Context) ([]domain.Dog, error) {
	query := `SELECT DogID, Name, IntakeDate, EstimatedBirthDate, Breed, Sex, Color, CageNumber
              FROM shelter.AvailableDogs`

	dogs, err := r.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	// Set IsAdopted to false for all dogs from the view
	for i := range dogs {
		dogs[i].IsAdopted = false
	}

	return dogs, nil
}

// GetByID retrieves a specific dog by its unique identifier
func (r *mssqlDogRepository) GetByID(ctx context.Context, dogID uuid.UUID) (*domain.Dog, error) {
	query := `SELECT DogID, Name, IntakeDate, EstimatedBirthDate, Breed, Sex, Color, CageNumber, IsAdopted 
              FROM shelter.Dog 
              WHERE DogID = @p1`

	var dog domain.Dog
	err := r.conn.QueryRowContext(ctx, query, dogID).Scan(
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
		if errors.Is(err, errors.New("sql: no rows in result set")) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &dog, nil
}

// Create inserts a new dog record into the database
func (r *mssqlDogRepository) Create(ctx context.Context, dog *domain.Dog) error {
	query := `INSERT INTO shelter.Dog
              (DogID, Name, IntakeDate, EstimatedBirthDate, Breed, Sex, Color, CageNumber, IsAdopted)
              VALUES
              (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8, @p9)`

	// Generate a new UUID if none is provided
	if dog.DogID == uuid.Nil {
		dog.DogID = uuid.New()
	}

	result, err := r.conn.ExecContext(
		ctx,
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

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrDatabaseError
	}

	return nil
}

// Update modifies an existing dog record in the database
func (r *mssqlDogRepository) Update(ctx context.Context, dog *domain.Dog) error {
	query := `UPDATE shelter.Dog 
              SET Name = @p1, IntakeDate = @p2, EstimatedBirthDate = @p3, Breed = @p4,
                  Sex = @p5, Color = @p6, CageNumber = @p7, IsAdopted = @p8
              WHERE DogID = @p9`

	if dog.DogID == uuid.Nil {
		return domain.ErrInvalidInput
	}

	result, err := r.conn.ExecContext(
		ctx,
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

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}

// Delete removes a dog record from the database
func (r *mssqlDogRepository) Delete(ctx context.Context, dogID uuid.UUID) error {
	query := `DELETE FROM shelter.Dog
              WHERE DogID = @p1`

	if dogID == uuid.Nil {
		return domain.ErrInvalidInput
	}

	result, err := r.conn.ExecContext(ctx, query, dogID)
	if err != nil {
		return fmt.Errorf("error deleting dog: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}

// MarkAsAdopted updates a dog's adoption status to adopted (IsAdopted = true)
func (r *mssqlDogRepository) MarkAsAdopted(ctx context.Context, dogID uuid.UUID) error {
	query := `UPDATE shelter.Dog
              SET IsAdopted = 1
              WHERE DogID = @p1`

	result, err := r.conn.ExecContext(ctx, query, dogID)
	if err != nil {
		return fmt.Errorf("error updating dog adoption status: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrNotFound
	}

	return nil
}
