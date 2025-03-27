package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
	"time"
)

type mssqlDogPrescriptionRepository struct {
	conn Connection
}

func NewMSSQLDogPrescription(conn Connection) domain.DogPrescriptionRepository {
	return &mssqlDogPrescriptionRepository{conn: conn}
}

// fetch is a helper function to retrieve multiple dog prescription records
func (r *mssqlDogPrescriptionRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.DogPrescription, error) {
	rows, err := r.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prescriptions []domain.DogPrescription
	for rows.Next() {
		var prescription domain.DogPrescription
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

// GetAll retrieves all dog prescriptions from the database
func (r *mssqlDogPrescriptionRepository) GetAll(ctx context.Context) ([]domain.DogPrescription, error) {
	query := `SELECT PrescriptionID, DogID, MedicineID, Dosage, Frequency, StartDate, EndDate, Notes, VetPrescriberID 
              FROM medical.DogPrescription`

	return r.fetch(ctx, query)
}

// GetByID retrieves a specific dog prescription by its ID
func (r *mssqlDogPrescriptionRepository) GetByID(ctx context.Context, prescriptionID int) (*domain.DogPrescription, error) {
	query := `SELECT PrescriptionID, DogID, MedicineID, Dosage, Frequency, StartDate, EndDate, Notes, VetPrescriberID 
              FROM medical.DogPrescription 
              WHERE PrescriptionID = @p1`

	var prescription domain.DogPrescription
	err := r.conn.QueryRowContext(ctx, query, prescriptionID).Scan(
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
		if errors.Is(err, errors.New("sql: no rows in result set")) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	return &prescription, nil
}

// GetByDogID retrieves all prescriptions for a specific dog
func (r *mssqlDogPrescriptionRepository) GetByDogID(ctx context.Context, dogID uuid.UUID) ([]domain.DogPrescription, error) {
	query := `SELECT PrescriptionID, DogID, MedicineID, Dosage, Frequency, StartDate, EndDate, Notes, VetPrescriberID 
              FROM medical.DogPrescription 
              WHERE DogID = @p1`

	return r.fetch(ctx, query, dogID)
}

// Create inserts a new dog prescription record into the database
func (r *mssqlDogPrescriptionRepository) Create(ctx context.Context, prescription *domain.DogPrescription) error {
	query := `INSERT INTO medical.DogPrescription
              (DogID, MedicineID, Dosage, Frequency, StartDate, EndDate, Notes, VetPrescriberID)
              VALUES
              (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8);
              SELECT SCOPE_IDENTITY()`

	// Execute the query and get the generated ID
	err := r.conn.QueryRowContext(
		ctx,
		query,
		prescription.DogID,
		prescription.MedicineID,
		prescription.Dosage,
		prescription.Frequency,
		prescription.StartDate,
		prescription.EndDate,
		prescription.Notes,
		prescription.VetPrescriberID,
	).Scan(&prescription.PrescriptionID)

	if err != nil {
		return fmt.Errorf("error creating dog prescription: %w", err)
	}

	return nil
}

// Update modifies an existing dog prescription record in the database
func (r *mssqlDogPrescriptionRepository) Update(ctx context.Context, prescription *domain.DogPrescription) error {
	query := `UPDATE medical.DogPrescription
              SET DogID = @p1, MedicineID = @p2, Dosage = @p3, Frequency = @p4, 
                  StartDate = @p5, EndDate = @p6, Notes = @p7, VetPrescriberID = @p8
              WHERE PrescriptionID = @p9`

	if prescription.PrescriptionID <= 0 {
		return domain.ErrInvalidInput
	}

	result, err := r.conn.ExecContext(
		ctx,
		query,
		prescription.DogID,
		prescription.MedicineID,
		prescription.Dosage,
		prescription.Frequency,
		prescription.StartDate,
		prescription.EndDate,
		prescription.Notes,
		prescription.VetPrescriberID,
		prescription.PrescriptionID,
	)
	if err != nil {
		return fmt.Errorf("error updating dog prescription: %w", err)
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

// Delete removes a dog prescription record from the database
func (r *mssqlDogPrescriptionRepository) Delete(ctx context.Context, prescriptionID int) error {
	query := `DELETE FROM medical.DogPrescription
              WHERE PrescriptionID = @p1`

	if prescriptionID <= 0 {
		return domain.ErrInvalidInput
	}

	result, err := r.conn.ExecContext(ctx, query, prescriptionID)
	if err != nil {
		return fmt.Errorf("error deleting dog prescription: %w", err)
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

// GetActivePrescriptions retrieves all active prescriptions for a specific dog
// Active prescriptions are those where the current date is between StartDate and EndDate
// or EndDate is not set (ongoing medication)
func (r *mssqlDogPrescriptionRepository) GetActivePrescriptions(ctx context.Context, dogID uuid.UUID) ([]domain.DogPrescription, error) {
	query := `SELECT PrescriptionID, DogID, MedicineID, Dosage, Frequency, StartDate, EndDate, Notes, VetPrescriberID 
              FROM medical.DogPrescription 
              WHERE DogID = @p1 
              AND (
                  (StartDate <= @p2 AND (EndDate IS NULL OR EndDate >= @p2))
              )`

	currentTime := time.Now()
	return r.fetch(ctx, query, dogID, currentTime)
}
