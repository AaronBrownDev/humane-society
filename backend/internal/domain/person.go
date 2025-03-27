package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
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
	Create(ctx context.Context, person *Person) error
	Update(ctx context.Context, person *Person) error
	Delete(ctx context.Context, personID uuid.UUID) error
}
