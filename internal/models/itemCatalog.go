package models

import "github.com/google/uuid"

type ItemCatalog struct {
	ItemID          uuid.UUID `json:"itemId" db:"ItemID"`
	Name            string    `json:"name" db:"Name"`
	Category        string    `json:"category" db:"Category"`
	Description     string    `json:"description" db:"Description"`
	MinimumQuantity int       `json:"minimumQuantity" db:"MinimumQuantity"`
	IsActive        bool      `json:"isActive" db:"IsActive"`
}
