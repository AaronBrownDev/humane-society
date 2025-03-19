package models

import "github.com/google/uuid"

type Veterinarian struct {
	VeterinarianID uuid.UUID `json:"veterinarianId" db:"VeterinarianID"`
}
