# Humane Society Database Roles and Security

This document outlines the security architecture for the Humane Society of Northwest Louisiana Management System database.

## Database Roles

The system uses SQL Server database roles to enforce least-privilege access and segregation of duties. These roles correspond to organizational responsibilities.

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

### Role Mappings

| Database Role           | Description                             | Organizational Position           |
|-------------------------|-----------------------------------------|-----------------------------------|
| HumaneSociety_Admin     | Full control over all database objects  | System Administrators, IT Staff   |
| HumaneSociety_Staff     | Data modification rights on most tables | Shelter Managers, Full-time Staff |
| HumaneSociety_Volunteer | Limited modification capabilities       | Shelter Volunteers                |
| HumaneSociety_ReadOnly  | Read-only access to operational data    | Public Website                    |
