package models

import (
	"github.com/google/uuid"
	"time"
)

type SurrenderForm struct {
	SurrenderFormID        int       `json:"surrenderFormId" db:"SurrenderFormID"`
	SurrendererID          uuid.UUID `json:"surrendererId" db:"SurrendererID"`
	SubmissionDate         time.Time `json:"submissionDate" db:"SubmissionDate"`
	DogName                string    `json:"dogName" db:"DogName"`
	DogAge                 int       `json:"dogAge" db:"DogAge"`
	WeightInPounds         float64   `json:"weightInPounds" db:"WeightInPounds"`
	Sex                    string    `json:"sex" db:"Sex"`
	Breed                  string    `json:"breed" db:"Breed"`
	Color                  string    `json:"color" db:"Color"`
	LivingEnvironment      string    `json:"livingEnvironment" db:"LivingEnvironment"`
	OwnershipDate          time.Time `json:"ownershipDate" db:"OwnershipDate"`
	VeterinarianID         uuid.UUID `json:"veterinarianId" db:"VeterinarianID"`
	LastVetVisitDate       time.Time `json:"lastVetVisitDate" db:"LastVetVisitDate"`
	IsGoodWithChildren     bool      `json:"isGoodWithChildren" db:"IsGoodWithChildren"`
	IsGoodWithDogs         bool      `json:"isGoodWithDogs" db:"IsGoodWithDogs"`
	IsGoodWithCats         bool      `json:"isGoodWithCats" db:"IsGoodWithCats"`
	IsGoodWithStrangers    bool      `json:"isGoodWithStrangers" db:"IsGoodWithStrangers"`
	IsHouseTrained         bool      `json:"isHouseTrained" db:"IsHouseTrained"`
	IsSterilized           bool      `json:"isSterilized" db:"IsSterilized"`
	MicroChipNumber        string    `json:"microChipNumber" db:"MicroChipNumber"`
	MedicalProblems        string    `json:"medicalProblems" db:"MedicalProblems"`
	BiteHistory            string    `json:"biteHistory" db:"BiteHistory"`
	SurrenderReason        string    `json:"surrenderReason" db:"SurrenderReason"`
	ProcessedByVolunteerID uuid.UUID `json:"processedByVolunteerId" db:"ProcessedByVolunteerID"`
	ProcessingDate         time.Time `json:"processingDate" db:"ProcessingDate"`
	ResultingDogID         uuid.UUID `json:"resultingDogId" db:"ResultingDogID"`
	Status                 string    `json:"status" db:"Status"`
}
