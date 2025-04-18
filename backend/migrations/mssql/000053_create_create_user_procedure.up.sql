CREATE OR ALTER PROCEDURE auth.CreateUser
    @UserID UNIQUEIDENTIFIER,
    @Email VARCHAR(100),
    @PasswordHash NVARCHAR(255),
    @PersonID UNIQUEIDENTIFIER = NULL,
    @Roles NVARCHAR(MAX) = NULL
AS
BEGIN
    SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;

        -- Insert user
        INSERT INTO auth.UserAccount (UserID, Email, PasswordHash, PersonID)
        VALUES (@UserID, @Email, @PasswordHash, @PersonID);

        -- Add roles if provided
        IF @Roles IS NOT NULL
            BEGIN
                INSERT INTO auth.UserRole (UserID, RoleID)
                SELECT @UserID, r.RoleID
                FROM auth.Role r
                WHERE r.Name IN (
                    SELECT value FROM STRING_SPLIT(@Roles, ',')
                );
            END
        ELSE
            BEGIN
                -- Add default Public role
                INSERT INTO auth.UserRole (UserID, RoleID)
                SELECT @UserID, r.RoleID
                FROM auth.Role r
                WHERE r.Name = 'Public';
            END

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        THROW;
    END CATCH;
END;