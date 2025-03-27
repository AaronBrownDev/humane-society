package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// SurrenderForm represents a form for surrendering a pet
type SurrenderForm struct {
	SurrenderFormID        int       `json:"surrenderFormId"`
	SurrendererID          uuid.UUID `json:"surrendererId"`
	SubmissionDate         time.Time `json:"submissionDate"`
	DogName                string    `json:"dogName"`
	DogAge                 int       `json:"dogAge"`
	WeightInPounds         float64   `json:"weightInPounds"`
	Sex                    string    `json:"sex"`
	Breed                  string    `json:"breed"`
	Color                  string    `json:"color"`
	LivingEnvironment      string    `json:"livingEnvironment"`
	OwnershipDate          time.Time `json:"ownershipDate"`
	VeterinarianID         uuid.UUID `json:"veterinarianId"`
	LastVetVisitDate       time.Time `json:"lastVetVisitDate"`
	IsGoodWithChildren     bool      `json:"isGoodWithChildren"`
	IsGoodWithDogs         bool      `json:"isGoodWithDogs"`
	IsGoodWithCats         bool      `json:"isGoodWithCats"`
	IsGoodWithStrangers    bool      `json:"isGoodWithStrangers"`
	IsHouseTrained         bool      `json:"isHouseTrained"`
	IsSterilized           bool      `json:"isSterilized"`
	MicroChipNumber        string    `json:"microChipNumber"`
	MedicalProblems        string    `json:"medicalProblems"`
	BiteHistory            string    `json:"biteHistory"`
	SurrenderReason        string    `json:"surrenderReason"`
	ProcessedByVolunteerID uuid.UUID `json:"processedByVolunteerId"`
	ProcessingDate         time.Time `json:"processingDate"`
	ResultingDogID         uuid.UUID `json:"resultingDogId"`
	Status                 string    `json:"status"`
}

// SurrenderFormRepository defines the interface for surrender form data access operations
type SurrenderFormRepository interface {
	GetAll(ctx context.Context) ([]SurrenderForm, error)
	GetByID(ctx context.Context, surrenderFormID int) (*SurrenderForm, error)
	GetBySurrendererID(ctx context.Context, surrendererID uuid.UUID) ([]SurrenderForm, error)
	Create(ctx context.Context, form *SurrenderForm) error
	Update(ctx context.Context, form *SurrenderForm) error
	Delete(ctx context.Context, surrenderFormID int) error

	// Domain-specific operations
	GetByStatus(ctx context.Context, status string) ([]SurrenderForm, error)
	Process(ctx context.Context, surrenderFormID int, volunteerID uuid.UUID, status string) error
	GetPendingForms(ctx context.Context) ([]SurrenderForm, error)
}
