package models

import (
	"github.com/google/uuid"
	"time"
)

type Volunteer struct {
	VolunteerID           uuid.UUID `json:"volunteerId" db:"VolunteerID"`
	VolunteerPosition     string    `json:"volunteerPosition" db:"VolunteerPosition"`
	StartDate             time.Time `json:"startDate" db:"StartDate"`
	EndDate               time.Time `json:"endDate" db:"EndDate"`
	EmergencyContactName  string    `json:"emergencyContactName" db:"EmergencyContactName"`
	EmergencyContactPhone string    `json:"emergencyContactPhone" db:"EmergencyContactPhone"`
	IsActive              bool      `json:"isActive" db:"IsActive"`
}
