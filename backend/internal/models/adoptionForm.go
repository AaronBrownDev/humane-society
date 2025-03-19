package models

import (
	"github.com/google/uuid"
	"time"
)

type AdoptionForm struct {
	AdoptionFormID         int       `json:"adoptionFormId" db:"AdoptionFormID"`
	AdopterID              uuid.UUID `json:"adopterId" db:"AdopterID"`
	DogID                  uuid.UUID `json:"dogId" db:"DogID"`
	SubmissionDate         time.Time `json:"submissionDate" db:"SubmissionDate"`
	ProcessedByVolunteerID uuid.UUID `json:"processedByVolunteerId" db:"ProcessedByVolunteerID"`
	ProcessingDate         time.Time `json:"processingDate" db:"ProcessingDate"`
	Status                 string    `json:"status" db:"Status"`
	RejectionReason        string    `json:"rejectionReason" db:"RejectionReason"`
}
