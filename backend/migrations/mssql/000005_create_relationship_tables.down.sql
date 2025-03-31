USE HumaneSociety;
GO

-- Drop in the correct order to respect dependencies
DROP TABLE IF EXISTS people.VolunteerSchedule;
DROP TABLE IF EXISTS shelter.VolunteerForm;
DROP INDEX IF EXISTS UQ_AdoptionForm_AdopterDog ON shelter.AdoptionForm;
DROP TABLE IF EXISTS shelter.AdoptionForm;
DROP TABLE IF EXISTS shelter.SurrenderForm;
DROP TABLE IF EXISTS people.PetOwnerPets;
DROP INDEX IF EXISTS UQ_DogPrescription_DogMedicine ON medical.DogPrescription;
DROP TABLE IF EXISTS medical.DogPrescription;
GO