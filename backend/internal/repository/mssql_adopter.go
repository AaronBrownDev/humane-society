package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlAdopterRepository struct {
	conn Connection
}

func NewMSSQLAdopter(conn Connection) domain.AdopterRepository {
	return &mssqlAdopterRepository{conn: conn}
}

// TODO: Implement functions
func (r *mssqlAdopterRepository) GetAll(ctx context.Context) ([]domain.Adopter, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlAdopterRepository) GetByID(ctx context.Context, adopterID uuid.UUID) (*domain.Adopter, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlAdopterRepository) Create(ctx context.Context, adopter *domain.Adopter) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlAdopterRepository) Update(ctx context.Context, adopter *domain.Adopter) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlAdopterRepository) Delete(ctx context.Context, adopterID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
