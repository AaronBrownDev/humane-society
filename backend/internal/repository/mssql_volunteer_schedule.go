package repository

import (
	"context"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
	"time"
)

type mssqlVolunteerScheduleRepository struct {
	conn Connection
}

func NewMSSQLVolunteerSchedule(conn Connection) domain.VolunteerScheduleRepository {
	return &mssqlVolunteerScheduleRepository{conn: conn}
}

// TODO: Implement functions
func (r *mssqlVolunteerScheduleRepository) GetAll(ctx context.Context) ([]domain.VolunteerSchedule, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerScheduleRepository) GetByID(ctx context.Context, scheduleID int) (*domain.VolunteerSchedule, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerScheduleRepository) GetByVolunteerID(ctx context.Context, volunteerID uuid.UUID) ([]domain.VolunteerSchedule, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerScheduleRepository) GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]domain.VolunteerSchedule, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerScheduleRepository) Create(ctx context.Context, schedule *domain.VolunteerSchedule) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerScheduleRepository) Update(ctx context.Context, schedule *domain.VolunteerSchedule) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerScheduleRepository) Delete(ctx context.Context, scheduleID int) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerScheduleRepository) GetByStatus(ctx context.Context, status string) ([]domain.VolunteerSchedule, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerScheduleRepository) UpdateStatus(ctx context.Context, scheduleID int, status string) error {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerScheduleRepository) GetUpcomingSchedules(ctx context.Context, volunteerID uuid.UUID) ([]domain.VolunteerSchedule, error) {
	//TODO implement me
	panic("implement me")
}

func (r *mssqlVolunteerScheduleRepository) GetSchedulesByTask(ctx context.Context, taskDescription string) ([]domain.VolunteerSchedule, error) {
	//TODO implement me
	panic("implement me")
}
