-- Delete data in reverse order of dependencies
DELETE FROM audit.ChangeLog;
DELETE FROM people.VolunteerSchedule;
DELETE FROM shelter.VolunteerForm;
DELETE FROM shelter.AdoptionForm;
DELETE FROM shelter.SurrenderForm;
DELETE FROM people.PetOwnerPets;
DELETE FROM medical.DogPrescription;
DELETE FROM shelter.Supply;
DELETE FROM shelter.ItemCatalog;
DELETE FROM medical.Medicine;
DELETE FROM shelter.Dog;
DELETE FROM people.Volunteer;
DELETE FROM people.PetOwner;
DELETE FROM people.PetSurrenderer;
DELETE FROM people.Veterinarian;
DELETE FROM people.Adopter;
DELETE FROM people.Person;
GO