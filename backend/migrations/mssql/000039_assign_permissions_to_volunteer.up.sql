-- Volunteer role gets limited rights
GRANT SELECT ON SCHEMA::shelter TO HumaneSociety_Volunteer;
GRANT SELECT ON SCHEMA::people TO HumaneSociety_Volunteer;
GRANT SELECT ON SCHEMA::medical TO HumaneSociety_Volunteer;

-- Allow volunteers to update specific tables
GRANT UPDATE ON shelter.Dog TO HumaneSociety_Volunteer;
GRANT INSERT, UPDATE ON shelter.AdoptionForm TO HumaneSociety_Volunteer;
GRANT INSERT, UPDATE ON shelter.SurrenderForm TO HumaneSociety_Volunteer;
GRANT INSERT, UPDATE, DELETE ON people.VolunteerSchedule TO HumaneSociety_Volunteer;