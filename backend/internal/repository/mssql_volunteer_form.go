package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlVolunteerFormRepository struct {
	conn Connection
}

func NewMSSQLVolunteerForm(conn Connection) domain.VolunteerFormRepository {
	return &mssqlVolunteerFormRepository{conn: conn}
}

// TODO: Implement functions
func (r *mssqlVolunteerFormRepository) GetAll(ctx context.Context) ([]domain.VolunteerForm, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerFormRepository) GetByID(ctx context.Context, formID int) (*domain.VolunteerForm, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerFormRepository) GetByApplicantID(ctx context.Context, applicantID uuid.UUID) ([]domain.VolunteerForm, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerFormRepository) Create(ctx context.Context, form *domain.VolunteerForm) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerFormRepository) Update(ctx context.Context, form *domain.VolunteerForm) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerFormRepository) Delete(ctx context.Context, formID int) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerFormRepository) GetByStatus(ctx context.Context, status string) ([]domain.VolunteerForm, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerFormRepository) Process(ctx context.Context, formID int, volunteerID uuid.UUID, status string, rejectionReason string) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerFormRepository) GetPendingForms(ctx context.Context) ([]domain.VolunteerForm, error) {
	//TODO implement me
	panic("implement me")
}
