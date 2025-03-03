package models

import (
	"github.com/google/uuid"
	"time"
)

type Person struct {
	PersonID        uuid.UUID `json:"personId" db:"PersonID"`
	FirstName       string    `json:"firstName" db:"FirstName"`
	LastName        string    `json:"lastName" db:"LastName"`
	BirthDate       time.Time `json:"birthDate" db:"BirthDate"`
	PhysicalAddress string    `json:"physicalAddress" db:"PhysicalAddress"`
	MailingAddress  string    `json:"mailingAddress" db:"MailingAddress"`
	EmailAddress    string    `json:"emailAddress" db:"EmailAddress"`
	PhoneNumber     string    `json:"phoneNumber" db:"PhoneNumber"`
}
