CREATE OR ALTER PROCEDURE auth.CreatePublicUser
    @UserID UNIQUEIDENTIFIER,
    @PasswordHash NVARCHAR(255)
AS
BEGIN
    SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;

        -- Insert into user account table
        INSERT INTO auth.UserAccount (UserID, PasswordHash)
        VALUES (@UserID, @PasswordHash);

        -- Get the Public role ID
        DECLARE @PublicRoleID INT;
        SELECT @PublicRoleID = RoleID
        FROM auth.Role
        WHERE Name = 'Public';

        -- Add default Public role
        IF @PublicRoleID IS NOT NULL
            BEGIN
                INSERT INTO auth.UserRole (UserID, RoleID, AssignedAt)
                VALUES (@UserID, @PublicRoleID, GETDATE());
            END

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        THROW;
    END CATCH;
END;