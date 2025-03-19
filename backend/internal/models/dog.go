package models

import (
	"github.com/google/uuid"
	"time"
)

type Dog struct {
	DogID              uuid.UUID `json:"dogId" db:"DogID"`
	Name               string    `json:"name" db:"Name"`
	IntakeDate         time.Time `json:"intakeDate" db:"IntakeDate"`
	EstimatedBirthDate time.Time `json:"estimatedBirthDate" db:"EstimatedBirthDate"`
	Breed              string    `json:"breed" db:"Breed"`
	Sex                string    `json:"sex" db:"Sex"`
	Color              string    `json:"color" db:"Color"`
	CageNumber         int       `json:"cageNumber" db:"CageNumber"`
	IsAdopted          bool      `json:"isAdopted" db:"IsAdopted"`
}
