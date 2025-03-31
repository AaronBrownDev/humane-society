USE HumaneSociety;
GO

-- Inserts the person into the person table
CREATE PROCEDURE INSERTPerson
    @PersonID UNIQUEIDENTIFIER,
    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE,
    @PhysicalAddress NVARCHAR(225),
    @MailingAddress NVARCHAR (225),
    @EmailAddress NVARCHAR(100),
    @PhoneNumber NVARCHAR(20)

AS
BEGIN
    SET NOCOUNT ON;
    INSERT INTO people.Person (PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, EmailAddress, PhoneNumber)
    VALUES ( @PersonID, @FirstName, @LastName, @BirthDate, @PhysicalAddress, @MailingAddress, @EmailAddress, @PhoneNumber);

END;
GO

CREATE PROCEDURE INSERTVeterinarian
    @VeterinarianID UNIQUEIDENTIFIER,

    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE,
    @PhysicalAddress NVARCHAR(225),
    @MailingAddress NVARCHAR (225),
    @EmailAddress NVARCHAR(100),
    @PhoneNumber NVARCHAR(20)

AS
BEGIN
    SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;
        IF NOT EXISTS (SELECT 1 FROM people.Person WHERE PersonID = @VeterinarianID)
            BEGIN
                EXEC INSERTPerson
                     @PersonID = @VeterinarianID ,
                     @FirstName = @FirstName,
                     @LastName = @LastName,
                     @BirthDate= @BirthDate,
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
        PRINT 'Error occured: '+ERROR_MESSAGE();
    END CATCH;

END;
GO

-- Inserts Pet Owner
CREATE PROCEDURE InsertPetOwner
    @PetOwnerID UNIQUEIDENTIFIER,
    @VeterinarianID UNIQUEIDENTIFIER,
    @PetsSterilized BIT ,
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
                EXEC INSERTPerson
                     @PersonID = @PetOwnerID ,
                     @FirstName = @FirstName,
                     @LastName = @LastName,
                     @BirthDate= @BirthDate,
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

-- Inserts into the adopter table. If the AdopterID is not in the Person table it executes the insertPerson procedure and
-- inserts into person table
CREATE PROCEDURE INSERTADOPTER
    @AdopterID UNIQUEIDENTIFIER,
    @PetAllergies BIT,
    @HaveSurrendered BIT ,
    @HomeStatus VARCHAR (20),

    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE,
    @PhysicalAddress NVARCHAR(225),
    @MailingAddress NVARCHAR (225),
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
                EXEC INSERTPerson
                     @PersonID = @AdopterID ,
                     @FirstName = @FirstName,
                     @LastName = @LastName,
                     @BirthDate= @BirthDate,
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
        PRINT 'Error occurred: '+ ERROR_MESSAGE();
    END CATCH;

END;
GO

-- Inserts the Volunteer
CREATE PROCEDURE INSERTVOLUNTEER
    @VolunteerID UNIQUEIDENTIFIER,
    @VolunteerPositon NVARCHAR(50),
    @StartDate DATE,
    @EndDate DATE,
    @EmergencyContactName NVARCHAR(100),
    @EmergencyContactPhone NVARCHAR(20),
    @IsActive BIT,

    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE,
    @PhysicalAddress NVARCHAR(225),
    @MailingAddress NVARCHAR (225),
    @EmailAddress NVARCHAR(100),
    @PhoneNumber NVARCHAR(20)

AS
BEGIN
    SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;
        --Ensures the person doesnt already exist in the person table
        IF NOT EXISTS (SELECT 1 FROM people.Person WHERE PersonID = @VolunteerID)
            BEGIN
                EXEC INSERTPerson
                     @PersonID = @VolunteerID,
                     @FirstName = @FirstName,
                     @LastName = @LastName,
                     @BirthDate= @BirthDate,
                     @PhysicalAddress = @PhysicalAddress,
                     @MailingAddress = @MailingAddress,
                     @EmailAddress = @EmailAddress,
                     @PhoneNumber = @PhoneNumber;
            END
        -- Inserts in the volunteer table
        INSERT INTO people.Volunteer(VolunteerID, VolunteerPosition, StartDate, EndDate, EmergencyContactName, EmergencyContactPhone, IsActive)
        VALUES (@VolunteerID, @VolunteerPositon, @StartDate, @EndDate, @EmergencyContactName, @EmergencyContactPhone, @IsActive);

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        PRINT 'Error occured: '+ ERROR_MESSAGE();
    END CATCH;

END;
GO