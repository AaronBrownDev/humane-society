package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlVeterinarianRepository struct {
	conn Connection
}

func NewMSSQLVeterinarian(conn Connection) domain.VeterinarianRepository {
	return &mssqlVeterinarianRepository{conn: conn}
}

// TODO: Implement functions
func (r *mssqlVeterinarianRepository) GetAll(ctx context.Context) ([]domain.Veterinarian, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVeterinarianRepository) GetByID(ctx context.Context, veterinarianID uuid.UUID) (*domain.Veterinarian, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVeterinarianRepository) Create(ctx context.Context, veterinarian *domain.Veterinarian) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVeterinarianRepository) Update(ctx context.Context, veterinarian *domain.Veterinarian) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVeterinarianRepository) Delete(ctx context.Context, veterinarianID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
