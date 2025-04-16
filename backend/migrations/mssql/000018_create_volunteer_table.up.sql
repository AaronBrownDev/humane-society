CREATE TABLE people.Volunteer (
    VolunteerID UNIQUEIDENTIFIER NOT NULL,
    VolunteerPosition NVARCHAR(50) NOT NULL,
    StartDate DATE NOT NULL DEFAULT GETDATE(),
    EndDate DATE NULL,
    EmergencyContactName NVARCHAR(100) NULL,
    EmergencyContactPhone VARCHAR(20) NULL,
    IsActive BIT NOT NULL DEFAULT 1,
    CONSTRAINT PK_Volunteer PRIMARY KEY (VolunteerID),
    CONSTRAINT FK_Volunteer_Person FOREIGN KEY (VolunteerID)
        REFERENCES people.Person(PersonID)
        ON DELETE CASCADE,
    CONSTRAINT CK_Volunteer_EmergencyContact CHECK
        ((EmergencyContactName IS NULL AND EmergencyContactPhone IS NULL) OR
        (EmergencyContactName IS NOT NULL AND EmergencyContactPhone IS NOT NULL))
);