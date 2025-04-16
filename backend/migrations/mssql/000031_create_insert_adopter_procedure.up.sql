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