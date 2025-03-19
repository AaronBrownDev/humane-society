package models

import "github.com/google/uuid"

type PetOwner struct {
	PetOwnerID                 uuid.UUID `json:"petOwnerId" db:"PetOwnerID"`
	VeterinarianID             uuid.UUID `json:"veterinarianId" db:"VeterinarianID"`
	HasSterilizedPets          bool      `json:"hasSterilizedPets" db:"HasSterilizedPets"`
	HasVaccinatedPets          bool      `json:"hasVaccinatedPets" db:"HasVaccinatedPets"`
	UsesVetHeartWormPrevention bool      `json:"usesVetHeartWormPrevention" db:"UsesVetHeartWormPrevention"`
}
