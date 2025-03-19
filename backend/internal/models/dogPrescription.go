package models

import (
	"github.com/google/uuid"
	"time"
)

type DogPrescription struct {
	PrescriptionID  int       `json:"prescriptionId" db:"PrescriptionID"`
	DogID           uuid.UUID `json:"dogId" db:"DogID"`
	MedicineID      int       `json:"medicineId" db:"MedicineID"`
	Dosage          float64   `json:"dosage" db:"Dosage"`
	Frequency       string    `json:"frequency" db:"Frequency"`
	StartDate       time.Time `json:"startDate" db:"StartDate"`
	EndDate         time.Time `json:"endDate" db:"EndDate"`
	Notes           string    `json:"notes" db:"Notes"`
	VetPrescriberID uuid.UUID `json:"vetPrescriberId" db:"VetPrescriberID"`
}
