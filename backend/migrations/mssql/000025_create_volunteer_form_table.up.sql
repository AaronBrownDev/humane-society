CREATE TABLE shelter.VolunteerForm (
    VolunteerFormID INT IDENTITY(1,1) NOT NULL,
    ApplicantID UNIQUEIDENTIFIER NOT NULL,
    SubmissionDate DATETIME2(0) NOT NULL DEFAULT GETDATE(),
    SupportsAnimalWelfareEducation BIT NOT NULL,
    AvailableShifts NVARCHAR(500),
    SupportsResponsibleBreeding BIT NOT NULL,
    AcceptsCleaningDuties BIT NOT NULL,
    AcceptsDogCare BIT NOT NULL,
    HasDogAllergies BIT NOT NULL,
    HasPhysicalLimitations BIT NOT NULL,
    IsForCommunityService BIT NOT NULL,
    RequiredServiceHours INT NULL,
    ReferralSource NVARCHAR(500) NOT NULL,
    CommentsAndQuestions NVARCHAR(500) NULL,
    ProcessedByVolunteerID UNIQUEIDENTIFIER NULL,
    ProcessingDate DATETIME2(0) NULL,
    Status VARCHAR(20) NOT NULL DEFAULT 'Pending',
    RejectionReason NVARCHAR(200) NULL,
    CONSTRAINT PK_VolunteerForm PRIMARY KEY (VolunteerFormID),
    CONSTRAINT FK_VolunteerForm_Person FOREIGN KEY (ApplicantID)
        REFERENCES people.Person(PersonID)
        ON DELETE CASCADE,
    CONSTRAINT FK_VolunteerForm_Volunteer FOREIGN KEY (ProcessedByVolunteerID)
        REFERENCES people.Volunteer(VolunteerID)
        ON DELETE NO ACTION,
    CONSTRAINT CK_VolunteerForm_Status CHECK (Status IN ('Pending', 'Approved', 'Rejected', 'Completed'))
);