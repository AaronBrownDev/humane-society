package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type SQLPersonRepository struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) domain.PersonRepository {
	return &SQLPersonRepository{db}
}

func (r *SQLPersonRepository) GetAll() ([]domain.Person, error) {
	query := `
			SELECT PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, EmailAddress, PhoneNumber FROM people.Person
	`
	rows, err := r.db.Query(query)
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
func (r *SQLPersonRepository) GetByID(personID uuid.UUID) (*domain.Person, error) {
	query := `
			SELECT PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, EmailAddress, PhoneNumber FROM people.Person 
			WHERE PersonID = @p1
	`
	row := r.db.QueryRow(query, personID)

	var person domain.Person

	err := row.Scan(
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("person not found")
		}
		return nil, err
	}

	return &person, nil
}

func (r *SQLPersonRepository) Create(person *domain.Person) error {
	query := `
			INSERT INTO people.Person
			(PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, EmailAddress, PhoneNumber)
			VALUES
			(@p1, @p2, @p3, @p4, @p5, @p6, @p7, @p8)
	`

	if person.PersonID == uuid.Nil {
		person.PersonID = uuid.New()
	}

	_, err := r.db.Exec(
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

	return nil
}

func (r *SQLPersonRepository) Update(person *domain.Person) error {
	query := `
			UPDATE people.Person
			SET FirstName = @p1, LastName = @p2, BirthDate = @p3, PhysicalAddress = @p4, MailingAddress = @p5, EmailAddress = @p6, PhoneNumber = @p7
			WHERE PersonID = @p8
	`

	result, err := r.db.Exec(
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

	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		if err != nil {
			return err
		}
		return errors.New("person not found")
	}

	return nil
}

func (r *SQLPersonRepository) Delete(personID uuid.UUID) error {
	query := `
			DELETE FROM people.Person WHERE PersonID = @p1
	`

	result, err := r.db.Exec(query, personID)
	if err != nil {
		return fmt.Errorf("error deleting person: %w", err)
	}

	if rowsAffected, err := result.RowsAffected(); rowsAffected == 0 {
		if err != nil {
			return err
		}
		return errors.New("person not found")
	}

	return nil
}
