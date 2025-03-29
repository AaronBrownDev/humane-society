# Humane Society Database Schema

This document contains the Data Definition Language (DDL) statements for creating the Humane Society of Northwest Louisiana Management System database schema.

## Database and Schema Creation

```sql
CREATE DATABASE HumaneSociety;
GO

USE HumaneSociety;
GO

-- SCHEMAS
CREATE SCHEMA shelter;
GO

CREATE SCHEMA medical;
GO

CREATE SCHEMA people;
GO

CREATE SCHEMA audit;
GO
```

## Audit System

The audit system tracks changes to database records for accountability and historical tracking.

```sql
-- AUDIT SYSTEM
CREATE TABLE audit.ChangeLog (
    LogID INT IDENTITY(1,1) NOT NULL,
    TableName NVARCHAR(128) NOT NULL,
    PrimaryKeyColumn NVARCHAR(128) NOT NULL,
    PrimaryKeyValue NVARCHAR(36) NOT NULL,
    ColumnName NVARCHAR(128) NOT NULL,
    OldValue NVARCHAR(MAX) NULL,
    NewValue NVARCHAR(MAX) NULL,
    ChangeDate DATETIME2(0) NOT NULL DEFAULT GETDATE(),
    ChangedBy NVARCHAR(128) NOT NULL DEFAULT SYSTEM_USER,
    AuditActionType CHAR(1) NOT NULL, -- I = Insert, U = Update, D = Delete
    CONSTRAINT PK_ChangeLog PRIMARY KEY (LogID),
    CONSTRAINT CK_ChangeLog_AuditActionType CHECK (AuditActionType IN ('I', 'U', 'D'))
);
GO
```

## Base Tables

### Person Table

The core entity representing any individual in the system.

```sql
-- Person Table
CREATE TABLE people.Person (
    PersonID UNIQUEIDENTIFIER NOT NULL,
    FirstName NVARCHAR(50) NOT NULL,
    LastName NVARCHAR(50) NOT NULL,
    BirthDate DATE NOT NULL,
    PhysicalAddress NVARCHAR(150) NOT NULL,
    MailingAddress NVARCHAR(150) NOT NULL,
    EmailAddress VARCHAR(100) NULL,
    PhoneNumber VARCHAR(20) NULL,
    CONSTRAINT PK_Person PRIMARY KEY (PersonID)
);
GO

-- Index on LastName, FirstName for name searches
CREATE INDEX IX_Person_Name ON people.Person(LastName, FirstName);
GO
```

### Dog Table

Represents dogs housed at the shelter.

```sql
-- Dog Table
CREATE TABLE shelter.Dog (
    DogID UNIQUEIDENTIFIER NOT NULL,
    Name NVARCHAR(50) NOT NULL,
    IntakeDate DATE NOT NULL,
    EstimatedBirthDate DATE NOT NULL,
    Breed NVARCHAR(50) NOT NULL,
    Sex VARCHAR(8) NOT NULL,
    Color NVARCHAR(30) NOT NULL,
    CageNumber INT NOT NULL,
    IsAdopted BIT NOT NULL DEFAULT 0,
    CONSTRAINT PK_Dog PRIMARY KEY (DogID),
    CONSTRAINT CK_Dog_Sex CHECK (Sex IN ('Male', 'Female', 'Intersex'))
);
GO

-- Index for looking up available dogs
CREATE INDEX IX_Dog_Adoption ON shelter.Dog(IsAdopted);
GO
```

### Medicine Table

Tracks medications used at the shelter.

```sql
-- Medicine table
CREATE TABLE medical.Medicine (
    MedicineID INT IDENTITY(1,1) NOT NULL,
    Name NVARCHAR(50) NOT NULL,
    Manufacturer NVARCHAR(50) NOT NULL,
    Description NVARCHAR(200) NULL,
    DosageUnit NVARCHAR(20) NULL,
    CONSTRAINT PK_Medicine PRIMARY KEY (MedicineID)
);
GO
```

### Inventory Management Tables

Tables for managing shelter supplies and items.

```sql
-- Item Catalog
CREATE TABLE shelter.ItemCatalog (
    ItemID UNIQUEIDENTIFIER NOT NULL,
    Name NVARCHAR(50) NOT NULL,
    Category NVARCHAR(30) NOT NULL,
    Description NVARCHAR(200) NULL,
    MinimumQuantity INT NOT NULL DEFAULT 0,
    IsActive BIT NOT NULL DEFAULT 1,
    CONSTRAINT PK_ItemCatalog PRIMARY KEY (ItemID),
    CONSTRAINT UK_ItemCatalog_Name UNIQUE (Name)
);
GO

-- Supply inventory
CREATE TABLE shelter.Supply (
    SupplyID INT IDENTITY(1,1) NOT NULL,
    ItemID UNIQUEIDENTIFIER NOT NULL,
    Quantity INT NOT NULL,
    StorageLocation NVARCHAR(50) NULL,
    ExpirationDate DATE NULL,
    BatchNumber NVARCHAR(50) NULL,
    AcquisitionDate DATE NULL DEFAULT GETDATE(),
    CONSTRAINT PK_Supply PRIMARY KEY (SupplyID),
    CONSTRAINT FK_Supply_ItemCatalog FOREIGN KEY (ItemID)
        REFERENCES shelter.ItemCatalog(ItemID)
        ON DELETE CASCADE,
    CONSTRAINT CK_Supply_Quantity CHECK (Quantity >= 0)
);
GO

-- Index for finding items by catalog ID
CREATE INDEX IX_Supply_ItemID ON shelter.Supply(ItemID);
GO
```

## Person Subtypes

Various roles a person can have in the system, implemented using single-table inheritance.

### Veterinarian

```sql
-- Veterinarian table
CREATE TABLE people.Veterinarian (
    VeterinarianID UNIQUEIDENTIFIER NOT NULL,
    CONSTRAINT PK_Veterinarian PRIMARY KEY (VeterinarianID),
    CONSTRAINT FK_Veterinarian_Person FOREIGN KEY (VeterinarianID)
        REFERENCES people.Person(PersonID)
        ON DELETE CASCADE
);
GO
```

### Adopter

```sql
-- Adopter table
CREATE TABLE people.Adopter (
    AdopterID UNIQUEIDENTIFIER NOT NULL,
    HasPetAllergies BIT NOT NULL DEFAULT 0,
    HasSurrenderedPets BIT NOT NULL DEFAULT 0,
    HomeStatus VARCHAR(20) NOT NULL DEFAULT 'Pending',
    CONSTRAINT PK_Adopter PRIMARY KEY (AdopterID),
    CONSTRAINT FK_Adopter_Person FOREIGN KEY (AdopterID)
        REFERENCES people.Person(PersonID)
        ON DELETE CASCADE,
    CONSTRAINT CK_Adopter_HomeStatus CHECK (HomeStatus IN ('Pending', 'Approved', 'Rejected'))
);
GO
```

### Pet Surrenderer

```sql
-- Pet Surrenderer table
CREATE TABLE people.PetSurrenderer (
    SurrendererID UNIQUEIDENTIFIER NOT NULL,
    CONSTRAINT PK_PetSurrenderer PRIMARY KEY (SurrendererID),
    CONSTRAINT FK_PetSurrenderer_Person FOREIGN KEY (SurrendererID)
        REFERENCES people.Person(PersonID)
        ON DELETE CASCADE
);
GO
```

### Pet Owner

```sql
-- Pet Owner table
CREATE TABLE people.PetOwner (
    PetOwnerID UNIQUEIDENTIFIER NOT NULL,
    VeterinarianID UNIQUEIDENTIFIER NULL,
    HasSterilizedPets BIT NOT NULL DEFAULT 0,
    HasVaccinatedPets BIT NOT NULL DEFAULT 0,
    UsesVetHeartWormPrevention BIT NOT NULL DEFAULT 0,
    CONSTRAINT PK_PetOwner PRIMARY KEY (PetOwnerID),
    CONSTRAINT FK_PetOwner_Person FOREIGN KEY (PetOwnerID)
        REFERENCES people.Person(PersonID)
        ON DELETE CASCADE,
    CONSTRAINT FK_PetOwner_Veterinarian FOREIGN KEY (VeterinarianID)
        REFERENCES people.Veterinarian(VeterinarianID)
        ON DELETE NO ACTION
);
GO
```

### Volunteer

```sql
-- Volunteer table
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
GO
```

## Relationship Tables

### Dog Prescription

Tracks medications prescribed to shelter dogs.

```sql
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
```

### Pet Owner's Pets

Tracks pets owned by registered pet owners.

```sql
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
```

## Forms

### Surrender Form

Records information about surrendered pets.

```sql
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
```

### Adoption Form

Tracks adoption applications.

```sql
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
```

### Volunteer Form

Tracks volunteer applications.

```sql
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
```

### Volunteer Schedule

Manages volunteer shifts and assignments.

```sql
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
```

## Views

### Available Dogs View

Provides a filtered view of dogs available for adoption with age calculation.

```sql
-- Create a view for available dogs
CREATE VIEW shelter.AvailableDogs AS
SELECT
    d.DogID,
    d.Name, 
    d.IntakeDate,
    d.EstimatedBirthDate,
    DATEDIFF(YEAR, d.EstimatedBirthDate, GETDATE()) AS AgeInYears,
    d.Breed,
    d.Sex,
    d.Color,
    d.CageNumber,
    d.IsAdopted
FROM
    shelter.Dog AS d
WHERE
    d.IsAdopted = 0;
GO
```