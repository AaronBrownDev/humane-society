package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlPetOwnerRepository struct {
	conn Connection
}

func NewMSSQLPetOwner(conn Connection) domain.PetOwnerRepository {
	return &mssqlPetOwnerRepository{conn: conn}
}

// TODO: Implement functions
func (r *mssqlPetOwnerRepository) GetAll(ctx context.Context) ([]domain.PetOwner, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlPetOwnerRepository) GetByID(ctx context.Context, petOwnerID uuid.UUID) (*domain.PetOwner, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlPetOwnerRepository) Create(ctx context.Context, petOwner *domain.PetOwner) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlPetOwnerRepository) Update(ctx context.Context, petOwner *domain.PetOwner) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlPetOwnerRepository) Delete(ctx context.Context, petOwnerID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlPetOwnerRepository) GetByVeterinarianID(ctx context.Context, veterinarianID uuid.UUID) ([]domain.PetOwner, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlPetOwnerRepository) GetWithPets(ctx context.Context, petOwnerID uuid.UUID) (*domain.PetOwner, []domain.PetOwnerPet, error) {
	//TODO implement me
	panic("implement me")
}
