CREATE TABLE people.VolunteerSchedule (
    ScheduleID INT IDENTITY(1,1) NOT NULL,
    VolunteerID UNIQUEIDENTIFIER NOT NULL,
    ScheduleDate DATE NOT NULL,
    StartTime TIME NOT NULL,
    EndTime TIME NOT NULL,
    TaskDescription NVARCHAR(100) NULL,
    Status VARCHAR(20) DEFAULT 'Scheduled',
    CONSTRAINT PK_VolunteerSchedule PRIMARY KEY (ScheduleID),
    CONSTRAINT FK_VolunteerSchedule_Volunteer FOREIGN KEY (VolunteerID)
        REFERENCES people.Volunteer(VolunteerID)
        ON DELETE CASCADE,
    CONSTRAINT CK_VolunteerSchedule_Status CHECK (Status IN ('Scheduled', 'Completed', 'Cancelled', 'NoShow')),
    CONSTRAINT CK_VolunteerSchedule_Times CHECK (EndTime > StartTime)
);