package domain

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestAdopter_Validate(t *testing.T) {
	testCaseUUID := uuid.New()
	testCasePerson := Person{
		PersonID:        testCaseUUID,
		FirstName:       "John",
		LastName:        "Doe",
		BirthDate:       time.Date(1980, time.January, 01, 0, 0, 0, 0, time.UTC),
		PhysicalAddress: "123 Main Street, Anytown, USA",
		MailingAddress:  "123 Main Street, Anytown, USA",
		EmailAddress:    "john@example.com",
		PhoneNumber:     "012-345-6789",
	}

	tests := []struct {
		name    string
		adopter Adopter
		wantErr bool
		errMsg  string
	}{
		{
			name: "valid adopter - pending home status",
			adopter: Adopter{
				AdopterID:          testCaseUUID,
				Person:             testCasePerson,
				HasPetAllergies:    false,
				HasSurrenderedPets: false,
				HomeStatus:         "Pending",
			},
			wantErr: false,
		},
		{
			name: "valid adopter - approved home status",
			adopter: Adopter{
				AdopterID:          testCaseUUID,
				Person:             testCasePerson,
				HasPetAllergies:    false,
				HasSurrenderedPets: false,
				HomeStatus:         "Approved",
			},
			wantErr: false,
		},
		{
			name: "valid adopter - rejected home status",
			adopter: Adopter{
				AdopterID:          testCaseUUID,
				Person:             testCasePerson,
				HasPetAllergies:    false,
				HasSurrenderedPets: false,
				HomeStatus:         "Rejected",
			},
			wantErr: false,
		},
		{
			name: "valid adopter - empty home status defaults to pending", // Should default to Pending
			adopter: Adopter{
				AdopterID:          testCaseUUID,
				Person:             testCasePerson,
				HasPetAllergies:    false,
				HasSurrenderedPets: false,
				HomeStatus:         "",
			},
			wantErr: false,
		},
		{
			name: "invalid adopter - invalid/empty person",
			adopter: Adopter{
				AdopterID:          testCaseUUID,
				Person:             Person{},
				HasPetAllergies:    false,
				HasSurrenderedPets: false,
				HomeStatus:         "Pending",
			},
			wantErr: true,
			errMsg:  ErrFirstNameRequired.Error(),
		},
		{
			name: "invalid adopter - invalid home status",
			adopter: Adopter{
				AdopterID:          testCaseUUID,
				Person:             testCasePerson,
				HasPetAllergies:    false,
				HasSurrenderedPets: false,
				HomeStatus:         "example",
			},
			wantErr: true,
			errMsg:  ErrInvalidHomeStatus.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.adopter.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, WantErr %v", err, tt.wantErr)
			}

			if err != nil && tt.errMsg != "" && !strings.Contains(err.Error(), tt.errMsg) {
				t.Errorf("Validate() error = %v, want error containing %v", err, tt.errMsg)
			}
		})
	}
}
