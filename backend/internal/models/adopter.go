package models

import "github.com/google/uuid"

type Adopter struct {
	AdopterID          uuid.UUID `json:"adopterId" db:"AdopterID"`
	HasPetAllergies    bool      `json:"hasPetAllergies" db:"HasPetAllergies"`
	HasSurrenderedPets bool      `json:"hasSurrenderedPets" db:"HasSurrenderedPets"`
	HomeStatus         string    `json:"homeStatus" db:"HomeStatus"`
}
