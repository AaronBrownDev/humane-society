USE HumaneSociety;
GO

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