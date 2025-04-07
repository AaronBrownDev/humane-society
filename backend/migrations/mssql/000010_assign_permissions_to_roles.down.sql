USE HumaneSociety;
GO

-- Revoke permissions from ReadOnly role
REVOKE SELECT ON SCHEMA::shelter FROM HumaneSociety_ReadOnly;
REVOKE SELECT ON SCHEMA::people FROM HumaneSociety_ReadOnly;
REVOKE SELECT ON SCHEMA::medical FROM HumaneSociety_ReadOnly;
GO

-- Revoke permissions from Volunteer role
REVOKE SELECT ON SCHEMA::shelter FROM HumaneSociety_Volunteer;
REVOKE SELECT ON SCHEMA::people FROM HumaneSociety_Volunteer;
REVOKE SELECT ON SCHEMA::medical FROM HumaneSociety_Volunteer;
REVOKE UPDATE ON shelter.Dog FROM HumaneSociety_Volunteer;
REVOKE INSERT, UPDATE ON shelter.AdoptionForm FROM HumaneSociety_Volunteer;
REVOKE INSERT, UPDATE ON shelter.SurrenderForm FROM HumaneSociety_Volunteer;
REVOKE INSERT, UPDATE, DELETE ON people.VolunteerSchedule FROM HumaneSociety_Volunteer;
GO

-- Revoke permissions from Staff role
REVOKE SELECT, INSERT, UPDATE, DELETE ON SCHEMA::shelter FROM HumaneSociety_Staff;
REVOKE SELECT, INSERT, UPDATE, DELETE ON SCHEMA::people FROM HumaneSociety_Staff;
REVOKE SELECT, INSERT, UPDATE, DELETE ON SCHEMA::medical FROM HumaneSociety_Staff;
REVOKE SELECT ON SCHEMA::audit FROM HumaneSociety_Staff;
GO

-- Revoke permissions from Admin role
REVOKE CONTROL ON SCHEMA::shelter FROM HumaneSociety_Admin;
REVOKE CONTROL ON SCHEMA::people FROM HumaneSociety_Admin;
REVOKE CONTROL ON SCHEMA::medical FROM HumaneSociety_Admin;
REVOKE CONTROL ON SCHEMA::audit FROM HumaneSociety_Admin;
GO