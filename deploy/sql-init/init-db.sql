-- =============================================
-- Humane Society of Northwest Louisiana Database
-- Combined Migration Script
-- =============================================

-- =============================================
-- 1. CREATE DATABASE AND SCHEMAS
-- =============================================
USE master;
GO

IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = 'HumaneSociety')
    BEGIN
        CREATE DATABASE HumaneSociety;
    END
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

-- =============================================
-- 2. CREATE AUDIT SYSTEM
-- =============================================
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

-- =============================================
-- 3. CREATE BASE TABLES
-- =============================================
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

-- =============================================
-- 4. CREATE PERSON SUBTYPES
-- =============================================
-- Veterinarian table
CREATE TABLE people.Veterinarian (
                                     VeterinarianID UNIQUEIDENTIFIER NOT NULL,
                                     CONSTRAINT PK_Veterinarian PRIMARY KEY (VeterinarianID),
                                     CONSTRAINT FK_Veterinarian_Person FOREIGN KEY (VeterinarianID)
                                         REFERENCES people.Person(PersonID)
                                         ON DELETE CASCADE
);
GO

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

-- Pet Surrenderer table
CREATE TABLE people.PetSurrenderer (
                                       SurrendererID UNIQUEIDENTIFIER NOT NULL,
                                       CONSTRAINT PK_PetSurrenderer PRIMARY KEY (SurrendererID),
                                       CONSTRAINT FK_PetSurrenderer_Person FOREIGN KEY (SurrendererID)
                                           REFERENCES people.Person(PersonID)
                                           ON DELETE CASCADE
);
GO

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

-- =============================================
-- 5. CREATE RELATIONSHIP TABLES
-- =============================================
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

-- =============================================
-- 6. CREATE VIEWS
-- =============================================
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

-- =============================================
-- 7. CREATE STORED PROCEDURES
-- =============================================
-- Inserts the person into the person table
CREATE OR ALTER PROCEDURE InsertPerson
    @PersonID UNIQUEIDENTIFIER,
    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE,
    @PhysicalAddress NVARCHAR(225),
    @MailingAddress NVARCHAR(225),
    @EmailAddress NVARCHAR(100),
    @PhoneNumber NVARCHAR(20)
AS
BEGIN
    SET NOCOUNT ON;
    INSERT INTO people.Person (PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, EmailAddress, PhoneNumber)
    VALUES (@PersonID, @FirstName, @LastName, @BirthDate, @PhysicalAddress, @MailingAddress, @EmailAddress, @PhoneNumber);
END;
GO

-- Inserts a veterinarian, including the person record if needed
CREATE OR ALTER PROCEDURE InsertVeterinarian
    @VeterinarianID UNIQUEIDENTIFIER,
    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE,
    @PhysicalAddress NVARCHAR(225),
    @MailingAddress NVARCHAR(225),
    @EmailAddress NVARCHAR(100),
    @PhoneNumber NVARCHAR(20)
AS
BEGIN
    SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;
        IF NOT EXISTS (SELECT 1 FROM people.Person WHERE PersonID = @VeterinarianID)
            BEGIN
                EXEC InsertPerson
                     @PersonID = @VeterinarianID,
                     @FirstName = @FirstName,
                     @LastName = @LastName,
                     @BirthDate = @BirthDate,
                     @PhysicalAddress = @PhysicalAddress,
                     @MailingAddress = @MailingAddress,
                     @EmailAddress = @EmailAddress,
                     @PhoneNumber = @PhoneNumber;

                PRINT 'Person inserted successfully';
            END

        INSERT INTO people.Veterinarian(VeterinarianID)
        VALUES (@VeterinarianID)

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION
        PRINT 'Error occurred: ' + ERROR_MESSAGE();
    END CATCH;
END;
GO

-- Inserts Pet Owner, including the person record if needed
CREATE OR ALTER PROCEDURE InsertPetOwner
    @PetOwnerID UNIQUEIDENTIFIER,
    @VeterinarianID UNIQUEIDENTIFIER,
    @PetsSterilized BIT,
    @PetsVaccinated BIT,
    @HeartWormPreventionFromVet BIT,
    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE,
    @PhysicalAddress NVARCHAR(255),
    @MailingAddress NVARCHAR(255),
    @EmailAddress NVARCHAR(100),
    @PhoneNumber NVARCHAR(20)
AS
BEGIN
    SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;
        -- Ensures the vetID already exists
        IF NOT EXISTS (SELECT 1 FROM people.Veterinarian WHERE VeterinarianID = @VeterinarianID)
            BEGIN
                INSERT INTO people.Veterinarian (VeterinarianID)
                VALUES (@VeterinarianID);
            END;

        -- Ensure the Person record is inserted first
        IF NOT EXISTS (SELECT 1 FROM people.Person WHERE PersonID = @PetOwnerID)
            BEGIN
                EXEC InsertPerson
                     @PersonID = @PetOwnerID,
                     @FirstName = @FirstName,
                     @LastName = @LastName,
                     @BirthDate = @BirthDate,
                     @PhysicalAddress = @PhysicalAddress,
                     @MailingAddress = @MailingAddress,
                     @EmailAddress = @EmailAddress,
                     @PhoneNumber = @PhoneNumber;
                PRINT 'Person inserted successfully';
            END

        -- Now insert into PetOwner (which references PersonID)
        INSERT INTO people.PetOwner (PetOwnerID, VeterinarianID, HasSterilizedPets, HasVaccinatedPets, UsesVetHeartWormPrevention)
        VALUES (@PetOwnerID, @VeterinarianID, @PetsSterilized, @PetsVaccinated, @HeartWormPreventionFromVet);

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        PRINT 'Error occurred: ' + ERROR_MESSAGE();
    END CATCH;
END;
GO

-- Inserts into the adopter table, including the person record if needed
CREATE OR ALTER PROCEDURE InsertAdopter
    @AdopterID UNIQUEIDENTIFIER,
    @PetAllergies BIT,
    @HaveSurrendered BIT,
    @HomeStatus VARCHAR(20),
    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE,
    @PhysicalAddress NVARCHAR(225),
    @MailingAddress NVARCHAR(225),
    @EmailAddress NVARCHAR(100),
    @PhoneNumber NVARCHAR(20)
AS
BEGIN
    SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;
        --Ensures the person doesn't already exist
        IF NOT EXISTS (SELECT 1 FROM people.Person WHERE PersonID = @AdopterID)
            BEGIN
                EXEC InsertPerson
                     @PersonID = @AdopterID,
                     @FirstName = @FirstName,
                     @LastName = @LastName,
                     @BirthDate = @BirthDate,
                     @PhysicalAddress = @PhysicalAddress,
                     @MailingAddress = @MailingAddress,
                     @EmailAddress = @EmailAddress,
                     @PhoneNumber = @PhoneNumber;
            END
        --Inserts the adopter
        INSERT INTO people.Adopter(AdopterID, HasPetAllergies, HasSurrenderedPets, HomeStatus)
        VALUES(@AdopterID, @PetAllergies, @HaveSurrendered, @HomeStatus);

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        PRINT 'Error occurred: ' + ERROR_MESSAGE();
    END CATCH;
END;
GO

-- Inserts a volunteer, including the person record if needed
CREATE OR ALTER PROCEDURE InsertVolunteer
    @VolunteerID UNIQUEIDENTIFIER,
    @VolunteerPosition NVARCHAR(50),
    @StartDate DATE,
    @EndDate DATE,
    @EmergencyContactName NVARCHAR(100),
    @EmergencyContactPhone NVARCHAR(20),
    @IsActive BIT,
    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE,
    @PhysicalAddress NVARCHAR(225),
    @MailingAddress NVARCHAR(225),
    @EmailAddress NVARCHAR(100),
    @PhoneNumber NVARCHAR(20)
AS
BEGIN
    SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;
        --Ensures the person does not already exist in the person table
        IF NOT EXISTS (SELECT 1 FROM people.Person WHERE PersonID = @VolunteerID)
            BEGIN
                EXEC InsertPerson
                     @PersonID = @VolunteerID,
                     @FirstName = @FirstName,
                     @LastName = @LastName,
                     @BirthDate = @BirthDate,
                     @PhysicalAddress = @PhysicalAddress,
                     @MailingAddress = @MailingAddress,
                     @EmailAddress = @EmailAddress,
                     @PhoneNumber = @PhoneNumber;
            END
        -- Inserts in the volunteer table
        INSERT INTO people.Volunteer(VolunteerID, VolunteerPosition, StartDate, EndDate, EmergencyContactName, EmergencyContactPhone, IsActive)
        VALUES (@VolunteerID, @VolunteerPosition, @StartDate, @EndDate, @EmergencyContactName, @EmergencyContactPhone, @IsActive);

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        PRINT 'Error occurred: ' + ERROR_MESSAGE();
    END CATCH;
END;
GO

-- =============================================
-- 8. INSERT SAMPLE DATA
-- =============================================
BEGIN TRY
    -- Start a transaction for data consistency
    BEGIN TRANSACTION;

    -- Declare all variables first
    DECLARE @PersonID1 UNIQUEIDENTIFIER, @PersonID2 UNIQUEIDENTIFIER, @PersonID3 UNIQUEIDENTIFIER, @PersonID4 UNIQUEIDENTIFIER, @PersonID5 UNIQUEIDENTIFIER;
    DECLARE @PersonID6 UNIQUEIDENTIFIER, @PersonID7 UNIQUEIDENTIFIER, @PersonID8 UNIQUEIDENTIFIER, @PersonID9 UNIQUEIDENTIFIER, @PersonID10 UNIQUEIDENTIFIER;
    DECLARE @PersonID11 UNIQUEIDENTIFIER, @PersonID12 UNIQUEIDENTIFIER, @PersonID13 UNIQUEIDENTIFIER, @PersonID14 UNIQUEIDENTIFIER, @PersonID15 UNIQUEIDENTIFIER;
    DECLARE @Dog1 UNIQUEIDENTIFIER, @Dog2 UNIQUEIDENTIFIER, @Dog3 UNIQUEIDENTIFIER, @Dog4 UNIQUEIDENTIFIER, @Dog5 UNIQUEIDENTIFIER;
    DECLARE @Dog6 UNIQUEIDENTIFIER, @Dog7 UNIQUEIDENTIFIER, @Dog8 UNIQUEIDENTIFIER;
    DECLARE @MedicineID1 INT, @MedicineID2 INT, @MedicineID3 INT, @MedicineID4 INT, @MedicineID5 INT;
    DECLARE @ItemID1 UNIQUEIDENTIFIER, @ItemID2 UNIQUEIDENTIFIER, @ItemID3 UNIQUEIDENTIFIER;
    DECLARE @ItemID4 UNIQUEIDENTIFIER, @ItemID5 UNIQUEIDENTIFIER, @ItemID6 UNIQUEIDENTIFIER;
    DECLARE @ItemID7 UNIQUEIDENTIFIER, @ItemID11 UNIQUEIDENTIFIER, @ItemID13 UNIQUEIDENTIFIER;
    DECLARE @ItemID14 UNIQUEIDENTIFIER;

    -- Pre-assign all GUIDs
    SET @PersonID1 = NEWID();
    SET @PersonID2 = NEWID();
    SET @PersonID3 = NEWID();
    SET @PersonID4 = NEWID();
    SET @PersonID5 = NEWID();
    SET @PersonID6 = NEWID();
    SET @PersonID7 = NEWID();
    SET @PersonID8 = NEWID();
    SET @PersonID9 = NEWID();
    SET @PersonID10 = NEWID();
    SET @PersonID11 = NEWID();
    SET @PersonID12 = NEWID();
    SET @PersonID13 = NEWID();
    SET @PersonID14 = NEWID();
    SET @PersonID15 = NEWID();

    SET @Dog1 = NEWID();
    SET @Dog2 = NEWID();
    SET @Dog3 = NEWID();
    SET @Dog4 = NEWID();
    SET @Dog5 = NEWID();
    SET @Dog6 = NEWID();
    SET @Dog7 = NEWID();
    SET @Dog8 = NEWID();

    SET @ItemID1 = NEWID();
    SET @ItemID2 = NEWID();
    SET @ItemID3 = NEWID();
    SET @ItemID4 = NEWID();
    SET @ItemID5 = NEWID();
    SET @ItemID6 = NEWID();
    SET @ItemID7 = NEWID();
    SET @ItemID11 = NEWID();
    SET @ItemID13 = NEWID();
    SET @ItemID14 = NEWID();

    -- Insert Person records
    INSERT INTO people.Person (PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, EmailAddress, PhoneNumber)
    VALUES
        (@PersonID1, 'John', 'Smith', '1985-04-12', '123 Main St, Springfield, IL 62701', '123 Main St, Springfield, IL 62701', 'john.smith@email.com', '217-555-1234'),
        (@PersonID2, 'Jane', 'Doe', '1990-08-24', '456 Oak Ave, Springfield, IL 62702', '456 Oak Ave, Springfield, IL 62702', 'jane.doe@email.com', '217-555-5678'),
        (@PersonID3, 'Robert', 'Johnson', '1978-02-15', '789 Pine Rd, Springfield, IL 62703', '789 Pine Rd, Springfield, IL 62703', 'robert.johnson@email.com', '217-555-9012'),
        (@PersonID4, 'Emily', 'Wilson', '1992-11-30', '321 Elm St, Springfield, IL 62704', '321 Elm St, Springfield, IL 62704', 'emily.wilson@email.com', '217-555-3456'),
        (@PersonID5, 'Michael', 'Brown', '1983-07-08', '654 Maple Dr, Springfield, IL 62705', '654 Maple Dr, Springfield, IL 62705', 'michael.brown@email.com', '217-555-7890'),
        (@PersonID6, 'Sarah', 'Davis', '1995-01-17', '987 Cedar Ln, Springfield, IL 62706', '987 Cedar Ln, Springfield, IL 62706', 'sarah.davis@email.com', '217-555-2345'),
        (@PersonID7, 'David', 'Miller', '1975-09-22', '135 Walnut Ct, Springfield, IL 62707', '135 Walnut Ct, Springfield, IL 62707', 'david.miller@email.com', '217-555-6789'),
        (@PersonID8, 'Jennifer', 'Taylor', '1988-03-05', '246 Birch Blvd, Springfield, IL 62708', '246 Birch Blvd, Springfield, IL 62708', 'jennifer.taylor@email.com', '217-555-0123'),
        (@PersonID9, 'James', 'Anderson', '1980-12-14', '357 Spruce Way, Springfield, IL 62709', '357 Spruce Way, Springfield, IL 62709', 'james.anderson@email.com', '217-555-4567'),
        (@PersonID10, 'Jessica', 'Thomas', '1993-06-27', '468 Aspen Pl, Springfield, IL 62710', '468 Aspen Pl, Springfield, IL 62710', 'jessica.thomas@email.com', '217-555-8901'),
        (@PersonID11, 'William', 'Jackson', '1982-10-09', '579 Fir Dr, Springfield, IL 62711', '579 Fir Dr, Springfield, IL 62711', 'william.jackson@email.com', '217-555-2345'),
        (@PersonID12, 'Amanda', 'White', '1991-05-03', '680 Hemlock St, Springfield, IL 62712', '680 Hemlock St, Springfield, IL 62712', 'amanda.white@email.com', '217-555-6789'),
        (@PersonID13, 'Christopher', 'Harris', '1977-08-19', '791 Locust Ave, Springfield, IL 62713', '791 Locust Ave, Springfield, IL 62713', 'christopher.harris@email.com', '217-555-0123'),
        (@PersonID14, 'Nicole', 'Martin', '1994-02-11', '802 Juniper Rd, Springfield, IL 62714', '802 Juniper Rd, Springfield, IL 62714', 'nicole.martin@email.com', '217-555-4567'),
        (@PersonID15, 'Daniel', 'Thompson', '1986-07-23', '913 Magnolia Dr, Springfield, IL 62715', '913 Magnolia Dr, Springfield, IL 62715', 'daniel.thompson@email.com', '217-555-8901');

    -- Assign roles to people
    -- Adopters
    INSERT INTO people.Adopter (AdopterID, HasPetAllergies, HasSurrenderedPets, HomeStatus)
    VALUES
        (@PersonID1, 0, 0, 'Approved'),
        (@PersonID4, 1, 0, 'Approved'),
        (@PersonID8, 0, 0, 'Pending'),
        (@PersonID10, 0, 1, 'Approved'),
        (@PersonID12, 0, 0, 'Rejected');

    -- Veterinarians
    INSERT INTO people.Veterinarian (VeterinarianID)
    VALUES
        (@PersonID3),
        (@PersonID11);

    -- Pet Surrenderers
    INSERT INTO people.PetSurrenderer (SurrendererID)
    VALUES
        (@PersonID5),
        (@PersonID9),
        (@PersonID10); -- Jessica is both a surrenderer and adopter

    -- Pet Owners
    INSERT INTO people.PetOwner (PetOwnerID, VeterinarianID, HasSterilizedPets, HasVaccinatedPets, UsesVetHeartWormPrevention)
    VALUES
        (@PersonID1, @PersonID3, 1, 1, 1), -- John is both an adopter and pet owner
        (@PersonID6, @PersonID11, 1, 1, 0),
        (@PersonID12, @PersonID3, 0, 1, 0); -- Amanda is both an adopter and pet owner

    -- Volunteers
    INSERT INTO people.Volunteer (VolunteerID, VolunteerPosition, StartDate, EndDate, EmergencyContactName, EmergencyContactPhone, IsActive)
    VALUES
        (@PersonID2, 'Dog Walker', '2023-01-15', NULL, 'Robert Johnson', '217-555-9012', 1),
        (@PersonID7, 'Cleaner', '2023-03-10', NULL, 'Mary Miller', '217-555-1234', 1),
        (@PersonID13, 'Admin', '2022-11-05', NULL, 'Lisa Harris', '217-555-5678', 1),
        (@PersonID14, 'Fundraiser', '2023-05-22', NULL, 'Mark Martin', '217-555-9876', 1),
        (@PersonID15, 'Dog Walker', '2022-08-18', '2023-06-30', 'Karen Thompson', '217-555-4321', 0);

    -- Pet Owner Pets
    INSERT INTO people.PetOwnerPets (PetOwnerID, Name, Type, Breed, Sex, OwnershipDate, LivingEnvironment)
    VALUES
        (@PersonID1, 'Max', 'Dog', 'Labrador Retriever', 'Male', '2020-05-10', 'Indoor'),
        (@PersonID1, 'Bella', 'Dog', 'Beagle', 'Female', '2021-03-15', 'Both'),
        (@PersonID6, 'Charlie', 'Dog', 'Golden Retriever', 'Male', '2019-08-22', 'Both'),
        (@PersonID6, 'Lucy', 'Cat', 'Siamese', 'Female', '2020-11-30', 'Indoor'),
        (@PersonID12, 'Cooper', 'Dog', 'German Shepherd', 'Male', '2022-01-05', 'Both');

    -- Create dogs
    INSERT INTO shelter.Dog (DogID, Name, IntakeDate, EstimatedBirthDate, Breed, Sex, Color, CageNumber, IsAdopted)
    VALUES
        (@Dog1, 'Rocky', '2023-02-15', '2021-06-10', 'Pit Bull Mix', 'Male', 'Brown', 1, 0),
        (@Dog2, 'Luna', '2023-03-22', '2022-01-05', 'Labrador Retriever', 'Female', 'Black', 2, 0),
        (@Dog3, 'Rex', '2023-01-08', '2020-11-12', 'German Shepherd', 'Male', 'Black and Tan', 3, 0),
        (@Dog4, 'Daisy', '2023-04-05', '2022-03-18', 'Beagle', 'Female', 'Tricolor', 4, 0),
        (@Dog5, 'Buddy', '2023-05-11', '2019-08-30', 'Golden Retriever', 'Male', 'Golden', 5, 0),
        (@Dog6, 'Mia', '2023-02-28', '2020-05-15', 'Boxer', 'Female', 'Brindle', 6, 1),
        (@Dog7, 'Duke', '2023-03-17', '2022-02-20', 'Husky', 'Male', 'Gray and White', 7, 0),
        (@Dog8, 'Stella', '2023-05-03', '2021-10-08', 'Australian Shepherd', 'Female', 'Merle', 8, 0);

    -- Medicine
    INSERT INTO medical.Medicine (Name, Manufacturer, Description, DosageUnit)
    VALUES
        ('Heartgard Plus', 'Boehringer Ingelheim', 'Heartworm prevention', 'Tablet'),
        ('Frontline Plus', 'Merial', 'Flea and tick prevention', 'ml'),
        ('Rimadyl', 'Zoetis', 'Pain relief and anti-inflammatory', 'mg'),
        ('Clavamox', 'Zoetis', 'Antibiotic', 'mg'),
        ('Cerenia', 'Zoetis', 'Anti-nausea and vomiting', 'mg');

    -- Store medicine IDs by directly querying
    SELECT @MedicineID1 = MedicineID FROM medical.Medicine WHERE Name = 'Heartgard Plus';
    SELECT @MedicineID2 = MedicineID FROM medical.Medicine WHERE Name = 'Frontline Plus';
    SELECT @MedicineID3 = MedicineID FROM medical.Medicine WHERE Name = 'Rimadyl';
    SELECT @MedicineID4 = MedicineID FROM medical.Medicine WHERE Name = 'Clavamox';
    SELECT @MedicineID5 = MedicineID FROM medical.Medicine WHERE Name = 'Cerenia';

    -- Dog Prescriptions
    INSERT INTO medical.DogPrescription (DogID, MedicineID, Dosage, Frequency, StartDate, EndDate, Notes, VetPrescriberID)
    VALUES
        (@Dog1, @MedicineID1, 1, 'Once monthly', '2023-02-20', NULL, 'Monthly heartworm prevention', @PersonID3),
        (@Dog1, @MedicineID2, 0.5, 'Once monthly', '2023-02-20', NULL, 'Monthly flea and tick prevention', @PersonID3),
        (@Dog3, @MedicineID4, 250, 'Twice daily', '2023-01-10', '2023-01-24', 'Treating respiratory infection', @PersonID11),
        (@Dog5, @MedicineID3, 75, 'Twice daily', '2023-05-15', '2023-05-29', 'Joint pain management', @PersonID3),
        (@Dog7, @MedicineID5, 16, 'Once daily', '2023-03-20', '2023-03-25', 'Motion sickness during transport', @PersonID11);

    -- Item Catalog
    INSERT INTO shelter.ItemCatalog (ItemID, Name, Category, Description, MinimumQuantity, IsActive)
    VALUES
        (@ItemID1, 'Dog Food - Adult', 'Food', 'Premium adult dog food', 5, 1),
        (@ItemID2, 'Dog Food - Puppy', 'Food', 'Premium puppy food', 3, 1),
        (@ItemID3, 'Dog Treats', 'Food', 'Training treats', 10, 1),
        (@ItemID4, 'Leash - Standard', 'Equipment', '6 foot standard leash', 8, 1),
        (@ItemID5, 'Collar - Small', 'Equipment', 'Small dog collar', 5, 1),
        (@ItemID6, 'Collar - Medium', 'Equipment', 'Medium dog collar', 5, 1),
        (@ItemID7, 'Collar - Large', 'Equipment', 'Large dog collar', 5, 1),
        (NEWID(), 'Bed - Small', 'Bedding', 'Small dog bed', 3, 1),
        (NEWID(), 'Bed - Large', 'Bedding', 'Large dog bed', 3, 1),
        (NEWID(), 'Blanket', 'Bedding', 'Fleece dog blanket', 10, 1),
        (@ItemID11, 'Cleaning Solution', 'Cleaning', 'Pet-safe cleaning solution', 2, 1),
        (NEWID(), 'Bleach', 'Cleaning', 'Disinfectant', 2, 1),
        (@ItemID13, 'Paper Towels', 'Cleaning', 'Paper towel rolls', 15, 1),
        (@ItemID14, 'Dog Shampoo', 'Grooming', 'Gentle dog shampoo', 5, 1),
        (NEWID(), 'Nail Clippers', 'Grooming', 'Dog nail clippers', 3, 1);

    -- Supplies
    INSERT INTO shelter.Supply (ItemID, Quantity, StorageLocation, ExpirationDate, BatchNumber, AcquisitionDate)
    VALUES
        (@ItemID1, 15, 'Main Storage', '2024-05-20', 'BT12345', '2023-05-15'),
        (@ItemID1, 10, 'Main Storage', '2024-08-10', 'BT12346', '2023-08-05'),
        (@ItemID2, 8, 'Main Storage', '2024-06-15', 'BT23456', '2023-06-10'),
        (@ItemID3, 25, 'Front Desk', '2024-12-01', 'BT34567', '2023-07-20'),
        (@ItemID4, 12, 'Equipment Room', NULL, 'EQ12345', '2023-03-15'),
        (@ItemID5, 6, 'Equipment Room', NULL, 'EQ23456', '2023-04-02'),
        (@ItemID6, 9, 'Equipment Room', NULL, 'EQ34567', '2023-04-02'),
        (@ItemID7, 7, 'Equipment Room', NULL, 'EQ45678', '2023-04-02'),
        (@ItemID11, 4, 'Janitorial Closet', '2025-01-15', 'CL12345', '2023-01-10'),
        (@ItemID13, 20, 'Janitorial Closet', NULL, 'CL23456', '2023-05-01'),
        (@ItemID14, 10, 'Grooming Area', '2024-09-30', 'GR12345', '2023-06-25');

    -- Surrender Form
    INSERT INTO shelter.SurrenderForm (
        SurrendererID, SubmissionDate, DogName, DogAge, WeightInPounds, Sex, Breed, Color,
        LivingEnvironment, OwnershipDate, VeterinarianID, LastVetVisitDate, IsGoodWithChildren, IsGoodWithDogs,
        IsGoodWithCats, IsGoodWithStrangers, IsHouseTrained, IsSterilized, MicroChipNumber,
        MedicalProblems, BiteHistory, SurrenderReason, ProcessedByVolunteerID, ProcessingDate,
        ResultingDogID, Status
    )
    VALUES
        (@PersonID5, '2023-02-10', 'Rocky', 2, 45.5, 'Male', 'Pit Bull Mix', 'Brown',
         'Both', '2021-07-15', @PersonID3, '2022-11-20', 1, 1, 0, 1, 1, 1, '985120035487621',
         'None', 'No bite history', 'Moving to apartment that doesn''t allow dogs', @PersonID13, '2023-02-15',
         @Dog1, 'Completed'),

        (@PersonID9, '2023-03-18', 'Luna', 1, 35.2, 'Female', 'Labrador Retriever', 'Black',
         'Indoor', '2022-02-10', @PersonID11, '2022-12-15', 1, 1, 1, 1, 0, 0, NULL,
         'None', 'No bite history', 'Unable to care for puppy with new baby', @PersonID13, '2023-03-22',
         @Dog2, 'Completed'),

        (@PersonID10, '2023-05-01', 'Stella', 1, 32.4, 'Female', 'Australian Shepherd', 'Merle',
         'Both', '2022-01-15', @PersonID3, '2023-03-10', 1, 0, 0, 0, 1, 1, '985120035782145',
         'Sensitive stomach, needs special diet', 'No bite history', 'Too high energy for our lifestyle', @PersonID13, '2023-05-03',
         @Dog8, 'Completed'),

        (@PersonID5, '2023-05-28', 'Baxter', 3, 28.6, 'Male', 'Cocker Spaniel', 'Blonde',
         'Indoor', '2020-04-30', @PersonID11, '2023-04-15', 1, 1, 1, 1, 1, 1, '985120035963214',
         'Slight hearing loss in left ear', 'No bite history', 'Cannot afford care anymore', NULL, NULL,
         NULL, 'Pending');

    -- Adoption Form
    INSERT INTO shelter.AdoptionForm (
        AdopterID, DogID, SubmissionDate, ProcessedByVolunteerID, ProcessingDate, Status, RejectionReason
    )
    VALUES
        (@PersonID1, @Dog5, '2023-05-20', @PersonID13, '2023-05-25', 'Approved', NULL),
        (@PersonID4, @Dog7, '2023-05-22', @PersonID13, '2023-05-27', 'Approved', NULL),
        (@PersonID8, @Dog4, '2023-05-15', NULL, NULL, 'Pending', NULL),
        (@PersonID10, @Dog3, '2023-05-18', @PersonID13, '2023-05-23', 'Pending', NULL),
        (@PersonID12, @Dog2, '2023-04-10', @PersonID13, '2023-04-15', 'Rejected', 'Home environment not suitable for this dog');

    -- Volunteer Form
    INSERT INTO shelter.VolunteerForm (
        ApplicantID, SubmissionDate, SupportsAnimalWelfareEducation, AvailableShifts,
        SupportsResponsibleBreeding, AcceptsCleaningDuties, AcceptsDogCare, HasDogAllergies, HasPhysicalLimitations,
        IsForCommunityService, RequiredServiceHours, ReferralSource, CommentsAndQuestions,
        ProcessedByVolunteerID, ProcessingDate, Status, RejectionReason
    )
    VALUES
        (@PersonID2, '2022-12-20', 1, 'Weekday mornings and weekends', 0, 1, 1, 0, 0, 0, NULL,
         'Through a friend who volunteers', 'Excited to help with the dogs!', @PersonID13, '2023-01-05', 'Approved', NULL),

        (@PersonID7, '2023-02-15', 1, 'Weekends only', 0, 1, 1, 0, 0, 0, NULL,
         'Local community board', 'I have experience with large dogs', @PersonID13, '2023-03-01', 'Approved', NULL),

        (@PersonID14, '2023-05-10', 1, 'Weekday evenings', 0, 1, 1, 0, 0, 0, NULL,
         'Social media post', 'I have fundraising experience', @PersonID13, '2023-05-15', 'Approved', NULL),

        (@PersonID15, '2022-07-25', 1, 'Flexible schedule', 0, 1, 1, 0, 0, 0, NULL,
         'Website', 'Looking to help out during my summer break', @PersonID13, '2022-08-05', 'Approved', NULL);

    -- Volunteer Schedule
    INSERT INTO people.VolunteerSchedule (
        VolunteerID, ScheduleDate, StartTime, EndTime, TaskDescription, Status
    )
    VALUES
        (@PersonID2, '2023-06-03', '09:00:00', '12:00:00', 'Dog walking and socialization', 'Scheduled'),
        (@PersonID2, '2023-06-10', '09:00:00', '12:00:00', 'Dog walking and socialization', 'Scheduled'),
        (@PersonID2, '2023-05-27', '09:00:00', '12:00:00', 'Dog walking and socialization', 'Completed'),
        (@PersonID2, '2023-05-20', '09:00:00', '12:00:00', 'Dog walking and socialization', 'Completed'),

        (@PersonID7, '2023-06-04', '10:00:00', '14:00:00', 'Kennel cleaning and laundry', 'Scheduled'),
        (@PersonID7, '2023-06-11', '10:00:00', '14:00:00', 'Kennel cleaning and laundry', 'Scheduled'),
        (@PersonID7, '2023-05-28', '10:00:00', '14:00:00', 'Kennel cleaning and laundry', 'Completed'),
        (@PersonID7, '2023-05-21', '10:00:00', '14:00:00', 'Kennel cleaning and laundry', 'Completed'),

        (@PersonID13, '2023-06-05', '13:00:00', '17:00:00', 'Office admin and adoption processing', 'Scheduled'),
        (@PersonID13, '2023-06-07', '13:00:00', '17:00:00', 'Office admin and adoption processing', 'Scheduled'),
        (@PersonID13, '2023-06-09', '13:00:00', '17:00:00', 'Office admin and adoption processing', 'Scheduled'),
        (@PersonID13, '2023-05-29', '13:00:00', '17:00:00', 'Office admin and adoption processing', 'Completed'),
        (@PersonID13, '2023-05-31', '13:00:00', '17:00:00', 'Office admin and adoption processing', 'Completed'),
        (@PersonID13, '2023-05-26', '13:00:00', '17:00:00', 'Office admin and adoption processing', 'Completed'),

        (@PersonID14, '2023-06-06', '17:00:00', '20:00:00', 'Fundraising call planning', 'Scheduled'),
        (@PersonID14, '2023-06-08', '17:00:00', '20:00:00', 'Fundraising call execution', 'Scheduled'),
        (@PersonID14, '2023-05-30', '17:00:00', '20:00:00', 'Fundraising call planning', 'Completed'),
        (@PersonID14, '2023-05-25', '17:00:00', '20:00:00', 'Fundraising call execution', 'Completed'),

        (@PersonID15, '2023-05-15', '09:00:00', '12:00:00', 'Dog walking and socialization', 'NoShow'),
        (@PersonID15, '2023-05-08', '09:00:00', '12:00:00', 'Dog walking and socialization', 'Completed'),
        (@PersonID15, '2023-05-01', '09:00:00', '12:00:00', 'Dog walking and socialization', 'Completed');

    -- Sample audit log entries
    INSERT INTO audit.ChangeLog (
        TableName, PrimaryKeyColumn, PrimaryKeyValue, ColumnName, OldValue, NewValue, ChangeDate, ChangedBy, AuditActionType
    )
    VALUES
        ('shelter.Dog', 'DogID', CONVERT(NVARCHAR(36), @Dog6), 'IsAdopted', '0', '1', '2023-06-01 10:15:23', 'system', 'U'),
        ('shelter.Dog', 'DogID', CONVERT(NVARCHAR(36), @Dog6), 'AdoptionDate', 'NULL', '2023-06-01', '2023-06-01 10:15:23', 'system', 'U'),
        ('shelter.AdoptionForm', 'AdoptionFormID', '1', 'Status', 'Approved', 'Completed', '2023-06-01 10:15:23', 'system', 'U'),
        ('medical.DogPrescription', 'PrescriptionID', '3', 'EndDate', '2023-01-24', '2023-01-31', '2023-01-24 14:22:10', 'system', 'U'),
        ('shelter.Supply', 'SupplyID', '1', 'Quantity', '20', '15', '2023-05-28 09:45:33', 'system', 'U');

    -- Mark one of the adoption forms as complete and the dog as adopted
    UPDATE shelter.Dog SET IsAdopted = 1 WHERE DogID = @Dog6;
    UPDATE shelter.AdoptionForm SET Status = 'Completed' WHERE AdopterID = @PersonID1 AND DogID = @Dog5;

    -- Commit the transaction if everything succeeded
    COMMIT TRANSACTION;
    PRINT 'Sample data inserted successfully.';
END TRY
BEGIN CATCH
    -- Roll back the transaction if an error occurred
    IF @@TRANCOUNT > 0
        ROLLBACK TRANSACTION;

    -- Print error information
    PRINT 'Error occurred: ' + ERROR_MESSAGE();
    PRINT 'Error line: ' + CAST(ERROR_LINE() AS VARCHAR(10));
    PRINT 'Error number: ' + CAST(ERROR_NUMBER() AS VARCHAR(10));
END CATCH;
GO

-- =============================================
-- 9. CREATE DATABASE ROLES
-- =============================================
-- Create database roles
CREATE ROLE HumaneSociety_Admin;
GO

CREATE ROLE HumaneSociety_Staff;
GO

CREATE ROLE HumaneSociety_Volunteer;
GO

CREATE ROLE HumaneSociety_ReadOnly;
GO

-- =============================================
-- 10. ASSIGN PERMISSIONS TO ROLES
-- =============================================
-- Admin role gets full control
GRANT CONTROL ON SCHEMA::shelter TO HumaneSociety_Admin;
GRANT CONTROL ON SCHEMA::people TO HumaneSociety_Admin;
GRANT CONTROL ON SCHEMA::medical TO HumaneSociety_Admin;
GRANT CONTROL ON SCHEMA::audit TO HumaneSociety_Admin;
GO

-- Staff role gets data modification rights on most tables
GRANT SELECT, INSERT, UPDATE, DELETE ON SCHEMA::shelter TO HumaneSociety_Staff;
GRANT SELECT, INSERT, UPDATE, DELETE ON SCHEMA::people TO HumaneSociety_Staff;
GRANT SELECT, INSERT, UPDATE, DELETE ON SCHEMA::medical TO HumaneSociety_Staff;
GRANT SELECT ON SCHEMA::audit TO HumaneSociety_Staff;
GO

-- Volunteer role gets limited rights
GRANT SELECT ON SCHEMA::shelter TO HumaneSociety_Volunteer;
GRANT SELECT ON SCHEMA::people TO HumaneSociety_Volunteer;
GRANT SELECT ON SCHEMA::medical TO HumaneSociety_Volunteer;

-- Allow volunteers to update specific tables
GRANT UPDATE ON shelter.Dog TO HumaneSociety_Volunteer;
GRANT INSERT, UPDATE ON shelter.AdoptionForm TO HumaneSociety_Volunteer;
GRANT INSERT, UPDATE ON shelter.SurrenderForm TO HumaneSociety_Volunteer;
GRANT INSERT, UPDATE, DELETE ON people.VolunteerSchedule TO HumaneSociety_Volunteer;
GO

-- ReadOnly role gets read-only access
GRANT SELECT ON SCHEMA::shelter TO HumaneSociety_ReadOnly;
GRANT SELECT ON SCHEMA::people TO HumaneSociety_ReadOnly;
GRANT SELECT ON SCHEMA::medical TO HumaneSociety_ReadOnly;
GO

-- =============================================
-- 11. CREATE AUDIT TRIGGERS
-- =============================================
-- Dog table audit trigger
CREATE TRIGGER shelter_Dog_Audit ON shelter.Dog
    AFTER INSERT, UPDATE, DELETE
    AS
BEGIN
    SET NOCOUNT ON;

    DECLARE @Action CHAR(1);

    -- Figure out what operation is happening (Insert, Update, Delete)
    IF EXISTS (SELECT * FROM inserted) AND EXISTS (SELECT * FROM deleted)
        SET @Action = 'U'; -- Update
    ELSE IF EXISTS (SELECT * FROM inserted)
        SET @Action = 'I'; -- Insert
    ELSE
        SET @Action = 'D'; -- Delete

    -- For INSERT operations
    IF @Action = 'I'
        BEGIN
            -- Insert records into the audit log for each inserted row
            INSERT INTO audit.ChangeLog
            (TableName, PrimaryKeyColumn, PrimaryKeyValue,
             ColumnName, OldValue, NewValue, AuditActionType)
            SELECT
                'shelter.Dog', -- Table Name
                'DogID',      -- Primary Key Column
                CONVERT(NVARCHAR(36), i.DogID), -- Primary Key Value
                'All Columns', -- Column Name (simplified to just track "all columns")
                NULL,         -- Old Value (none for INSERT)
                'New Dog: ' + i.Name, -- Simplified new value
                'I'           -- Action Type
            FROM
                inserted i;
        END

    -- For UPDATE operations
    IF @Action = 'U'
        BEGIN
            -- Log name changes
            IF UPDATE(Name)
                BEGIN
                    INSERT INTO audit.ChangeLog
                    (TableName, PrimaryKeyColumn, PrimaryKeyValue,
                     ColumnName, OldValue, NewValue, AuditActionType)
                    SELECT
                        'shelter.Dog',
                        'DogID',
                        CONVERT(NVARCHAR(36), i.DogID),
                        'Name',
                        d.Name,
                        i.Name,
                        'U'
                    FROM
                        inserted i
                            JOIN deleted d ON i.DogID = d.DogID
                    WHERE
                        i.Name <> d.Name;
                END

            -- Log adoption status changes
            IF UPDATE(IsAdopted)
                BEGIN
                    INSERT INTO audit.ChangeLog
                    (TableName, PrimaryKeyColumn, PrimaryKeyValue,
                     ColumnName, OldValue, NewValue, AuditActionType)
                    SELECT
                        'shelter.Dog',
                        'DogID',
                        CONVERT(NVARCHAR(36), i.DogID),
                        'IsAdopted',
                        CASE WHEN d.IsAdopted = 1 THEN 'Yes' ELSE 'No' END,
                        CASE WHEN i.IsAdopted = 1 THEN 'Yes' ELSE 'No' END,
                        'U'
                    FROM
                        inserted i
                            JOIN deleted d ON i.DogID = d.DogID
                    WHERE
                        i.IsAdopted <> d.IsAdopted;
                END
        END

    -- For DELETE operations
    IF @Action = 'D'
        BEGIN
            INSERT INTO audit.ChangeLog
            (TableName, PrimaryKeyColumn, PrimaryKeyValue,
             ColumnName, OldValue, NewValue, AuditActionType)
            SELECT
                'shelter.Dog',
                'DogID',
                CONVERT(NVARCHAR(36), d.DogID),
                'All Columns',
                'Dog Deleted: ' + d.Name,
                NULL,
                'D'
            FROM
                deleted d;
        END
END;
GO

-- =============================================
-- 12. CREATE ADOPTION FORM PROCEDURES
-- =============================================
CREATE OR ALTER PROCEDURE ApproveAdoptionForm
    @AdoptionFormID INT,
    @ProcessedByVolunteerID UNIQUEIDENTIFIER
AS
BEGIN
    SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;

        DECLARE @DogID UNIQUEIDENTIFIER;
        DECLARE @CurrentStatus VARCHAR(20);

        -- Get the dog ID and current status from the adoption form
        SELECT @DogID = DogID, @CurrentStatus = Status
        FROM shelter.AdoptionForm
        WHERE AdoptionFormID = @AdoptionFormID;

        -- Only proceed if form exists and is in a valid state for approval
        IF @DogID IS NULL
            BEGIN
                RAISERROR('Adoption form not found', 16, 1);
                RETURN;
            END

        IF @CurrentStatus NOT IN ('Pending', 'HomeVisitScheduled')
            BEGIN
                RAISERROR('Adoption form cannot be approved from current status: %s', 16, 1, @CurrentStatus);
                RETURN;
            END

        -- Update the adoption form to approved
        UPDATE shelter.AdoptionForm
        SET Status = 'Approved',
            ProcessedByVolunteerID = @ProcessedByVolunteerID,
            ProcessingDate = GETDATE()
        WHERE AdoptionFormID = @AdoptionFormID;

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        THROW;
    END CATCH;
END;
GO

CREATE OR ALTER PROCEDURE CompleteAdoptionForm
    @AdoptionFormID INT,
    @ProcessedByVolunteerID UNIQUEIDENTIFIER
AS
BEGIN
    SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;

        DECLARE @DogID UNIQUEIDENTIFIER;
        DECLARE @CurrentStatus VARCHAR(20);

        -- Get the dog ID and current status from the adoption form
        SELECT @DogID = DogID, @CurrentStatus = Status
        FROM shelter.AdoptionForm
        WHERE AdoptionFormID = @AdoptionFormID;

        -- Only proceed if form exists and is in Approved status
        IF @DogID IS NULL
            BEGIN
                RAISERROR('Adoption form not found', 16, 1);
                RETURN;
            END

        IF @CurrentStatus <> 'Approved'
            BEGIN
                RAISERROR('Adoption form must be approved before completion', 16, 1);
                RETURN;
            END

        -- Update the adoption form to completed
        UPDATE shelter.AdoptionForm
        SET Status = 'Completed',
            ProcessedByVolunteerID = @ProcessedByVolunteerID,
            ProcessingDate = GETDATE()
        WHERE AdoptionFormID = @AdoptionFormID;

        -- Mark the dog as adopted
        UPDATE shelter.Dog
        SET IsAdopted = 1
        WHERE DogID = @DogID;

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        THROW;
    END CATCH;
END;
GO

CREATE OR ALTER PROCEDURE RejectAdoptionForm
    @AdoptionFormID INT,
    @ProcessedByVolunteerID UNIQUEIDENTIFIER,
    @RejectionReason NVARCHAR(200)
AS
BEGIN
    SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;

        DECLARE @CurrentStatus VARCHAR(20);

        -- Get current status from the adoption form
        SELECT @CurrentStatus = Status
        FROM shelter.AdoptionForm
        WHERE AdoptionFormID = @AdoptionFormID;

        -- Only proceed if form exists and is in a valid state for rejection
        IF @CurrentStatus IS NULL
            BEGIN
                RAISERROR('Adoption form not found', 16, 1);
                RETURN;
            END

        IF @CurrentStatus NOT IN ('Pending', 'HomeVisitScheduled', 'Approved')
            BEGIN
                RAISERROR('Adoption form cannot be rejected from current status: %s', 16, 1, @CurrentStatus);
                RETURN;
            END

        -- Update the adoption form to rejected
        UPDATE shelter.AdoptionForm
        SET Status = 'Rejected',
            ProcessedByVolunteerID = @ProcessedByVolunteerID,
            ProcessingDate = GETDATE(),
            RejectionReason = @RejectionReason
        WHERE AdoptionFormID = @AdoptionFormID;

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        THROW;
    END CATCH;
END;
GO