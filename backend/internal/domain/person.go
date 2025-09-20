package domain

import (
	"context"
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
)

// Person represents a base person entity
type Person struct {
	PersonID        uuid.UUID `json:"personId"`
	FirstName       string    `json:"firstName"`
	LastName        string    `json:"lastName"`
	BirthDate       time.Time `json:"birthDate"`
	PhysicalAddress string    `json:"physicalAddress"`
	MailingAddress  string    `json:"mailingAddress"`
	EmailAddress    string    `json:"emailAddress"`
	PhoneNumber     string    `json:"phoneNumber"`
}

// PersonRepository defines the interface for person data access operations
type PersonRepository interface {
	GetAll(ctx context.Context) ([]Person, error)
	GetByID(ctx context.Context, personID uuid.UUID) (*Person, error)
	GetByEmail(ctx context.Context, emailAddress string) (*Person, error)
	Create(ctx context.Context, person *Person) error
	Update(ctx context.Context, person *Person) error
	Delete(ctx context.Context, personID uuid.UUID) error
}

var (
	ErrFirstNameRequired = errors.New("first name is required")
	ErrLastNameRequired  = errors.New("last name is required")
	ErrInvalidEmail      = errors.New("invalid email address")
	ErrInvalidPhone      = errors.New("invalid phone number")
	ErrInvalidBirthDate  = errors.New("birth date cannot be in the future")
	ErrTooYoung          = errors.New("person must be at least 18 years old")
)

// Validate
func (p *Person) Validate() error {
	// Check required fields
	if strings.TrimSpace(p.FirstName) == "" {
		return ErrFirstNameRequired
	}

	if strings.TrimSpace(p.LastName) == "" {
		return ErrLastNameRequired
	}

	if !isValidEmail(p.EmailAddress) {
		return ErrInvalidEmail
	}

	if !isValidPhone(p.PhoneNumber) {
		return ErrInvalidPhone
	}

	if p.BirthDate.After(time.Now()) {
		return ErrInvalidBirthDate
	}

	if calculateAge(p.BirthDate) < 18 {
		return ErrTooYoung
	}

	return nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return emailRegex.MatchString(strings.ToLower(email))
}

func isValidPhone(phone string) bool {
	// Remove all non-digits
	digitsOnly := regexp.MustCompile(`\D`).ReplaceAllString(phone, "")
	// TODO: might allow other types so might need to change later
	// US phone should have 10 digits
	return len(digitsOnly) == 10
}

func calculateAge(birthDate time.Time) int {
	now := time.Now()
	years := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay() {
		years--
	}
	return years
}
