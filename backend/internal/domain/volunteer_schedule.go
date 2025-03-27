package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

// VolunteerSchedule represents a volunteer's scheduled shift
type VolunteerSchedule struct {
	ScheduleID      int       `json:"scheduleId"`
	VolunteerID     uuid.UUID `json:"volunteerId"`
	ScheduleDate    time.Time `json:"scheduleDate"`
	StartTime       time.Time `json:"startTime"`
	EndTime         time.Time `json:"endTime"`
	TaskDescription string    `json:"taskDescription"`
	Status          string    `json:"status"`
}

// VolunteerScheduleRepository defines the interface for volunteer schedule data access operations
type VolunteerScheduleRepository interface {
	GetAll(ctx context.Context) ([]VolunteerSchedule, error)
	GetByID(ctx context.Context, scheduleID int) (*VolunteerSchedule, error)
	GetByVolunteerID(ctx context.Context, volunteerID uuid.UUID) ([]VolunteerSchedule, error)
	GetByDateRange(ctx context.Context, startDate, endDate time.Time) ([]VolunteerSchedule, error)
	Create(ctx context.Context, schedule *VolunteerSchedule) error
	Update(ctx context.Context, schedule *VolunteerSchedule) error
	Delete(ctx context.Context, scheduleID int) error

	// Domain-specific operations
	GetByStatus(ctx context.Context, status string) ([]VolunteerSchedule, error)
	UpdateStatus(ctx context.Context, scheduleID int, status string) error
	GetUpcomingSchedules(ctx context.Context, volunteerID uuid.UUID) ([]VolunteerSchedule, error)
	GetSchedulesByTask(ctx context.Context, taskDescription string) ([]VolunteerSchedule, error)
}
