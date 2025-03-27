package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type mssqlDogPrescriptionRepository struct {
	conn Connection
}

func NewMSSQLDogPrescription(conn Connection) domain.DogPrescriptionRepository {
	return &mssqlDogPrescriptionRepository{conn: conn}
}

func (m mssqlDogPrescriptionRepository) GetAll(ctx context.Context) ([]domain.DogPrescription, error) {
	//TODO implement me
	panic("implement me")
}

func (m mssqlDogPrescriptionRepository) GetByID(ctx context.Context, prescriptionID int) (*domain.DogPrescription, error) {
	//TODO implement me
	panic("implement me")
}

func (m mssqlDogPrescriptionRepository) GetByDogID(ctx context.Context, dogID uuid.UUID) ([]domain.DogPrescription, error) {
	//TODO implement me
	panic("implement me")
}

func (m mssqlDogPrescriptionRepository) Create(ctx context.Context, prescription *domain.DogPrescription) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlDogPrescriptionRepository) Update(ctx context.Context, prescription *domain.DogPrescription) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlDogPrescriptionRepository) Delete(ctx context.Context, prescriptionID int) error {
	//TODO implement me
	panic("implement me")
}

func (m mssqlDogPrescriptionRepository) GetActivePrescriptions(ctx context.Context, dogID uuid.UUID) ([]domain.DogPrescription, error) {
	//TODO implement me
	panic("implement me")
}
