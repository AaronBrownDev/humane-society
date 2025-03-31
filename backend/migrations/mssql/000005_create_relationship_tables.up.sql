USE HumaneSociety;
GO

-- Dog Prescription table
CREATE TABLE medical.DogPrescription (
    PrescriptionID INT IDENTITY(1, 1) NOT NULL,
    DogID UNIQUEIDENTIFIER NOT NULL,
    MedicineID INT NOT NULL,
    Dosage DECIMAL(5,2) NOT NULL,
    Frequency NVARCHAR(50) NULL,
    StartDate DATE NOT NULL DEFAULT GETDATE(),
    EndDate DATE NULL,
    Notes NVARCHAR(200) NULL,
    VetPrescriberID UNIQUEIDENTIFIER NULL,
    CONSTRAINT PK_DogPrescription PRIMARY KEY (PrescriptionID),
    CONSTRAINT FK_DogPrescription_Dog FOREIGN KEY (DogID)
        REFERENCES shelter.Dog(DogID)
        ON DELETE CASCADE,
    CONSTRAINT FK_DogPrescription_Medicine FOREIGN KEY (MedicineID)
        REFERENCES medical.Medicine(MedicineID)
        ON DELETE CASCADE,
    CONSTRAINT FK_DogPrescription_Veterinarian FOREIGN KEY (VetPrescriberID)
        REFERENCES people.Veterinarian(VeterinarianID)
        ON DELETE NO ACTION
);
GO

-- Create unique constraint for dog and medicine combination
CREATE UNIQUE INDEX UQ_DogPrescription_DogMedicine
    ON medical.DogPrescription(DogID, MedicineID, StartDate);
GO

-- Pet Owner's Pets table
CREATE TABLE people.PetOwnerPets (
    PetID INT IDENTITY(1,1) NOT NULL,
    PetOwnerID UNIQUEIDENTIFIER NOT NULL,
    Name NVARCHAR(50) NOT NULL,
    Type NVARCHAR(20) DEFAULT 'Dog',
    Breed NVARCHAR(50) NOT NULL,
    Sex VARCHAR(8) NOT NULL,
    OwnershipDate DATE NOT NULL,
    LivingEnvironment VARCHAR(7) NOT NULL,
    CONSTRAINT PK_PetOwnerPets PRIMARY KEY (PetID),
    CONSTRAINT FK_PetOwnerPets_PetOwner FOREIGN KEY (PetOwnerID)
        REFERENCES people.PetOwner(PetOwnerID)
        ON DELETE CASCADE,
    CONSTRAINT CK_PetOwnerPets_LivingEnvironment CHECK (LivingEnvironment IN ('Indoor', 'Outdoor', 'Both')),
    CONSTRAINT CK_PetOwnerPets_Sex CHECK (Sex IN ('Male', 'Female', 'Intersex'))
);
GO

-- Surrender Form table
CREATE TABLE shelter.SurrenderForm (
    SurrenderFormID INT IDENTITY(1,1) NOT NULL,
    SurrendererID UNIQUEIDENTIFIER NOT NULL,
    SubmissionDate DATETIME2(0) NOT NULL DEFAULT GETDATE(),
    DogName NVARCHAR(50) NOT NULL,
    DogAge INT NOT NULL,
    WeightInPounds DECIMAL(5,2) NOT NULL,
    Sex VARCHAR(8) NOT NULL,
    Breed NVARCHAR(50) NULL,
    Color NVARCHAR(30) NULL,
    LivingEnvironment VARCHAR(7) NOT NULL,
    OwnershipDate DATE NOT NULL,
    VeterinarianID UNIQUEIDENTIFIER NULL,
    LastVetVisitDate DATE NULL,
    IsGoodWithChildren BIT NOT NULL DEFAULT 0,
    IsGoodWithDogs BIT NOT NULL DEFAULT 0,
    IsGoodWithCats BIT NOT NULL DEFAULT 0,
    IsGoodWithStrangers BIT NOT NULL DEFAULT 0,
    IsHouseTrained BIT NOT NULL DEFAULT 0,
    IsSterilized BIT NOT NULL DEFAULT 0,
    MicroChipNumber VARCHAR(15) NULL,
    MedicalProblems NVARCHAR(500) NULL,
    BiteHistory NVARCHAR(500) NULL,
    SurrenderReason NVARCHAR(500) NOT NULL,
    ProcessedByVolunteerID UNIQUEIDENTIFIER NULL,
    ProcessingDate DATETIME2(0) NULL,
    ResultingDogID UNIQUEIDENTIFIER NULL,
    Status VARCHAR(20) DEFAULT 'Pending',
    CONSTRAINT PK_SurrenderForm PRIMARY KEY (SurrenderFormID),
    CONSTRAINT FK_SurrenderForm_PetSurrenderer FOREIGN KEY (SurrendererID)
        REFERENCES people.PetSurrenderer(SurrendererID)
        ON DELETE CASCADE,
    CONSTRAINT FK_SurrenderForm_Veterinarian FOREIGN KEY (VeterinarianID)
        REFERENCES people.Veterinarian(VeterinarianID)
        ON DELETE NO ACTION,
    CONSTRAINT FK_SurrenderForm_Volunteer FOREIGN KEY (ProcessedByVolunteerID)
        REFERENCES people.Volunteer(VolunteerID)
        ON DELETE NO ACTION,
    CONSTRAINT FK_SurrenderForm_Dog FOREIGN KEY (ResultingDogID)
        REFERENCES shelter.Dog(DogID)
        ON DELETE NO ACTION,
    CONSTRAINT CK_SurrenderForm_Sex CHECK (Sex IN ('Male', 'Female', 'Intersex')),
    CONSTRAINT CK_SurrenderForm_LivingEnvironment CHECK (LivingEnvironment IN ('Indoor', 'Outdoor', 'Both')),
    CONSTRAINT CK_SurrenderForm_Status CHECK (Status IN ('Pending', 'Approved', 'Rejected', 'Completed'))
);
GO

-- Adoption Form
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
GO

-- Create unique constraint on adopter and dog
CREATE UNIQUE INDEX UQ_AdoptionForm_AdopterDog
    ON shelter.AdoptionForm(AdopterID, DogID)
    WHERE Status IN ('Pending', 'HomeVisitScheduled', 'Approved');
GO

-- Volunteer Form table
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
GO

-- Volunteer Schedule table
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
GO