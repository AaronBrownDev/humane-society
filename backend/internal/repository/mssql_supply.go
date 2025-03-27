package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlSupplyRepository struct {
	conn Connection
}

func NewMSSQLSupply(conn Connection) domain.SupplyRepository {
	return &mssqlSupplyRepository{
		conn: conn,
	}
}

// TODO: implement functions

func (m mssqlSupplyRepository) GetAll(ctx context.Context) ([]domain.Supply, error) {
	//TODO implement me
	panic("implement me")
}

func (m mssqlSupplyRepository) GetByID(ctx context.Context, supplyID int) (*domain.Supply, error) {
	//TODO implement me
	panic("implement me")
}

func (m mssqlSupplyRepository) GetByItemID(ctx context.Context, itemID uuid.UUID) ([]domain.Supply, error) {
	//TODO implement me
	panic("implement me")
}

func (m mssqlSupplyRepository) Create(ctx context.Context, supply *domain.Supply) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlSupplyRepository) Update(ctx context.Context, supply *domain.Supply) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlSupplyRepository) Delete(ctx context.Context, supplyID int) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlSupplyRepository) UpdateQuantity(ctx context.Context, supplyID int, quantity int) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlSupplyRepository) GetExpiringSoon(ctx context.Context, days int) ([]domain.Supply, error) {
	//TODO implement me
	panic("implement me")
}

func (m mssqlSupplyRepository) GetLowStock(ctx context.Context) ([]domain.Supply, error) {
	//TODO implement me
	panic("implement me")
}
