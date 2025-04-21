package repository

import (
	"context"
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
		err = rows.Scan(
			&person.PersonID,
			&person.FirstName,
			&person.LastName,
			&person.BirthDate,
			&person.PhysicalAddress,
			&person.MailingAddress,
			&person.EmailAddress,
			&person.PhoneNumber,
		)
		if err != nil {
			return nil, err
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
	err := r.conn.QueryRowContext(ctx, query, personID).Scan(
		&person.PersonID,
		&person.FirstName,
		&person.LastName,
		&person.BirthDate,
		&person.PhysicalAddress,
		&person.MailingAddress,
		&person.EmailAddress,
		&person.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, errors.New("sql: no rows in result set")) {
			return nil, domain.ErrNotFound
		}
		return nil, err
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
	err := r.conn.QueryRowContext(ctx, query, email).Scan(
		&person.PersonID,
		&person.FirstName,
		&person.LastName,
		&person.BirthDate,
		&person.PhysicalAddress,
		&person.MailingAddress,
		&person.EmailAddress,
		&person.PhoneNumber,
	)
	if err != nil {
		if errors.Is(err, errors.New("sql: no rows in result set")) {
			return nil, domain.ErrNotFound
		}
		return nil, err
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

	result, err := r.conn.ExecContext(
		ctx,
		query,
		person.PersonID,
		person.FirstName,
		person.LastName,
		person.BirthDate,
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
              SET FirstName = @p1, LastName = @p2, BirthDate = @p3, PhysicalAddress = @p4,
                  MailingAddress = @p5, EmailAddress = @p6, PhoneNumber = @p7
              WHERE PersonID = @p8`

	if person.PersonID == uuid.Nil {
		return domain.ErrInvalidInput
	}

	result, err := r.conn.ExecContext(
		ctx,
		query,
		person.FirstName,
		person.LastName,
		person.BirthDate,
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
