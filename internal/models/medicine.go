package models

type Medicine struct {
	MedicineID   int    `json:"medicineId" db:"MedicineID"`
	Name         string `json:"name" db:"Name"`
	Manufacturer string `json:"manufacturer" db:"Manufacturer"`
	Description  string `json:"description" db:"Description"`
	DosageUnit   string `json:"dosageUnit" db:"DosageUnit"`
}
