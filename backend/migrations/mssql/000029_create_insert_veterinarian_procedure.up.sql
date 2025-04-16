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