package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlItemCatalogRepository struct {
	conn Connection
}

func NewMSSQLItemCatalog(conn Connection) domain.ItemCatalogRepository {
	return &mssqlItemCatalogRepository{conn: conn}
}

// TODO: Implement functions
func (r *mssqlItemCatalogRepository) GetAll(ctx context.Context) ([]domain.ItemCatalog, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlItemCatalogRepository) GetByID(ctx context.Context, itemID uuid.UUID) (*domain.ItemCatalog, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlItemCatalogRepository) Create(ctx context.Context, item *domain.ItemCatalog) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlItemCatalogRepository) Update(ctx context.Context, item *domain.ItemCatalog) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlItemCatalogRepository) Delete(ctx context.Context, itemID uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlItemCatalogRepository) GetByCategory(ctx context.Context, category string) ([]domain.ItemCatalog, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlItemCatalogRepository) GetActive(ctx context.Context) ([]domain.ItemCatalog, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlItemCatalogRepository) SetActiveStatus(ctx context.Context, itemID uuid.UUID, isActive bool) error {
	//TODO implement me
	panic("implement me")
}
