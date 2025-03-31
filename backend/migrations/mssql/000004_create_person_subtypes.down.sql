USE HumaneSociety;
GO

-- Drop tables in the correct order to respect dependencies
DROP TABLE IF EXISTS people.Volunteer;
DROP TABLE IF EXISTS people.PetOwner;
DROP TABLE IF EXISTS people.PetSurrenderer;
DROP TABLE IF EXISTS people.Adopter;
DROP TABLE IF EXISTS people.Veterinarian;
GO