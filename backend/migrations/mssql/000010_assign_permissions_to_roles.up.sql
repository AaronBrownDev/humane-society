USE HumaneSociety;
GO

-- Admin role gets full control
GRANT CONTROL ON SCHEMA::shelter TO HumaneSociety_Admin;
GRANT CONTROL ON SCHEMA::people TO HumaneSociety_Admin;
GRANT CONTROL ON SCHEMA::medical TO HumaneSociety_Admin;
GRANT CONTROL ON SCHEMA::audit TO HumaneSociety_Admin;
GO

-- Staff role gets data modification rights on most tables
GRANT SELECT, INSERT, UPDATE, DELETE ON SCHEMA::shelter TO HumaneSociety_Staff;
GRANT SELECT, INSERT, UPDATE, DELETE ON SCHEMA::people TO HumaneSociety_Staff;
GRANT SELECT, INSERT, UPDATE, DELETE ON SCHEMA::medical TO HumaneSociety_Staff;
GRANT SELECT ON SCHEMA::audit TO HumaneSociety_Staff;
GO

-- Volunteer role gets limited rights
GRANT SELECT ON SCHEMA::shelter TO HumaneSociety_Volunteer;
GRANT SELECT ON SCHEMA::people TO HumaneSociety_Volunteer;
GRANT SELECT ON SCHEMA::medical TO HumaneSociety_Volunteer;

-- Allow volunteers to update specific tables
GRANT UPDATE ON shelter.Dog TO HumaneSociety_Volunteer;
GRANT INSERT, UPDATE ON shelter.AdoptionForm TO HumaneSociety_Volunteer;
GRANT INSERT, UPDATE ON shelter.SurrenderForm TO HumaneSociety_Volunteer;
GRANT INSERT, UPDATE, DELETE ON people.VolunteerSchedule TO HumaneSociety_Volunteer;
GO

-- ReadOnly role gets read-only access
GRANT SELECT ON SCHEMA::shelter TO HumaneSociety_ReadOnly;
GRANT SELECT ON SCHEMA::people TO HumaneSociety_ReadOnly;
GRANT SELECT ON SCHEMA::medical TO HumaneSociety_ReadOnly;
GO