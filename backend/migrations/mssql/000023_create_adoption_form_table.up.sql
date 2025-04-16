CREATE TABLE shelter.AdoptionForm (
    AdoptionFormID INT IDENTITY(1,1) NOT NULL,
    AdopterID UNIQUEIDENTIFIER NOT NULL,
    DogID UNIQUEIDENTIFIER NOT NULL,
    SubmissionDate DATETIME2(0) NOT NULL DEFAULT GETDATE(),
    ProcessedByVolunteerID UNIQUEIDENTIFIER NULL,
    ProcessingDate DATETIME2(0) NULL,
    Status VARCHAR(20) NOT NULL DEFAULT 'Pending',
    RejectionReason NVARCHAR(200) NULL,
    CONSTRAINT PK_AdoptionForm PRIMARY KEY (AdoptionFormID),
    CONSTRAINT FK_AdoptionForm_Adopter FOREIGN KEY (AdopterID)
        REFERENCES people.Adopter(AdopterID)
        ON DELETE CASCADE,
    CONSTRAINT FK_AdoptionForm_Dog FOREIGN KEY (DogID)
        REFERENCES shelter.Dog(DogID)
        ON DELETE CASCADE,
    CONSTRAINT FK_AdoptionForm_Volunteer FOREIGN KEY (ProcessedByVolunteerID)
        REFERENCES people.Volunteer(VolunteerID)
        ON DELETE NO ACTION,
    CONSTRAINT CK_AdoptionForm_Status CHECK (Status IN ('Pending', 'HomeVisitScheduled', 'Approved', 'Rejected', 'Completed'))
);