CREATE PROCEDURE InsertPetOwner
    @PetOwnerID UNIQUEIDENTIFIER,
    @VeterinarianID UNIQUEIDENTIFIER,
    @HasSterilizedPets BIT,
    @HasVaccinatedPets BIT,
    @UsesVetHeartWormPrevention BIT,

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
        IF NOT EXISTS (SELECT 1 FROM people.Veterinarian WHERE VeterinarianID = @VeterinarianID)
            BEGIN
                INSERT INTO people.Veterinarian (VeterinarianID)
                VALUES (@VeterinarianID);
            END;

        -- Ensure the Person record is inserted first
        IF NOT EXISTS (SELECT 1 FROM people.Person WHERE PersonID = @PetOwnerID)
            BEGIN
                INSERT INTO people.Person (PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, EmailAddress, PhoneNumber)
                VALUES (@PetOwnerID, @FirstName, @LastName, @BirthDate, @PhysicalAddress, @MailingAddress, @EmailAddress, @PhoneNumber);

                PRINT 'Person inserted successfully';
            END

        -- Now insert into PetOwner (which references PersonID)
        INSERT INTO people.PetOwner (PetOwnerID, VeterinarianID, HasSterilizedPets, HasVaccinatedPets, UsesVetHeartWormPrevention)
        VALUES (@PetOwnerID, @VeterinarianID, @HasSterilizedPets, @HasVaccinatedPets, @UsesVetHeartWormPrevention);

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        PRINT 'Error occurred: ' + ERROR_MESSAGE();
    END CATCH;

END;
