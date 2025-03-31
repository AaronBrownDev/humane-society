package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlPetOwnerPetRepository struct {
	conn Connection
}

func NewMSSQLPetOwnerPet(conn Connection) domain.PetOwnerPetRepository {
	return &mssqlPetOwnerPetRepository{conn: conn}
}

// TODO: Implement functions
func (r *mssqlPetOwnerPetRepository) GetAll(ctx context.Context) ([]domain.PetOwnerPet, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlPetOwnerPetRepository) GetByID(ctx context.Context, petID int) (*domain.PetOwnerPet, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlPetOwnerPetRepository) GetByPetOwnerID(ctx context.Context, petOwnerID uuid.UUID) ([]domain.PetOwnerPet, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlPetOwnerPetRepository) Create(ctx context.Context, pet *domain.PetOwnerPet) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlPetOwnerPetRepository) Update(ctx context.Context, pet *domain.PetOwnerPet) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlPetOwnerPetRepository) Delete(ctx context.Context, petID int) error {
	//TODO implement me
	panic("implement me")
}
