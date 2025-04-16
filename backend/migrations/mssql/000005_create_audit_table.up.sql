CREATE TABLE audit.ChangeLog (
    LogID INT IDENTITY(1,1) NOT NULL,
    TableName NVARCHAR(128) NOT NULL,
    PrimaryKeyColumn NVARCHAR(128) NOT NULL,
    PrimaryKeyValue NVARCHAR(36) NOT NULL,
    ColumnName NVARCHAR(128) NOT NULL,
    OldValue NVARCHAR(MAX) NULL,
    NewValue NVARCHAR(MAX) NULL,
    ChangeDate DATETIME2(0) NOT NULL DEFAULT GETDATE(),
    ChangedBy NVARCHAR(128) NOT NULL DEFAULT SYSTEM_USER,
    AuditActionType CHAR(1) NOT NULL, -- I = Insert, U = Update, D = Delete
    CONSTRAINT PK_ChangeLog PRIMARY KEY (LogID),
    CONSTRAINT CK_ChangeLog_AuditActionType CHECK (AuditActionType IN ('I', 'U', 'D'))
);