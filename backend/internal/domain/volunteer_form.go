package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// VolunteerForm represents an application to volunteer
type VolunteerForm struct {
	VolunteerFormID                int       `json:"volunteerFormId"`
	ApplicantID                    uuid.UUID `json:"applicantId"`
	SubmissionDate                 time.Time `json:"submissionDate"`
	SupportsAnimalWelfareEducation bool      `json:"supportsAnimalWelfareEducation"`
	AvailableShifts                string    `json:"availableShifts"`
	SupportsResponsibleBreeding    bool      `json:"supportsResponsibleBreeding"`
	AcceptsCleaningDuties          bool      `json:"acceptsCleaningDuties"`
	AcceptsDogCare                 bool      `json:"acceptsDogCare"`
	HasDogAllergies                bool      `json:"hasDogAllergies"`
	HasPhysicalLimitations         bool      `json:"hasPhysicalLimitations"`
	IsForCommunityService          bool      `json:"isForCommunityService"`
	RequiredServiceHours           int       `json:"requiredServiceHours"`
	ReferralSource                 string    `json:"referralSource"`
	CommentsAndQuestions           string    `json:"commentsAndQuestions"`
	ProcessedByVolunteerID         uuid.UUID `json:"processedByVolunteerId"`
	ProcessingDate                 time.Time `json:"processingDate"`
	Status                         string    `json:"status"`
	RejectionReason                string    `json:"rejectionReason"`
}

// VolunteerFormRepository defines the interface for volunteer form data access operations
type VolunteerFormRepository interface {
	GetAll(ctx context.Context) ([]VolunteerForm, error)
	GetByID(ctx context.Context, formID int) (*VolunteerForm, error)
	GetByApplicantID(ctx context.Context, applicantID uuid.UUID) ([]VolunteerForm, error)
	Create(ctx context.Context, form *VolunteerForm) error
	Update(ctx context.Context, form *VolunteerForm) error
	Delete(ctx context.Context, formID int) error

	// Domain-specific operations
	GetByStatus(ctx context.Context, status string) ([]VolunteerForm, error)
	Process(ctx context.Context, formID int, volunteerID uuid.UUID, status string, rejectionReason string) error
	GetPendingForms(ctx context.Context) ([]VolunteerForm, error)
}
