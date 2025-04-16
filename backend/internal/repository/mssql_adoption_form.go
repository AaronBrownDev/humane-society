package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
	"time"
)

type mssqlAdoptionFormRepository struct {
	conn Connection
}

func NewMSSQLAdoptionForm(conn Connection) domain.AdoptionFormRepository {
	return &mssqlAdoptionFormRepository{conn: conn}
}

// fetch is a helper function to retrieve multiple adoption form records
func (r *mssqlAdoptionFormRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.AdoptionForm, error) {
	rows, err := r.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var forms []domain.AdoptionForm
	for rows.Next() {
		var form domain.AdoptionForm
		var procVolID sql.NullString
		var procDate sql.NullTime
		var rejReason sql.NullString

		err = rows.Scan(
			&form.AdoptionFormID,
			&form.AdopterID,
			&form.DogID,
			&form.SubmissionDate,
			&procVolID,
			&procDate,
			&form.Status,
			&rejReason,
		)
		if err != nil {
			return nil, err
		}

		// Handle NULL values
		if procVolID.Valid {
			form.ProcessedByVolunteerID, err = uuid.Parse(procVolID.String)
			if err != nil {
				return nil, fmt.Errorf("error parsing ProcessedByVolunteerID: %w", err)
			}

			form.ProcessedByVolunteerID, err = SwapUUIDFormat(form.ProcessedByVolunteerID)
			if err != nil {
				return nil, fmt.Errorf("error converting ProcessedByVolunteerID UUID: %w", err)
			}
		}

		if procDate.Valid {
			form.ProcessingDate = procDate.Time
		}

		if rejReason.Valid {
			form.RejectionReason = rejReason.String
		}

		// Convert UUIDs for SQL Server format
		form.AdopterID, err = SwapUUIDFormat(form.AdopterID)
		if err != nil {
			return nil, fmt.Errorf("error converting AdopterID UUID: %w", err)
		}

		form.DogID, err = SwapUUIDFormat(form.DogID)
		if err != nil {
			return nil, fmt.Errorf("error converting DogID UUID: %w", err)
		}

		forms = append(forms, form)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return forms, nil
}

// GetAll retrieves all adoption forms from the database
func (r *mssqlAdoptionFormRepository) GetAll(ctx context.Context) ([]domain.AdoptionForm, error) {
	query := `SELECT AdoptionFormID, AdopterID, DogID, SubmissionDate, ProcessedByVolunteerID, ProcessingDate, Status, RejectionReason 
              FROM shelter.AdoptionForm`

	return r.fetch(ctx, query)
}

// GetPending retrieves all pending adoption forms
func (r *mssqlAdoptionFormRepository) GetPending(ctx context.Context) ([]domain.AdoptionForm, error) {
	query := `SELECT AdoptionFormID, AdopterID, DogID, SubmissionDate, ProcessedByVolunteerID, ProcessingDate, Status, RejectionReason 
              FROM shelter.AdoptionForm 
              WHERE Status = 'Pending'`

	return r.fetch(ctx, query)
}

// GetByID retrieves a specific adoption form by its ID
func (r *mssqlAdoptionFormRepository) GetByID(ctx context.Context, formID int) (*domain.AdoptionForm, error) {
	query := `SELECT AdoptionFormID, AdopterID, DogID, SubmissionDate, ProcessedByVolunteerID, ProcessingDate, Status, RejectionReason 
              FROM shelter.AdoptionForm 
              WHERE AdoptionFormID = @p1`

	var form domain.AdoptionForm
	var procVolID sql.NullString
	var procDate sql.NullTime
	var rejReason sql.NullString

	err := r.conn.QueryRowContext(ctx, query, formID).Scan(
		&form.AdoptionFormID,
		&form.AdopterID,
		&form.DogID,
		&form.SubmissionDate,
		&procVolID,
		&procDate,
		&form.Status,
		&rejReason,
	)

	if err != nil {
		if errors.Is(err, errors.New("sql: no rows in result set")) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	// Handle NULL values
	if procVolID.Valid {
		form.ProcessedByVolunteerID, err = uuid.Parse(procVolID.String)
		if err != nil {
			return nil, fmt.Errorf("error parsing ProcessedByVolunteerID: %w", err)
		}

		form.ProcessedByVolunteerID, err = SwapUUIDFormat(form.ProcessedByVolunteerID)
		if err != nil {
			return nil, fmt.Errorf("error converting ProcessedByVolunteerID UUID: %w", err)
		}
	}

	if procDate.Valid {
		form.ProcessingDate = procDate.Time
	}

	if rejReason.Valid {
		form.RejectionReason = rejReason.String
	}

	// Convert UUIDs
	form.AdopterID, err = SwapUUIDFormat(form.AdopterID)
	if err != nil {
		return nil, fmt.Errorf("error converting AdopterID UUID: %w", err)
	}

	form.DogID, err = SwapUUIDFormat(form.DogID)
	if err != nil {
		return nil, fmt.Errorf("error converting DogID UUID: %w", err)
	}

	return &form, nil
}

// Create inserts a new adoption form into the database
func (r *mssqlAdoptionFormRepository) Create(ctx context.Context, form *domain.AdoptionForm) error {
	query := `INSERT INTO shelter.AdoptionForm
              (AdopterID, DogID, SubmissionDate, ProcessedByVolunteerID, ProcessingDate, Status, RejectionReason)
              VALUES
              (@p1, @p2, @p3, @p4, @p5, @p6, @p7);
              SELECT SCOPE_IDENTITY()`

	// Set default values if not provided
	if form.SubmissionDate.IsZero() {
		form.SubmissionDate = time.Now()
	}

	if form.Status == "" {
		form.Status = "Pending"
	}

	// Handle NULL values
	var procVolID sql.NullString
	if form.ProcessedByVolunteerID != uuid.Nil {
		procVolID = sql.NullString{String: form.ProcessedByVolunteerID.String(), Valid: true}
	}

	var procDate sql.NullTime
	if !form.ProcessingDate.IsZero() {
		procDate = sql.NullTime{Time: form.ProcessingDate, Valid: true}
	}

	var rejReason sql.NullString
	if form.RejectionReason != "" {
		rejReason = sql.NullString{String: form.RejectionReason, Valid: true}
	}

	// Execute the query and get the generated ID
	err := r.conn.QueryRowContext(
		ctx,
		query,
		form.AdopterID,
		form.DogID,
		form.SubmissionDate,
		procVolID,
		procDate,
		form.Status,
		rejReason,
	).Scan(&form.AdoptionFormID)

	if err != nil {
		return fmt.Errorf("error creating adoption form: %w", err)
	}

	return nil
}

// Update modifies an existing adoption form in the database
func (r *mssqlAdoptionFormRepository) Update(ctx context.Context, form *domain.AdoptionForm) error {
	query := `UPDATE shelter.AdoptionForm
              SET AdopterID = @p1, 
                  DogID = @p2, 
                  SubmissionDate = @p3, 
                  ProcessedByVolunteerID = @p4, 
                  ProcessingDate = @p5, 
                  Status = @p6, 
                  RejectionReason = @p7
              WHERE AdoptionFormID = @p8`

	if form.AdoptionFormID <= 0 {
		return domain.ErrInvalidInput
	}

	// Handle NULL values
	var procVolID sql.NullString
	if form.ProcessedByVolunteerID != uuid.Nil {
		procVolID = sql.NullString{String: form.ProcessedByVolunteerID.String(), Valid: true}
	}

	var procDate sql.NullTime
	if !form.ProcessingDate.IsZero() {
		procDate = sql.NullTime{Time: form.ProcessingDate, Valid: true}
	}

	var rejReason sql.NullString
	if form.RejectionReason != "" {
		rejReason = sql.NullString{String: form.RejectionReason, Valid: true}
	}

	result, err := r.conn.ExecContext(
		ctx,
		query,
		form.AdopterID,
		form.DogID,
		form.SubmissionDate,
		procVolID,
		procDate,
		form.Status,
		rejReason,
		form.AdoptionFormID,
	)
	if err != nil {
		return fmt.Errorf("error updating adoption form: %w", err)
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

// Delete removes an adoption form from the database
func (r *mssqlAdoptionFormRepository) Delete(ctx context.Context, formID int) error {
	query := `DELETE FROM shelter.AdoptionForm
              WHERE AdoptionFormID = @p1`

	if formID <= 0 {
		return domain.ErrInvalidInput
	}

	result, err := r.conn.ExecContext(ctx, query, formID)
	if err != nil {
		return fmt.Errorf("error deleting adoption form: %w", err)
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

// ApproveForm changes an adoption form status to "Approved" using a stored procedure
func (r *mssqlAdoptionFormRepository) ApproveForm(ctx context.Context, formID int, volunteerID uuid.UUID) error {
	query := `EXEC ApproveAdoptionForm @AdoptionFormID = @p1, @ProcessedByVolunteerID = @p2`

	if formID <= 0 {
		return domain.ErrInvalidInput
	}

	if volunteerID == uuid.Nil {
		return domain.ErrInvalidInput
	}

	result, err := r.conn.ExecContext(ctx, query, formID, volunteerID)
	if err != nil {
		return fmt.Errorf("error approving adoption form: %w", err)
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

// RejectForm changes an adoption form status to "Rejected" with a reason using a stored procedure
func (r *mssqlAdoptionFormRepository) RejectForm(ctx context.Context, formID int, volunteerID uuid.UUID, reason string) error {
	query := `EXEC RejectAdoptionForm @AdoptionFormID = @p1, @ProcessedByVolunteerID = @p2, @RejectionReason = @p3`

	if formID <= 0 {
		return domain.ErrInvalidInput
	}

	if volunteerID == uuid.Nil {
		return domain.ErrInvalidInput
	}

	result, err := r.conn.ExecContext(ctx, query, formID, volunteerID, reason)
	if err != nil {
		return fmt.Errorf("error rejecting adoption form: %w", err)
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

// CompleteForm changes an adoption form status to "Completed" and marks the dog as adopted
func (r *mssqlAdoptionFormRepository) CompleteForm(ctx context.Context, formID int, volunteerID uuid.UUID) error {
	query := `EXEC CompleteAdoptionForm @AdoptionFormID = @p1, @ProcessedByVolunteerID = @p2`

	if formID <= 0 {
		return domain.ErrInvalidInput
	}

	if volunteerID == uuid.Nil {
		return domain.ErrInvalidInput
	}

	result, err := r.conn.ExecContext(ctx, query, formID, volunteerID)
	if err != nil {
		return fmt.Errorf("error completing adoption form: %w", err)
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
