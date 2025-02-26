CREATE DATABASE HumaneSociety

USE HumaneSociety;

CREATE TABLE Dog (
    DogID VARCHAR(32) NOT NULL,
    DogName VARCHAR(32) NOT NULL,
    IntakeDate DATETIME NOT NULL,
    EstimatedBirthDate DATETIME NOT NULL,
    Breed VARCHAR(32) NOT NULL,
    Sex VARCHAR(8) CHECK (Sex IN ('Male', 'Female', 'Intersex')),
    Color VARCHAR(32) NOT NULL,
    CageNumber INT NOT NULL,
    IsAdopted BIT DEFAULT 0,
    CONSTRAINT PK_Dog PRIMARY KEY (DogID),
)

CREATE TABLE Medicine (
    MedicineID INT NOT NULL,
    MedicationName VARCHAR(32) NOT NULL,
    Manufacturer VARCHAR(32) NOT NULL,
    CONSTRAINT PK_Medicine PRIMARY KEY (MedicineID),
)

CREATE TABLE DogPrescription (
    DogID VARCHAR(32) NOT NULL FOREIGN KEY REFERENCES Dog(DogID),
    MedicineID INT NOT NULL FOREIGN KEY REFERENCES Medicine(MedicineID),
    Dosage DECIMAL(5,2) NOT NULL,
    CONSTRAINT PK_DogPrescription PRIMARY KEY (DogID, MedicineID),
)

CREATE TABLE Person (
    PersonID VARCHAR(32) NOT NULL,
    FirstName VARCHAR(32) NOT NULL,
    LastName VARCHAR(32) NOT NULL,
    BirthDate DATE NOT NULL,
    PhysicalAddress VARCHAR(100) NOT NULL,
    MailingAddress VARCHAR(100) NOT NULL,
    CONSTRAINT PK_Person PRIMARY KEY (PersonID),
)

CREATE TABLE Adopter (
    AdopterID VARCHAR(32) NOT NULL FOREIGN KEY REFERENCES Person(PersonID),
    IsPetOwner BIT NOT NULL,
    PetAllergies BIT NOT NULL,
    HaveSurrended BIT NOT NULL,
    CONSTRAINT PK_Adopter PRIMARY KEY (AdopterID),
)

CREATE TABLE PetOwnerAdopter (
    PetOwnerID VARCHAR(32) NOT NULL FOREIGN KEY REFERENCES Adopter(AdopterID),
    PetsSterilized BIT NOT NULL,
    PetsVaccinated BIT NOT NULL,
    HeartWormPreventionFromVet BIT NOT NULL,
    CONSTRAINT PK_PetOwnerAdopter PRIMARY KEY (PetOwnerID),
)

CREATE TABLE AdopterPets (
    PetOwnerID VARCHAR(32) NOT NULL FOREIGN KEY REFERENCES PetOwnerAdopter(PetOwnerID),
    PetName VARCHAR(32) NOT NULL,
    PetBreed VARCHAR(32) NOT NULL,
    OwnershipDate DATE NOT NULL,
    LivingSpace VARCHAR(7) CHECK (LivingSpace IN ('Indoor', 'Outdoor', 'Both')),
    CONSTRAINT PK_AdopterPets PRIMARY KEY (PetOwnerID, PetName)
)

CREATE TABLE AdoptionForm (
    AdopterID VARCHAR(32) FOREIGN KEY REFERENCES Adopter(AdopterID),
    InterestedPetID VARCHAR(32) FOREIGN KEY REFERENCES Dog(DogID),
    FormDate DATETIME DEFAULT GETDATE(),
    Status VARCHAR(10) CHECK (Status IN ('Pending', 'Approved', 'Rejected')),
    CONSTRAINT PK_AdoptionForm PRIMARY KEY (AdopterID, InterestedPetID)
)

CREATE TABLE Employee (
    EmployeeID VARCHAR(32) NOT NULL FOREIGN KEY REFERENCES Person(PersonID),
    EmployeeRole VARCHAR(32) NOT NULL,
    CONSTRAINT PK_Employee PRIMARY KEY (EmployeeID),
)