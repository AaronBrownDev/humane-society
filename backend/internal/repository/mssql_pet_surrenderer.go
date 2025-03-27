package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlPetSurrendererRepository struct {
	conn Connection
}

func NewMSSQLSurrenderer(conn Connection) domain.PetSurrendererRepository {
	return &mssqlPetSurrendererRepository{
		conn: conn,
	}
}

// TODO: Implement functions
func (m mssqlPetSurrendererRepository) GetAll(ctx context.Context) ([]domain.PetSurrenderer, error) {
	//TODO implement me
	panic("implement me")
}

func (m mssqlPetSurrendererRepository) GetByID(ctx context.Context, surrendererID uuid.UUID) (*domain.PetSurrenderer, error) {
	//TODO implement me
	panic("implement me")
}

func (m mssqlPetSurrendererRepository) Create(ctx context.Context, surrenderer *domain.PetSurrenderer) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlPetSurrendererRepository) Update(ctx context.Context, surrenderer *domain.PetSurrenderer) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlPetSurrendererRepository) Delete(ctx context.Context, surrendererID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlPetSurrendererRepository) GetSurrenderHistory(ctx context.Context, surrendererID uuid.UUID) ([]domain.SurrenderForm, error) {
	//TODO implement me
	panic("implement me")
}
