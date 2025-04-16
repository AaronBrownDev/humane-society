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