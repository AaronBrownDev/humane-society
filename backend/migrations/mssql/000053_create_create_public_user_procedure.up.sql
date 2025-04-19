CREATE OR ALTER PROCEDURE auth.CreatePublicUser
    @UserID UNIQUEIDENTIFIER,
    @PasswordHash NVARCHAR(255)
AS
BEGIN
    SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;

        INSERT INTO auth.UserAccount (UserID, PasswordHash)
        VALUES (@UserID, @PasswordHash);

        -- Add default Public role (assuming ID 4)
        INSERT INTO auth.UserRole (UserID, RoleID)
        SELECT @UserID, r.RoleID
        FROM auth.Role r
        WHERE r.Name = 'Public';

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        THROW;
    END CATCH;
END;