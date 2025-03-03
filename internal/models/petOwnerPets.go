package models

import (
	"github.com/google/uuid"
	"time"
)

type PetOwnerPets struct {
	PetID             int       `json:"petId" db:"PetID"`
	PetOwnerID        uuid.UUID `json:"petOwnerId" db:"PetOwnerID"`
	Name              string    `json:"name" db:"Name"`
	Type              string    `json:"type" db:"Type"`
	Breed             string    `json:"breed" db:"Breed"`
	Sex               string    `json:"sex" db:"Sex"`
	OwnershipDate     time.Time `json:"ownershipDate" db:"OwnershipDate"`
	LivingEnvironment string    `json:"livingEnvironment" db:"LivingEnvironment"`
}
