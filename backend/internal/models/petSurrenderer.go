package models

import "github.com/google/uuid"

type PetSurrenderer struct {
	SurrendererID uuid.UUID `json:"surrendererId" db:"SurrendererID"`
}
