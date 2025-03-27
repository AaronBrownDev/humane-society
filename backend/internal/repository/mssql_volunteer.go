package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlVolunteerRepository struct {
	conn Connection
}

func NewMSSQLVolunteer(conn Connection) domain.VolunteerRepository {
	return &mssqlVolunteerRepository{conn: conn}
}

// TODO: Implement functions
func (r *mssqlVolunteerRepository) GetAll(ctx context.Context) ([]domain.Volunteer, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerRepository) GetActive(ctx context.Context) ([]domain.Volunteer, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerRepository) GetByID(ctx context.Context, volunteerID uuid.UUID) (*domain.Volunteer, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerRepository) Create(ctx context.Context, volunteer *domain.Volunteer) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerRepository) Update(ctx context.Context, volunteer *domain.Volunteer) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerRepository) Delete(ctx context.Context, volunteerID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerRepository) SetActiveStatus(ctx context.Context, volunteerID uuid.UUID, isActive bool) error {
	//TODO implement me
	panic("implement me")
}
