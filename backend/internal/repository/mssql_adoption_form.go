package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlAdoptionFormRepository struct {
	conn Connection
}

func (m mssqlAdoptionFormRepository) CompleteForm(ctx context.Context, formID int, volunteerID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func NewMSSQLAdoptionForm(conn Connection) domain.AdoptionFormRepository {
	return &mssqlAdoptionFormRepository{conn: conn}
}

// TODO Implement functions
func (m mssqlAdoptionFormRepository) GetAll(ctx context.Context) ([]domain.AdoptionForm, error) {
	//TODO implement me
	panic("implement me")
}

func (m mssqlAdoptionFormRepository) GetPending(ctx context.Context) ([]domain.AdoptionForm, error) {
	//TODO implement me
	panic("implement me")
}

func (m mssqlAdoptionFormRepository) GetByID(ctx context.Context, formID int) (*domain.AdoptionForm, error) {
	//TODO implement me
	panic("implement me")
}

func (m mssqlAdoptionFormRepository) Create(ctx context.Context, form *domain.AdoptionForm) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlAdoptionFormRepository) Update(ctx context.Context, form *domain.AdoptionForm) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlAdoptionFormRepository) Delete(ctx context.Context, formID int) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlAdoptionFormRepository) ApproveForm(ctx context.Context, formID int, volunteerID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlAdoptionFormRepository) RejectForm(ctx context.Context, formID int, volunteerID uuid.UUID, reason string) error {
	//TODO implement me
	panic("implement me")
}
