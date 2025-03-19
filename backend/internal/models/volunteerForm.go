package models

import (
	"github.com/google/uuid"
	"time"
)

type VolunteerForm struct {
	VolunteerFormID                int       `json:"volunteerFormId" db:"VolunteerFormID"`
	ApplicantID                    uuid.UUID `json:"applicantId" db:"ApplicantID"`
	SubmissionDate                 time.Time `json:"submissionDate" db:"SubmissionDate"`
	SupportsAnimalWelfareEducation bool      `json:"supportsAnimalWelfareEducation" db:"SupportsAnimalWelfareEducation"`
	AvailableShifts                string    `json:"availableShifts" db:"AvailableShifts"`
	SupportsResponsibleBreeding    bool      `json:"supportsResponsibleBreeding" db:"SupportsResponsibleBreeding"`
	AcceptsCleaningDuties          bool      `json:"acceptsCleaningDuties" db:"AcceptsCleaningDuties"`
	AcceptsDogCare                 bool      `json:"acceptsDogCare" db:"AcceptsDogCare"`
	HasDogAllergies                bool      `json:"hasDogAllergies" db:"HasDogAllergies"`
	HasPhysicalLimitations         bool      `json:"hasPhysicalLimitations" db:"HasPhysicalLimitations"`
	IsForCommunityService          bool      `json:"isForCommunityService" db:"IsForCommunityService"`
	RequiredServiceHours           int       `json:"requiredServiceHours" db:"RequiredServiceHours"`
	ReferralSource                 string    `json:"referralSource" db:"ReferralSource"`
	CommentsAndQuestions           string    `json:"commentsAndQuestions" db:"CommentsAndQuestions"`
	ProcessedByVolunteerID         uuid.UUID `json:"processedByVolunteerId" db:"ProcessedByVolunteerID"`
	ProcessingDate                 time.Time `json:"processingDate" db:"ProcessingDate"`
	Status                         string    `json:"status" db:"Status"`
	RejectionReason                string    `json:"rejectionReason" db:"RejectionReason"`
}
