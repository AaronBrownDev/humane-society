package repository

import (
	"database/sql"
	"github.com/AaronBrownDev/HumaneSociety/internal/models"
	"github.com/google/uuid"
)

type VolunteerRepository interface {
	// Volunteer Form CRUD
	GetAllVolunteerForms() ([]models.VolunteerForm, error)
	GetVolunteerForm(formID int) (*models.VolunteerForm, error)
	InsertVolunteerForm(form *models.VolunteerForm) error
	UpdateVolunteerForm(form *models.VolunteerForm) error
	DeleteVolunteerForm(formID int) error

	// Volunteer CRUD
	GetAllVolunteers() ([]models.Volunteer, error)
	GetVolunteer(volunteerID uuid.UUID) (*models.Volunteer, error)
	InsertVolunteer(volunteer *models.Volunteer) error
	UpdateVolunteer(volunteer *models.Volunteer) error
	DeleteVolunteer(volunteerID uuid.UUID) error

	// Volunteer Schedule
	GetAllSchedules() ([]models.VolunteerSchedule, error)
	GetSchedule(scheduleID int) (*models.VolunteerSchedule, error)
	InsertSchedule(schedule *models.VolunteerSchedule) error
	UpdateSchedule(schedule *models.VolunteerSchedule) error
	DeleteSchedule(scheduleID int) error

	GetVolunteerSchedules(volunteerID uuid.UUID) ([]models.VolunteerSchedule, error)
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
