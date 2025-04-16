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