package repository

import (
	"database/sql"
	"github.com/AaronBrownDev/HumaneSociety/internal/domain"
	"github.com/google/uuid"
)

type VolunteerRepository interface {
	// Volunteer Form CRUD
	GetAllVolunteerForms() ([]domain.VolunteerForm, error)
	GetVolunteerForm(formID int) (*domain.VolunteerForm, error)
	InsertVolunteerForm(form *domain.VolunteerForm) error
	UpdateVolunteerForm(form *domain.VolunteerForm) error
	DeleteVolunteerForm(formID int) error

	// Volunteer CRUD
	GetAllVolunteers() ([]domain.Volunteer, error)
	GetVolunteer(volunteerID uuid.UUID) (*domain.Volunteer, error)
	InsertVolunteer(volunteer *domain.Volunteer) error
	UpdateVolunteer(volunteer *domain.Volunteer) error
	DeleteVolunteer(volunteerID uuid.UUID) error

	// Volunteer Schedule
	GetAllSchedules() ([]domain.VolunteerSchedule, error)
	GetSchedule(scheduleID int) (*domain.VolunteerSchedule, error)
	InsertSchedule(schedule *domain.VolunteerSchedule) error
	UpdateSchedule(schedule *domain.VolunteerSchedule) error
	DeleteSchedule(scheduleID int) error

	GetVolunteerSchedules(volunteerID uuid.UUID) ([]domain.VolunteerSchedule, error)
}

type SQLVolunteerRepository struct {
	db *sql.DB
}

func NewSQLVolunteerRepository(db *sql.DB) VolunteerRepository {
	return &SQLVolunteerRepository{
		db: db,
	}
}

// TODO: Implement functions
