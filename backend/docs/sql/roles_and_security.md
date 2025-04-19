# Humane Society Database Roles and Security

This document outlines the security architecture for the Humane Society of Northwest Louisiana Management System database.

## Database Security Model

The system uses a dual-layer security approach:

1. **Database-Level Roles**: SQL Server database roles with specific permissions to enforce least-privilege access.
2. **Application-Level Roles**: User roles stored in the database that control access within the application.

## Database Roles

Database roles correspond to organizational responsibilities and are used by the application to connect to the database with appropriate permissions.

### Role Definitions

```sql
-- Administrator role with full system control
CREATE ROLE HumaneSociety_Admin;

-- Staff role for everyday shelter operations
CREATE ROLE HumaneSociety_Staff;

-- Volunteer role with limited data modification capabilities
CREATE ROLE HumaneSociety_Volunteer;

-- Read-only role for reporting and limited access
CREATE ROLE HumaneSociety_ReadOnly;
```

### Role Permissions

#### Administrator Role
```sql
-- Grant control over all schemas
GRANT CONTROL ON SCHEMA::shelter TO HumaneSociety_Admin;
GRANT CONTROL ON SCHEMA::people TO HumaneSociety_Admin;
GRANT CONTROL ON SCHEMA::medical TO HumaneSociety_Admin;
GRANT CONTROL ON SCHEMA::audit TO HumaneSociety_Admin;
```

#### Staff Role
```sql
-- Grant read/write permissions to most tables
GRANT SELECT, INSERT, UPDATE, DELETE ON SCHEMA::shelter TO HumaneSociety_Staff;
GRANT SELECT, INSERT, UPDATE, DELETE ON SCHEMA::people TO HumaneSociety_Staff;
GRANT SELECT, INSERT, UPDATE, DELETE ON SCHEMA::medical TO HumaneSociety_Staff;
GRANT SELECT ON SCHEMA::audit TO HumaneSociety_Staff;
```

#### Volunteer Role
```sql
-- Grant limited read/write permissions
GRANT SELECT ON SCHEMA::shelter TO HumaneSociety_Volunteer;
GRANT SELECT ON SCHEMA::people TO HumaneSociety_Volunteer;
GRANT SELECT ON SCHEMA::medical TO HumaneSociety_Volunteer;

-- Allow volunteers to update specific tables
GRANT UPDATE ON shelter.Dog TO HumaneSociety_Volunteer;
GRANT INSERT, UPDATE ON shelter.AdoptionForm TO HumaneSociety_Volunteer;
GRANT INSERT, UPDATE ON shelter.SurrenderForm TO HumaneSociety_Volunteer;
GRANT INSERT, UPDATE, DELETE ON people.VolunteerSchedule TO HumaneSociety_Volunteer;
```

#### Read-Only Role
```sql
-- Grant read-only access to operational data
GRANT SELECT ON SCHEMA::shelter TO HumaneSociety_ReadOnly;
GRANT SELECT ON SCHEMA::people TO HumaneSociety_ReadOnly;
GRANT SELECT ON SCHEMA::medical TO HumaneSociety_ReadOnly;
```

### Role Mappings

| Database Role | Description | Organizational Position |
|---------------|-------------|-------------------------|
| HumaneSociety_Admin | Full control over all database objects | System Administrators, IT Staff |
| HumaneSociety_Staff | Data modification rights on most tables | Shelter Managers, Full-time Staff |
| HumaneSociety_Volunteer | Limited modification capabilities | Shelter Volunteers |
| HumaneSociety_ReadOnly | Read-only access to operational data | Public Website, Basic Users |

## Application Roles

The application uses a role-based authentication system stored in the `auth` schema.

### Application Role Definitions

```sql
-- Default application roles
INSERT INTO auth.Role (Name, Description)
VALUES
('Admin', 'System administrators with full access'),
('Staff', 'Staff members with access to most functions'),
('Volunteer', 'Volunteers with limited access'),
('Public', 'Basic public user role');
```

### Role Hierarchy

1. **Admin**: Full access to all system functions
2. **Staff**: Access to most operational functions
3. **Volunteer**: Limited access to volunteer-specific functions
4. **Public**: Basic access for adoption applications and information viewing

### User-Role Assignment

Users are assigned roles via the `auth.UserRole` table. By default, new users are assigned the 'Public' role through the `auth.CreatePublicUser` stored procedure.

## Refresh Token Security

The system uses refresh tokens for maintaining authenticated sessions:

1. **Token Storage**: Refresh tokens are stored in the `auth.RefreshToken` table
2. **Token Rotation**: When used, tokens are rotated for security (old token references new token)
3. **Expiration**: Tokens have defined expiration times
4. **Revocation**: Tokens can be manually revoked

## Audit System

All important data changes are tracked in the audit system:

```sql
-- Dog table audit trigger example
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

    -- Insert appropriate audit records based on action type
    -- ...
END;
```

## Security Best Practices

1. **Connection Segregation**: Different application components connect with different database roles
2. **Password Hashing**: User passwords are hashed before storage
3. **Error Handling**: Security-related errors are logged but not exposed to users
4. **Transaction Integrity**: All operations that modify multiple tables use transactions
5. **Connection Pooling**: Configured with appropriate timeouts and limits for security