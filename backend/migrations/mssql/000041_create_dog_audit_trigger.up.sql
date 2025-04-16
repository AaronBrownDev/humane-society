-- Dog table audit trigger
CREATE TRIGGER shelter_Dog_Audit ON shelter.Dog
    AFTER INSERT, UPDATE, DELETE
    AS
BEGIN
    SET NOCOUNT ON;

    DECLARE @Action CHAR(1);

    -- Figure out what operation is happening (Insert, Update, Delete)
    IF EXISTS (SELECT * FROM inserted) AND EXISTS (SELECT * FROM deleted)
        SET @Action = 'U'; -- Update
    ELSE IF EXISTS (SELECT * FROM inserted)
        SET @Action = 'I'; -- Insert
    ELSE
        SET @Action = 'D'; -- Delete

    -- For INSERT operations
    IF @Action = 'I'
        BEGIN
            -- Insert records into the audit log for each inserted row
            INSERT INTO audit.ChangeLog
            (TableName, PrimaryKeyColumn, PrimaryKeyValue,
             ColumnName, OldValue, NewValue, AuditActionType)
            SELECT
                'shelter.Dog', -- Table Name
                'DogID',      -- Primary Key Column
                CONVERT(NVARCHAR(36), i.DogID), -- Primary Key Value
                'All Columns', -- Column Name (simplified to just track "all columns")
                NULL,         -- Old Value (none for INSERT)
                'New Dog: ' + i.Name, -- Simplified new value
                'I'           -- Action Type
            FROM
                inserted i;
        END

    -- For UPDATE operations
    IF @Action = 'U'
        BEGIN
            -- Log name changes
            IF UPDATE(Name)
                BEGIN
                    INSERT INTO audit.ChangeLog
                    (TableName, PrimaryKeyColumn, PrimaryKeyValue,
                     ColumnName, OldValue, NewValue, AuditActionType)
                    SELECT
                        'shelter.Dog',
                        'DogID',
                        CONVERT(NVARCHAR(36), i.DogID),
                        'Name',
                        d.Name,
                        i.Name,
                        'U'
                    FROM
                        inserted i
                            JOIN deleted d ON i.DogID = d.DogID
                    WHERE
                        i.Name <> d.Name;
                END

            -- Log adoption status changes
            IF UPDATE(IsAdopted)
                BEGIN
                    INSERT INTO audit.ChangeLog
                    (TableName, PrimaryKeyColumn, PrimaryKeyValue,
                     ColumnName, OldValue, NewValue, AuditActionType)
                    SELECT
                        'shelter.Dog',
                        'DogID',
                        CONVERT(NVARCHAR(36), i.DogID),
                        'IsAdopted',
                        CASE WHEN d.IsAdopted = 1 THEN 'Yes' ELSE 'No' END,
                        CASE WHEN i.IsAdopted = 1 THEN 'Yes' ELSE 'No' END,
                        'U'
                    FROM
                        inserted i
                            JOIN deleted d ON i.DogID = d.DogID
                    WHERE
                        i.IsAdopted <> d.IsAdopted;
                END
        END

    -- For DELETE operations
    IF @Action = 'D'
        BEGIN
            INSERT INTO audit.ChangeLog
            (TableName, PrimaryKeyColumn, PrimaryKeyValue,
             ColumnName, OldValue, NewValue, AuditActionType)
            SELECT
                'shelter.Dog',
                'DogID',
                CONVERT(NVARCHAR(36), d.DogID),
                'All Columns',
                'Dog Deleted: ' + d.Name,
                NULL,
                'D'
            FROM
                deleted d;
        END
END;