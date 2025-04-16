REVOKE SELECT ON SCHEMA::shelter FROM HumaneSociety_Volunteer;
REVOKE SELECT ON SCHEMA::people FROM HumaneSociety_Volunteer;
REVOKE SELECT ON SCHEMA::medical FROM HumaneSociety_Volunteer;
REVOKE UPDATE ON shelter.Dog FROM HumaneSociety_Volunteer;
REVOKE INSERT, UPDATE ON shelter.AdoptionForm FROM HumaneSociety_Volunteer;
REVOKE INSERT, UPDATE ON shelter.SurrenderForm FROM HumaneSociety_Volunteer;
REVOKE INSERT, UPDATE, DELETE ON people.VolunteerSchedule FROM HumaneSociety_Volunteer;