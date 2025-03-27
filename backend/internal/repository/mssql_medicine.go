package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
)

type mssqlMedicineRepository struct {
	conn Connection
}

func NewMSSQLMedicine(conn Connection) domain.MedicineRepository {
	return &mssqlMedicineRepository{conn: conn}
}

// TODO: Implement functions
func (r *mssqlMedicineRepository) GetAll(ctx context.Context) ([]domain.Medicine, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlMedicineRepository) GetByID(ctx context.Context, medicineID int) (*domain.Medicine, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlMedicineRepository) Create(ctx context.Context, medicine *domain.Medicine) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlMedicineRepository) Update(ctx context.Context, medicine *domain.Medicine) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlMedicineRepository) Delete(ctx context.Context, medicineID int) error {
	//TODO implement me
	panic("implement me")
}
