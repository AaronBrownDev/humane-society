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