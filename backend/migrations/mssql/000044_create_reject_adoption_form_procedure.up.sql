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