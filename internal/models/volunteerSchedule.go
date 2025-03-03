package models

import (
	"github.com/google/uuid"
	"time"
)

type VolunteerSchedule struct {
	ScheduleID      int       `json:"scheduleId" db:"ScheduleID"`
	VolunteerID     uuid.UUID `json:"volunteerId" db:"VolunteerID"`
	ScheduleDate    time.Time `json:"scheduleDate" db:"ScheduleDate"`
	StartTime       time.Time `json:"startTime" db:"StartTime"`
	EndTime         time.Time `json:"endTime" db:"EndTime"`
	TaskDescription string    `json:"taskDescription" db:"TaskDescription"`
	Status          string    `json:"status" db:"Status"`
}
