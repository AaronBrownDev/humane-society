package domain

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestPerson_Validate(t *testing.T) {
	tests := []struct {
		name    string
		person  Person
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid person",
			person: Person{
				PersonID:        uuid.New(),
				FirstName:       "John",
				LastName:        "Doe",
				BirthDate:       time.Date(1980, time.January, 01, 0, 0, 0, 0, time.UTC),
				PhysicalAddress: "123 Main Street, Anytown, USA",
				MailingAddress:  "123 Main Street, Anytown, USA",
				EmailAddress:    "john@example.com",
				PhoneNumber:     "012-345-6789",
			},
			wantErr: false,
		},
		{
			name: "invalid person - missing first name",
			person: Person{
				PersonID:        uuid.New(),
				LastName:        "Doe",
				BirthDate:       time.Date(1980, time.January, 01, 0, 0, 0, 0, time.UTC),
				PhysicalAddress: "123 Main Street, Anytown, USA",
				MailingAddress:  "123 Main Street, Anytown, USA",
				EmailAddress:    "john@example.com",
				PhoneNumber:     "012-345-6789",
			},
			wantErr: true,
			errMsg:  ErrFirstNameRequired.Error(),
		},
		{
			name: "invalid person - missing last name",
			person: Person{
				PersonID:        uuid.New(),
				FirstName:       "John",
				BirthDate:       time.Date(1980, time.January, 01, 0, 0, 0, 0, time.UTC),
				PhysicalAddress: "123 Main Street, Anytown, USA",
				MailingAddress:  "123 Main Street, Anytown, USA",
				EmailAddress:    "john@example.com",
				PhoneNumber:     "012-345-6789",
			},
			wantErr: true,
			errMsg:  ErrLastNameRequired.Error(),
		},
		{
			name: "invalid person - invalid email address",
			person: Person{
				PersonID:        uuid.New(),
				FirstName:       "John",
				LastName:        "Doe",
				BirthDate:       time.Date(1980, time.January, 01, 0, 0, 0, 0, time.UTC),
				PhysicalAddress: "123 Main Street, Anytown, USA",
				MailingAddress:  "123 Main Street, Anytown, USA",
				EmailAddress:    "johnexample.com",
				PhoneNumber:     "012-345-6789",
			},
			wantErr: true,
			errMsg:  ErrInvalidEmail.Error(),
		},
		{
			name: "invalid person - invalid phone number",
			person: Person{
				PersonID:        uuid.New(),
				FirstName:       "John",
				LastName:        "Doe",
				BirthDate:       time.Date(1980, time.January, 01, 0, 0, 0, 0, time.UTC),
				PhysicalAddress: "123 Main Street, Anytown, USA",
				MailingAddress:  "123 Main Street, Anytown, USA",
				EmailAddress:    "john@example.com",
				PhoneNumber:     "012-345-678",
			},
			wantErr: true,
			errMsg:  ErrInvalidPhone.Error(),
		},
		{
			name: "invalid person - future age",
			person: Person{
				PersonID:        uuid.New(),
				FirstName:       "John",
				LastName:        "Doe",
				BirthDate:       time.Now().AddDate(1, 0, 0),
				PhysicalAddress: "123 Main Street, Anytown, USA",
				MailingAddress:  "123 Main Street, Anytown, USA",
				EmailAddress:    "john@example.com",
				PhoneNumber:     "012-345-6789",
			},
			wantErr: true,
			errMsg:  ErrInvalidBirthDate.Error(),
		},
		{
			name: "invalid person - not old enough",
			person: Person{
				PersonID:        uuid.New(),
				FirstName:       "John",
				LastName:        "Doe",
				BirthDate:       time.Now().AddDate(-17, 0, 0),
				PhysicalAddress: "123 Main Street, Anytown, USA",
				MailingAddress:  "123 Main Street, Anytown, USA",
				EmailAddress:    "john@example.com",
				PhoneNumber:     "012-345-6789",
			},
			wantErr: true,
			errMsg:  ErrTooYoung.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.person.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, WantErr %v", err, tt.wantErr)
			}

			if err != nil && tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("Validate() error = %v, want error containing %v", err, tt.errMsg)
			}
		})
	}
}
