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