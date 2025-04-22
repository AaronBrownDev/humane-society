package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

// mssqlPersonRepository implements domain.PersonRepository interface using SQL Server.
// It handles database operations for person entities including core CRUD operations
// and email-based lookups.
type mssqlPersonRepository struct {
	conn Connection
}

// NewMSSQLPerson creates a new repository instance for person operations
// using the provided database connection.
func NewMSSQLPerson(conn Connection) domain.PersonRepository {
	return &mssqlPersonRepository{conn: conn}
}

// fetch is a helper function to retrieve multiple person records based on the provided query.
// It handles row scanning, result collection, and proper resource cleanup.
// The function returns the matched people or an error if the database operation fails.
func (r *mssqlPersonRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]domain.Person, error) {
	rows, err := r.conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var people []domain.Person
	for rows.Next() {
		var person domain.Person
		var birthDate sql.NullTime

		err = rows.Scan(
			&person.PersonID,
			&person.FirstName,
			&person.LastName,
			&birthDate,
			&person.PhysicalAddress,
			&person.MailingAddress,
			&person.EmailAddress,
			&person.PhoneNumber,
		)
		if err != nil {
			return nil, err
		}

		if birthDate.Valid {
			person.BirthDate = birthDate.Time
		}

		// Convert UUID from SQL Server format to standard format
		person.PersonID, err = SwapUUIDFormat(person.PersonID)
		if err != nil {
			return nil, fmt.Errorf("error converting PersonID UUID: %w", err)
		}

		people = append(people, person)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return people, nil
}

// GetAll returns all person records from the database.
// Returns a slice of all persons or an error if the database operation fails.
func (r *mssqlPersonRepository) GetAll(ctx context.Context) ([]domain.Person, error) {
	query := `SELECT PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, EmailAddress, PhoneNumber 
              FROM people.Person`

	return r.fetch(ctx, query)
}

// GetByID returns a person with the specified ID.
// Returns a pointer to the person if found, domain.ErrNotFound if no matching person exists,
// or another error if the database operation fails.
func (r *mssqlPersonRepository) GetByID(ctx context.Context, personID uuid.UUID) (*domain.Person, error) {
	query := `SELECT PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, EmailAddress, PhoneNumber 
              FROM people.Person 
              WHERE PersonID = @p1`

	var person domain.Person
	var birthDate sql.NullTime

	err := r.conn.QueryRowContext(ctx, query, personID).Scan(
		&person.PersonID,
		&person.FirstName,
		&person.LastName,
		&birthDate,
		&person.PhysicalAddress,
		&person.MailingAddress,
		&person.EmailAddress,
		&person.PhoneNumber,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	if birthDate.Valid {
		person.BirthDate = birthDate.Time
	}

	// Convert UUID from SQL Server format to standard format
	person.PersonID, err = SwapUUIDFormat(person.PersonID)
	if err != nil {
		return nil, fmt.Errorf("error converting PersonID UUID: %w", err)
	}

	return &person, nil
}

// GetByEmail retrieves a specific person by their email address.
// Returns domain.ErrInvalidInput if the email is empty,
// domain.ErrNotFound if no person with the given email exists,
// or another error if the database operation fails.
func (r *mssqlPersonRepository) GetByEmail(ctx context.Context, email string) (*domain.Person, error) {
	query := `SELECT PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, EmailAddress, PhoneNumber 
              FROM people.Person 
              WHERE EmailAddress = @p1`

	if email == "" {
		return nil, domain.ErrInvalidInput
	}

	var person domain.Person
	var birthDate sql.NullTime // Handle nullable birth date

	err := r.conn.QueryRowContext(ctx, query, email).Scan(
		&person.PersonID,
		&person.FirstName,
		&person.LastName,
		&birthDate,
		&person.PhysicalAddress,
		&person.MailingAddress,
		&person.EmailAddress,
		&person.PhoneNumber,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}

	// Handle nullable birth date
	if birthDate.Valid {
		person.BirthDate = birthDate.Time
	}

	// Convert UUID from SQL Server format to standard format
	person.PersonID, err = SwapUUIDFormat(person.PersonID)
	if err != nil {
		return nil, fmt.Errorf("error converting PersonID UUID: %w", err)
	}

	return &person, nil
}

// Create inserts a new person record into the database.
// Generates a new UUID for the person if one is not provided.
// Returns domain.ErrDatabaseError if the insertion fails to affect any rows,
// or another error if the database operation fails.
func (r *mssqlPersonRepository) Create(ctx context.Context, person *domain.Person) error {
	query := `INSERT INTO people.Person
              (PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, EmailAddress, PhoneNumber)
              VALUES
              (@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8)`

	// Generate a new UUID if none is provided
	if person.PersonID == uuid.Nil {
		person.PersonID = uuid.New()
	}

	// Convert to SQL Server UUID format
	sqlUUID, err := SwapUUIDFormat(person.PersonID)
	if err != nil {
		return fmt.Errorf("error converting UUID format: %w", err)
	}

	// Handle NULL birth date
	var birthDate interface{}
	if person.BirthDate.IsZero() {
		birthDate = nil
	} else {
		birthDate = person.BirthDate
	}

	result, err := r.conn.ExecContext(
		ctx,
		query,
		sqlUUID,
		person.FirstName,
		person.LastName,
		birthDate,
		person.PhysicalAddress,
		person.MailingAddress,
		person.EmailAddress,
		person.PhoneNumber,
	)
	if err != nil {
		return fmt.Errorf("error creating person: %w", err)
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

// Update modifies an existing person record in the database.
// Returns domain.ErrInvalidInput if the person ID is nil,
// domain.ErrNotFound if no person with the given ID exists,
// or another error if the database operation fails.
func (r *mssqlPersonRepository) Update(ctx context.Context, person *domain.Person) error {
	query := `UPDATE people.Person
              SET FirstName = @p1, 
                  LastName = @p2, 
                  BirthDate = @p3, 
                  PhysicalAddress = @p4, 
                  MailingAddress = @p5, 
                  EmailAddress = @p6, 
                  PhoneNumber = @p7
              WHERE PersonID = @p8`

	if person.PersonID == uuid.Nil {
		return domain.ErrInvalidInput
	}

	// Handle NULL birth date
	var birthDate interface{}
	if person.BirthDate.IsZero() {
		birthDate = nil
	} else {
		birthDate = person.BirthDate
	}

	result, err := r.conn.ExecContext(
		ctx,
		query,
		person.FirstName,
		person.LastName,
		birthDate,
		person.PhysicalAddress,
		person.MailingAddress,
		person.EmailAddress,
		person.PhoneNumber,
		person.PersonID,
	)
	if err != nil {
		return fmt.Errorf("error updating person: %w", err)
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

// Delete removes a person record from the database by its UUID.
// Returns domain.ErrInvalidInput if the person ID is nil,
// domain.ErrNotFound if no person with the given ID exists,
// or another error if the database operation fails.
func (r *mssqlPersonRepository) Delete(ctx context.Context, personID uuid.UUID) error {
	query := `DELETE FROM people.Person WHERE PersonID = @p1`

	if personID == uuid.Nil {
		return domain.ErrInvalidInput
	}

	result, err := r.conn.ExecContext(ctx, query, personID)
	if err != nil {
		return fmt.Errorf("error deleting person: %w", err)
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
