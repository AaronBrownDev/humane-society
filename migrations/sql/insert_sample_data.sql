USE HumaneSociety;
GO

-- Clear existing data (if any) while maintaining referential integrity
-- Delete data in reverse order of dependencies

-- First delete from relationship tables
DELETE FROM audit.ChangeLog;
DELETE FROM people.VolunteerSchedule;
DELETE FROM shelter.VolunteerForm;
DELETE FROM shelter.AdoptionForm;
DELETE FROM shelter.SurrenderForm;
DELETE FROM people.PetOwnerPets;
DELETE FROM medical.DogPrescription;
DELETE FROM shelter.Supply;

-- Then delete from entity tables
DELETE FROM shelter.ItemCatalog;
DELETE FROM medical.Medicine;
DELETE FROM shelter.Dog;
DELETE FROM people.Volunteer;
DELETE FROM people.PetOwner;
DELETE FROM people.PetSurrenderer;
DELETE FROM people.Veterinarian;
DELETE FROM people.Adopter;
DELETE FROM people.Person;

------------------------------------------
-- SAMPLE PEOPLE DATA
------------------------------------------

-- Insert Person records
INSERT INTO people.Person (PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, Email, Phone)
VALUES
    (NEWID(), 'John', 'Smith', '1985-04-12', '123 Main St, Springfield, IL 62701', '123 Main St, Springfield, IL 62701', 'john.smith@email.com', '217-555-1234'),
    (NEWID(), 'Jane', 'Doe', '1990-08-24', '456 Oak Ave, Springfield, IL 62702', '456 Oak Ave, Springfield, IL 62702', 'jane.doe@email.com', '217-555-5678'),
    (NEWID(), 'Robert', 'Johnson', '1978-02-15', '789 Pine Rd, Springfield, IL 62703', '789 Pine Rd, Springfield, IL 62703', 'robert.johnson@email.com', '217-555-9012'),
    (NEWID(), 'Emily', 'Wilson', '1992-11-30', '321 Elm St, Springfield, IL 62704', '321 Elm St, Springfield, IL 62704', 'emily.wilson@email.com', '217-555-3456'),
    (NEWID(), 'Michael', 'Brown', '1983-07-08', '654 Maple Dr, Springfield, IL 62705', '654 Maple Dr, Springfield, IL 62705', 'michael.brown@email.com', '217-555-7890'),
    (NEWID(), 'Sarah', 'Davis', '1995-01-17', '987 Cedar Ln, Springfield, IL 62706', '987 Cedar Ln, Springfield, IL 62706', 'sarah.davis@email.com', '217-555-2345'),
    (NEWID(), 'David', 'Miller', '1975-09-22', '135 Walnut Ct, Springfield, IL 62707', '135 Walnut Ct, Springfield, IL 62707', 'david.miller@email.com', '217-555-6789'),
    (NEWID(), 'Jennifer', 'Taylor', '1988-03-05', '246 Birch Blvd, Springfield, IL 62708', '246 Birch Blvd, Springfield, IL 62708', 'jennifer.taylor@email.com', '217-555-0123'),
    (NEWID(), 'James', 'Anderson', '1980-12-14', '357 Spruce Way, Springfield, IL 62709', '357 Spruce Way, Springfield, IL 62709', 'james.anderson@email.com', '217-555-4567'),
    (NEWID(), 'Jessica', 'Thomas', '1993-06-27', '468 Aspen Pl, Springfield, IL 62710', '468 Aspen Pl, Springfield, IL 62710', 'jessica.thomas@email.com', '217-555-8901'),
    (NEWID(), 'William', 'Jackson', '1982-10-09', '579 Fir Dr, Springfield, IL 62711', '579 Fir Dr, Springfield, IL 62711', 'william.jackson@email.com', '217-555-2345'),
    (NEWID(), 'Amanda', 'White', '1991-05-03', '680 Hemlock St, Springfield, IL 62712', '680 Hemlock St, Springfield, IL 62712', 'amanda.white@email.com', '217-555-6789'),
    (NEWID(), 'Christopher', 'Harris', '1977-08-19', '791 Locust Ave, Springfield, IL 62713', '791 Locust Ave, Springfield, IL 62713', 'christopher.harris@email.com', '217-555-0123'),
    (NEWID(), 'Nicole', 'Martin', '1994-02-11', '802 Juniper Rd, Springfield, IL 62714', '802 Juniper Rd, Springfield, IL 62714', 'nicole.martin@email.com', '217-555-4567'),
    (NEWID(), 'Daniel', 'Thompson', '1986-07-23', '913 Magnolia Dr, Springfield, IL 62715', '913 Magnolia Dr, Springfield, IL 62715', 'daniel.thompson@email.com', '217-555-8901');

-- Save PersonIDs for reference in other tables
DECLARE @PersonID1 UNIQUEIDENTIFIER, @PersonID2 UNIQUEIDENTIFIER, @PersonID3 UNIQUEIDENTIFIER, @PersonID4 UNIQUEIDENTIFIER, @PersonID5 UNIQUEIDENTIFIER;
DECLARE @PersonID6 UNIQUEIDENTIFIER, @PersonID7 UNIQUEIDENTIFIER, @PersonID8 UNIQUEIDENTIFIER, @PersonID9 UNIQUEIDENTIFIER, @PersonID10 UNIQUEIDENTIFIER;
DECLARE @PersonID11 UNIQUEIDENTIFIER, @PersonID12 UNIQUEIDENTIFIER, @PersonID13 UNIQUEIDENTIFIER, @PersonID14 UNIQUEIDENTIFIER, @PersonID15 UNIQUEIDENTIFIER;

SELECT TOP 1 @PersonID1 = PersonID FROM people.Person WHERE FirstName = 'John' AND LastName = 'Smith';
SELECT TOP 1 @PersonID2 = PersonID FROM people.Person WHERE FirstName = 'Jane' AND LastName = 'Doe';
SELECT TOP 1 @PersonID3 = PersonID FROM people.Person WHERE FirstName = 'Robert' AND LastName = 'Johnson';
SELECT TOP 1 @PersonID4 = PersonID FROM people.Person WHERE FirstName = 'Emily' AND LastName = 'Wilson';
SELECT TOP 1 @PersonID5 = PersonID FROM people.Person WHERE FirstName = 'Michael' AND LastName = 'Brown';
SELECT TOP 1 @PersonID6 = PersonID FROM people.Person WHERE FirstName = 'Sarah' AND LastName = 'Davis';
SELECT TOP 1 @PersonID7 = PersonID FROM people.Person WHERE FirstName = 'David' AND LastName = 'Miller';
SELECT TOP 1 @PersonID8 = PersonID FROM people.Person WHERE FirstName = 'Jennifer' AND LastName = 'Taylor';
SELECT TOP 1 @PersonID9 = PersonID FROM people.Person WHERE FirstName = 'James' AND LastName = 'Anderson';
SELECT TOP 1 @PersonID10 = PersonID FROM people.Person WHERE FirstName = 'Jessica' AND LastName = 'Thomas';
SELECT TOP 1 @PersonID11 = PersonID FROM people.Person WHERE FirstName = 'William' AND LastName = 'Jackson';
SELECT TOP 1 @PersonID12 = PersonID FROM people.Person WHERE FirstName = 'Amanda' AND LastName = 'White';
SELECT TOP 1 @PersonID13 = PersonID FROM people.Person WHERE FirstName = 'Christopher' AND LastName = 'Harris';
SELECT TOP 1 @PersonID14 = PersonID FROM people.Person WHERE FirstName = 'Nicole' AND LastName = 'Martin';
SELECT TOP 1 @PersonID15 = PersonID FROM people.Person WHERE FirstName = 'Daniel' AND LastName = 'Thompson';

-- Assign roles to people
-- Adopters
INSERT INTO people.Adopter (AdopterID, PetAllergies, HaveSurrendered, HomeStatus)
VALUES
    (@PersonID1, 0, 0, 'Approved'),
    (@PersonID4, 1, 0, 'Approved'),
    (@PersonID8, 0, 0, 'Pending'),
    (@PersonID10, 0, 1, 'Approved'),
    (@PersonID12, 0, 0, 'Rejected');

-- Veterinarians
INSERT INTO people.Veterinarian (VeterinarianID)
VALUES
    (@PersonID3),
    (@PersonID11);

-- Pet Surrenderers
INSERT INTO people.PetSurrenderer (SurrendererID)
VALUES
    (@PersonID5),
    (@PersonID9),
    (@PersonID10); -- Jessica is both a surrenderer and adopter

-- Pet Owners
INSERT INTO people.PetOwner (PetOwnerID, VetID, PetsSterilized, PetsVaccinated, HeartWormPreventionFromVet)
VALUES
    (@PersonID1, @PersonID3, 1, 1, 1), -- John is both an adopter and pet owner
    (@PersonID6, @PersonID11, 1, 1, 0),
    (@PersonID12, @PersonID3, 0, 1, 0); -- Amanda is both an adopter and pet owner

-- Volunteers
INSERT INTO people.Volunteer (VolunteerID, VolunteerRole, StartDate, EndDate, EmergencyContactName, EmergencyContactPhone, IsActive)
VALUES
    (@PersonID2, 'Dog Walker', '2023-01-15', NULL, 'Robert Johnson', '217-555-9012', 1),
    (@PersonID7, 'Cleaner', '2023-03-10', NULL, 'Mary Miller', '217-555-1234', 1),
    (@PersonID13, 'Admin', '2022-11-05', NULL, 'Lisa Harris', '217-555-5678', 1),
    (@PersonID14, 'Fundraiser', '2023-05-22', NULL, 'Mark Martin', '217-555-9876', 1),
    (@PersonID15, 'Dog Walker', '2022-08-18', '2023-06-30', 'Karen Thompson', '217-555-4321', 0);

-- Pet Owner Pets
INSERT INTO people.PetOwnerPets (PetOwnerID, PetName, PetType, PetBreed, Sex, OwnershipDate, LivingSpace)
VALUES
    (@PersonID1, 'Max', 'Dog', 'Labrador Retriever', 'Male', '2020-05-10', 'Indoor'),
    (@PersonID1, 'Bella', 'Dog', 'Beagle', 'Female', '2021-03-15', 'Both'),
    (@PersonID6, 'Charlie', 'Dog', 'Golden Retriever', 'Male', '2019-08-22', 'Both'),
    (@PersonID6, 'Lucy', 'Cat', 'Siamese', 'Female', '2020-11-30', 'Indoor'),
    (@PersonID12, 'Cooper', 'Dog', 'German Shepherd', 'Male', '2022-01-05', 'Both');

------------------------------------------
-- SAMPLE DOGS DATA
------------------------------------------

-- Create dogs
DECLARE @Dog1 UNIQUEIDENTIFIER = NEWID();
DECLARE @Dog2 UNIQUEIDENTIFIER = NEWID();
DECLARE @Dog3 UNIQUEIDENTIFIER = NEWID();
DECLARE @Dog4 UNIQUEIDENTIFIER = NEWID();
DECLARE @Dog5 UNIQUEIDENTIFIER = NEWID();
DECLARE @Dog6 UNIQUEIDENTIFIER = NEWID();
DECLARE @Dog7 UNIQUEIDENTIFIER = NEWID();
DECLARE @Dog8 UNIQUEIDENTIFIER = NEWID();

INSERT INTO shelter.Dog (DogID, DogName, IntakeDate, EstimatedBirthDate, Breed, Sex, Color, CageNumber, IsAdopted)
VALUES
    (@Dog1, 'Rocky', '2023-02-15', '2021-06-10', 'Pit Bull Mix', 'Male', 'Brown', 1, 0),
    (@Dog2, 'Luna', '2023-03-22', '2022-01-05', 'Labrador Retriever', 'Female', 'Black', 2, 0),
    (@Dog3, 'Rex', '2023-01-08', '2020-11-12', 'German Shepherd', 'Male', 'Black and Tan', 3, 0),
    (@Dog4, 'Daisy', '2023-04-05', '2022-03-18', 'Beagle', 'Female', 'Tricolor', 4, 0),
    (@Dog5, 'Buddy', '2023-05-11', '2019-08-30', 'Golden Retriever', 'Male', 'Golden', 5, 0),
    (@Dog6, 'Mia', '2023-02-28', '2020-05-15', 'Boxer', 'Female', 'Brindle', 6, 1),
    (@Dog7, 'Duke', '2023-03-17', '2022-02-20', 'Husky', 'Male', 'Gray and White', 7, 0),
    (@Dog8, 'Stella', '2023-05-03', '2021-10-08', 'Australian Shepherd', 'Female', 'Merle', 8, 0);

------------------------------------------
-- SAMPLE MEDICAL DATA
------------------------------------------

-- Medicine
INSERT INTO medical.Medicine (MedicationName, Manufacturer, Description, DosageUnit)
VALUES
    ('Heartgard Plus', 'Boehringer Ingelheim', 'Heartworm prevention', 'Tablet'),
    ('Frontline Plus', 'Merial', 'Flea and tick prevention', 'ml'),
    ('Rimadyl', 'Zoetis', 'Pain relief and anti-inflammatory', 'mg'),
    ('Clavamox', 'Zoetis', 'Antibiotic', 'mg'),
    ('Cerenia', 'Zoetis', 'Anti-nausea and vomiting', 'mg');

-- Store medicine IDs in variables
DECLARE @MedicineID1 INT, @MedicineID2 INT, @MedicineID3 INT, @MedicineID4 INT, @MedicineID5 INT;

SELECT @MedicineID1 = MedicineID FROM medical.Medicine WHERE MedicationName = 'Heartgard Plus';
SELECT @MedicineID2 = MedicineID FROM medical.Medicine WHERE MedicationName = 'Frontline Plus';
SELECT @MedicineID3 = MedicineID FROM medical.Medicine WHERE MedicationName = 'Rimadyl';
SELECT @MedicineID4 = MedicineID FROM medical.Medicine WHERE MedicationName = 'Clavamox';
SELECT @MedicineID5 = MedicineID FROM medical.Medicine WHERE MedicationName = 'Cerenia';

-- Dog Prescriptions
INSERT INTO medical.DogPrescription (DogID, MedicineID, Dosage, Frequency, StartDate, EndDate, Notes, PrescribedBy)
VALUES
    (@Dog1, @MedicineID1, 1, 'Once monthly', '2023-02-20', NULL, 'Monthly heartworm prevention', @PersonID3),
    (@Dog1, @MedicineID2, 0.5, 'Once monthly', '2023-02-20', NULL, 'Monthly flea and tick prevention', @PersonID3),
    (@Dog3, @MedicineID4, 250, 'Twice daily', '2023-01-10', '2023-01-24', 'Treating respiratory infection', @PersonID11),
    (@Dog5, @MedicineID3, 75, 'Twice daily', '2023-05-15', '2023-05-29', 'Joint pain management', @PersonID3),
    (@Dog7, @MedicineID5, 16, 'Once daily', '2023-03-20', '2023-03-25', 'Motion sickness during transport', @PersonID11);

------------------------------------------
-- SAMPLE INVENTORY DATA
------------------------------------------

-- Item Catalog
INSERT INTO shelter.ItemCatalog (ItemID, ItemName, Category, Description, MinimumQuantity, IsActive)
VALUES
    (NEWID(), 'Dog Food - Adult', 'Food', 'Premium adult dog food', 5, 1),
    (NEWID(), 'Dog Food - Puppy', 'Food', 'Premium puppy food', 3, 1),
    (NEWID(), 'Dog Treats', 'Food', 'Training treats', 10, 1),
    (NEWID(), 'Leash - Standard', 'Equipment', '6 foot standard leash', 8, 1),
    (NEWID(), 'Collar - Small', 'Equipment', 'Small dog collar', 5, 1),
    (NEWID(), 'Collar - Medium', 'Equipment', 'Medium dog collar', 5, 1),
    (NEWID(), 'Collar - Large', 'Equipment', 'Large dog collar', 5, 1),
    (NEWID(), 'Bed - Small', 'Bedding', 'Small dog bed', 3, 1),
    (NEWID(), 'Bed - Large', 'Bedding', 'Large dog bed', 3, 1),
    (NEWID(), 'Blanket', 'Bedding', 'Fleece dog blanket', 10, 1),
    (NEWID(), 'Cleaning Solution', 'Cleaning', 'Pet-safe cleaning solution', 2, 1),
    (NEWID(), 'Bleach', 'Cleaning', 'Disinfectant', 2, 1),
    (NEWID(), 'Paper Towels', 'Cleaning', 'Paper towel rolls', 15, 1),
    (NEWID(), 'Dog Shampoo', 'Grooming', 'Gentle dog shampoo', 5, 1),
    (NEWID(), 'Nail Clippers', 'Grooming', 'Dog nail clippers', 3, 1);

-- Save ItemIDs for supply inventory
DECLARE @ItemID1 UNIQUEIDENTIFIER, @ItemID2 UNIQUEIDENTIFIER, @ItemID3 UNIQUEIDENTIFIER;
DECLARE @ItemID4 UNIQUEIDENTIFIER, @ItemID5 UNIQUEIDENTIFIER, @ItemID6 UNIQUEIDENTIFIER;
DECLARE @ItemID7 UNIQUEIDENTIFIER, @ItemID11 UNIQUEIDENTIFIER, @ItemID13 UNIQUEIDENTIFIER;
DECLARE @ItemID14 UNIQUEIDENTIFIER;

SELECT TOP 1 @ItemID1 = ItemID FROM shelter.ItemCatalog WHERE ItemName = 'Dog Food - Adult';
SELECT TOP 1 @ItemID2 = ItemID FROM shelter.ItemCatalog WHERE ItemName = 'Dog Food - Puppy';
SELECT TOP 1 @ItemID3 = ItemID FROM shelter.ItemCatalog WHERE ItemName = 'Dog Treats';
SELECT TOP 1 @ItemID4 = ItemID FROM shelter.ItemCatalog WHERE ItemName = 'Leash - Standard';
SELECT TOP 1 @ItemID5 = ItemID FROM shelter.ItemCatalog WHERE ItemName = 'Collar - Small';
SELECT TOP 1 @ItemID6 = ItemID FROM shelter.ItemCatalog WHERE ItemName = 'Collar - Medium';
SELECT TOP 1 @ItemID7 = ItemID FROM shelter.ItemCatalog WHERE ItemName = 'Collar - Large';
SELECT TOP 1 @ItemID11 = ItemID FROM shelter.ItemCatalog WHERE ItemName = 'Cleaning Solution';
SELECT TOP 1 @ItemID13 = ItemID FROM shelter.ItemCatalog WHERE ItemName = 'Paper Towels';
SELECT TOP 1 @ItemID14 = ItemID FROM shelter.ItemCatalog WHERE ItemName = 'Dog Shampoo';

-- Supplies
INSERT INTO shelter.Supply (ItemID, Quantity, Location, ExpirationDate, BatchNumber, AcquisitionDate)
VALUES
    (@ItemID1, 15, 'Main Storage', '2024-05-20', 'BT12345', '2023-05-15'),
    (@ItemID1, 10, 'Main Storage', '2024-08-10', 'BT12346', '2023-08-05'),
    (@ItemID2, 8, 'Main Storage', '2024-06-15', 'BT23456', '2023-06-10'),
    (@ItemID3, 25, 'Front Desk', '2024-12-01', 'BT34567', '2023-07-20'),
    (@ItemID4, 12, 'Equipment Room', NULL, 'EQ12345', '2023-03-15'),
    (@ItemID5, 6, 'Equipment Room', NULL, 'EQ23456', '2023-04-02'),
    (@ItemID6, 9, 'Equipment Room', NULL, 'EQ34567', '2023-04-02'),
    (@ItemID7, 7, 'Equipment Room', NULL, 'EQ45678', '2023-04-02'),
    (@ItemID11, 4, 'Janitorial Closet', '2025-01-15', 'CL12345', '2023-01-10'),
    (@ItemID13, 20, 'Janitorial Closet', NULL, 'CL23456', '2023-05-01'),
    (@ItemID14, 10, 'Grooming Area', '2024-09-30', 'GR12345', '2023-06-25');

------------------------------------------
-- SAMPLE FORMS DATA
------------------------------------------

-- Surrender Form
INSERT INTO shelter.SurrenderForm (
    SurrendererID, FormDate, DogName, DogAge, WeightInPounds, Sex, Breed, Color,
    LivingSpace, OwnershipDate, VetID, LastVetVisit, GoodWithChildren, GoodWithDogs,
    GoodWithCats, GoodWithStrangers, HouseTrained, Sterilized, MicroChipNumber,
    MedicalProblems, BiteHistory, Reason, ProcessedByVolunteerID, ProcessingDate,
    ResultingDogID, Status
)
VALUES
    (@PersonID5, '2023-02-10', 'Rocky', 2, 45.5, 'Male', 'Pit Bull Mix', 'Brown', 
     'Both', '2021-07-15', @PersonID3, '2022-11-20', 1, 1, 0, 1, 1, 1, '985120035487621',
     'None', 'No bite history', 'Moving to apartment that doesn''t allow dogs', @PersonID13, '2023-02-15',
     @Dog1, 'Completed'),
    
    (@PersonID9, '2023-03-18', 'Luna', 1, 35.2, 'Female', 'Labrador Retriever', 'Black', 
     'Indoor', '2022-02-10', @PersonID11, '2022-12-15', 1, 1, 1, 1, 0, 0, NULL,
     'None', 'No bite history', 'Unable to care for puppy with new baby', @PersonID13, '2023-03-22',
     @Dog2, 'Completed'),
    
    (@PersonID10, '2023-05-01', 'Stella', 1, 32.4, 'Female', 'Australian Shepherd', 'Merle', 
     'Both', '2022-01-15', @PersonID3, '2023-03-10', 1, 0, 0, 0, 1, 1, '985120035782145',
     'Sensitive stomach, needs special diet', 'No bite history', 'Too high energy for our lifestyle', @PersonID13, '2023-05-03',
     @Dog8, 'Completed'),
    
    (@PersonID5, '2023-05-28', 'Baxter', 3, 28.6, 'Male', 'Cocker Spaniel', 'Blonde',
     'Indoor', '2020-04-30', @PersonID11, '2023-04-15', 1, 1, 1, 1, 1, 1, '985120035963214',
     'Slight hearing loss in left ear', 'No bite history', 'Cannot afford care anymore', NULL, NULL,
     NULL, 'Pending');

-- Adoption Form
INSERT INTO shelter.AdoptionForm (
    AdopterID, InterestedPetID, FormDate, ProcessedByVolunteerID, ProcessingDate, Status, RejectionReason
)
VALUES
    (@PersonID1, @Dog5, '2023-05-20', @PersonID13, '2023-05-25', 'Approved', NULL),
    (@PersonID4, @Dog7, '2023-05-22', @PersonID13, '2023-05-27', 'Approved', NULL),
    (@PersonID8, @Dog4, '2023-05-15', NULL, NULL, 'Pending', NULL),
    (@PersonID10, @Dog3, '2023-05-18', @PersonID13, '2023-05-23', 'Pending', NULL),
    (@PersonID12, @Dog2, '2023-04-10', @PersonID13, '2023-04-15', 'Rejected', 'Home environment not suitable for this dog');

-- Volunteer Form
INSERT INTO shelter.VolunteerForm (
    ApplicantID, FormDate, SterilizationAndPetEducationPromotion, ShiftAvailable,
    BreedingBelief, CleaningDutiesAgreement, DogCareAgreement, DogAllergies, AnyLimitations,
    ForCommunityServiceHours, NeededCommunityServiceHours, HowDidYouHearAboutUs, QuestionsAndComments,
    ProcessedByVolunteerID, ProcessingDate, Status, RejectionReason
)
VALUES
    (@PersonID2, '2022-12-20', 1, 'Weekday mornings and weekends', 0, 1, 1, 0, 0, 0, NULL,
     'Through a friend who volunteers', 'Excited to help with the dogs!', @PersonID13, '2023-01-05', 'Approved', NULL),
    
    (@PersonID7, '2023-02-15', 1, 'Weekends only', 0, 1, 1, 0, 0, 0, NULL,
     'Local community board', 'I have experience with large dogs', @PersonID13, '2023-03-01', 'Approved', NULL),
    
    (@PersonID14, '2023-05-10', 1, 'Weekday evenings', 0, 1, 1, 0, 0, 0, NULL,
     'Social media post', 'I have fundraising experience', @PersonID13, '2023-05-15', 'Approved', NULL),
    
    (@PersonID15, '2022-07-25', 1, 'Flexible schedule', 0, 1, 1, 0, 0, 0, NULL,
     'Website', 'Looking to help out during my summer break', @PersonID13, '2022-08-05', 'Approved', NULL);

-- Volunteer Schedule
INSERT INTO people.VolunteerSchedule (
    VolunteerID, ScheduleDate, StartTime, EndTime, TaskDescription, Status
)
VALUES
    (@PersonID2, '2023-06-03', '09:00:00', '12:00:00', 'Dog walking and socialization', 'Scheduled'),
    (@PersonID2, '2023-06-10', '09:00:00', '12:00:00', 'Dog walking and socialization', 'Scheduled'),
    (@PersonID2, '2023-05-27', '09:00:00', '12:00:00', 'Dog walking and socialization', 'Completed'),
    (@PersonID2, '2023-05-20', '09:00:00', '12:00:00', 'Dog walking and socialization', 'Completed'),
    
    (@PersonID7, '2023-06-04', '10:00:00', '14:00:00', 'Kennel cleaning and laundry', 'Scheduled'),
    (@PersonID7, '2023-06-11', '10:00:00', '14:00:00', 'Kennel cleaning and laundry', 'Scheduled'),
    (@PersonID7, '2023-05-28', '10:00:00', '14:00:00', 'Kennel cleaning and laundry', 'Completed'),
    (@PersonID7, '2023-05-21', '10:00:00', '14:00:00', 'Kennel cleaning and laundry', 'Completed'),
    
    (@PersonID13, '2023-06-05', '13:00:00', '17:00:00', 'Office admin and adoption processing', 'Scheduled'),
    (@PersonID13, '2023-06-07', '13:00:00', '17:00:00', 'Office admin and adoption processing', 'Scheduled'),
    (@PersonID13, '2023-06-09', '13:00:00', '17:00:00', 'Office admin and adoption processing', 'Scheduled'),
    (@PersonID13, '2023-05-29', '13:00:00', '17:00:00', 'Office admin and adoption processing', 'Completed'),
    (@PersonID13, '2023-05-31', '13:00:00', '17:00:00', 'Office admin and adoption processing', 'Completed'),
    (@PersonID13, '2023-05-26', '13:00:00', '17:00:00', 'Office admin and adoption processing', 'Completed'),
    
    (@PersonID14, '2023-06-06', '17:00:00', '20:00:00', 'Fundraising call planning', 'Scheduled'),
    (@PersonID14, '2023-06-08', '17:00:00', '20:00:00', 'Fundraising call execution', 'Scheduled'),
    (@PersonID14, '2023-05-30', '17:00:00', '20:00:00', 'Fundraising call planning', 'Completed'),
    (@PersonID14, '2023-05-25', '17:00:00', '20:00:00', 'Fundraising call execution', 'Completed'),
    
    (@PersonID15, '2023-05-15', '09:00:00', '12:00:00', 'Dog walking and socialization', 'NoShow'),
    (@PersonID15, '2023-05-08', '09:00:00', '12:00:00', 'Dog walking and socialization', 'Completed'),
    (@PersonID15, '2023-05-01', '09:00:00', '12:00:00', 'Dog walking and socialization', 'Completed');

-- Sample audit log entries
INSERT INTO audit.ChangeLog (
    TableName, PrimaryKeyColumn, PrimaryKeyValue, ColumnName, OldValue, NewValue, ChangeDate, ChangedBy, Operation
)
VALUES
    ('shelter.Dog', 'DogID', CONVERT(NVARCHAR(36), @Dog6), 'IsAdopted', '0', '1', '2023-06-01 10:15:23', 'system', 'U'),
    ('shelter.Dog', 'DogID', CONVERT(NVARCHAR(36), @Dog6), 'AdoptionDate', 'NULL', '2023-06-01', '2023-06-01 10:15:23', 'system', 'U'),
    ('shelter.AdoptionForm', 'AdoptionFormID', '1', 'Status', 'Approved', 'Completed', '2023-06-01 10:15:23', 'system', 'U'),
    ('medical.DogPrescription', 'PrescriptionID', '3', 'EndDate', '2023-01-24', '2023-01-31', '2023-01-24 14:22:10', 'system', 'U'),
    ('shelter.Supply', 'SupplyID', '1', 'Quantity', '20', '15', '2023-05-28 09:45:33', 'system', 'U');

-- Mark one of the adoption forms as complete and the dog as adopted
UPDATE shelter.Dog SET IsAdopted = 1 WHERE DogID = @Dog6;
UPDATE shelter.AdoptionForm SET Status = 'Completed' WHERE AdopterID = @PersonID1 AND InterestedPetID = @Dog5;
