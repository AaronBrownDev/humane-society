package models

import (
	"github.com/google/uuid"
	"time"
)

type Supply struct {
	SupplyID        int       `json:"supplyId" db:"SupplyID"`
	ItemID          uuid.UUID `json:"itemId" db:"ItemID"`
	Quantity        int       `json:"quantity" db:"Quantity"`
	StorageLocation string    `json:"storageLocation" db:"StorageLocation"`
	ExpirationDate  time.Time `json:"expirationDate" db:"ExpirationDate"`
	BatchNumber     string    `json:"batchNumber" db:"BatchNumber"`
	AcquisitionDate time.Time `json:"acquisitionDate" db:"AcquisitionDate"`
}
