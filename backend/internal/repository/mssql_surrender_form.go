package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlSurrenderFormRepository struct {
	conn Connection
}

func NewMSSQLSurrenderForm(conn Connection) domain.SurrenderFormRepository {
	return &mssqlSurrenderFormRepository{conn: conn}
}

// TODO: Implement functions
func (r *mssqlSurrenderFormRepository) GetAll(ctx context.Context) ([]domain.SurrenderForm, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlSurrenderFormRepository) GetByID(ctx context.Context, surrenderFormID int) (*domain.SurrenderForm, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlSurrenderFormRepository) GetBySurrendererID(ctx context.Context, surrendererID uuid.UUID) ([]domain.SurrenderForm, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlSurrenderFormRepository) Create(ctx context.Context, form *domain.SurrenderForm) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlSurrenderFormRepository) Update(ctx context.Context, form *domain.SurrenderForm) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlSurrenderFormRepository) Delete(ctx context.Context, surrenderFormID int) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlSurrenderFormRepository) GetByStatus(ctx context.Context, status string) ([]domain.SurrenderForm, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlSurrenderFormRepository) Process(ctx context.Context, surrenderFormID int, volunteerID uuid.UUID, status string) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlSurrenderFormRepository) GetPendingForms(ctx context.Context) ([]domain.SurrenderForm, error) {
	//TODO implement me
	panic("implement me")
}
