USE HumaneSociety;
GO

-- Veterinarian table
CREATE TABLE people.Veterinarian (
    VeterinarianID UNIQUEIDENTIFIER NOT NULL,
    CONSTRAINT PK_Veterinarian PRIMARY KEY (VeterinarianID),
    CONSTRAINT FK_Veterinarian_Person FOREIGN KEY (VeterinarianID)
        REFERENCES people.Person(PersonID)
        ON DELETE CASCADE
);
GO

-- Adopter table
CREATE TABLE people.Adopter (
    AdopterID UNIQUEIDENTIFIER NOT NULL,
    HasPetAllergies BIT NOT NULL DEFAULT 0,
    HasSurrenderedPets BIT NOT NULL DEFAULT 0,
    HomeStatus VARCHAR(20) NOT NULL DEFAULT 'Pending',
    CONSTRAINT PK_Adopter PRIMARY KEY (AdopterID),
    CONSTRAINT FK_Adopter_Person FOREIGN KEY (AdopterID)
        REFERENCES people.Person(PersonID)
        ON DELETE CASCADE,
    CONSTRAINT CK_Adopter_HomeStatus CHECK (HomeStatus IN ('Pending', 'Approved', 'Rejected'))
);
GO

-- Pet Surrenderer table
CREATE TABLE people.PetSurrenderer (
    SurrendererID UNIQUEIDENTIFIER NOT NULL,
    CONSTRAINT PK_PetSurrenderer PRIMARY KEY (SurrendererID),
    CONSTRAINT FK_PetSurrenderer_Person FOREIGN KEY (SurrendererID)
        REFERENCES people.Person(PersonID)
        ON DELETE CASCADE
);
GO

-- Pet Owner table
CREATE TABLE people.PetOwner (
    PetOwnerID UNIQUEIDENTIFIER NOT NULL,
    VeterinarianID UNIQUEIDENTIFIER NULL,
    HasSterilizedPets BIT NOT NULL DEFAULT 0,
    HasVaccinatedPets BIT NOT NULL DEFAULT 0,
    UsesVetHeartWormPrevention BIT NOT NULL DEFAULT 0,
    CONSTRAINT PK_PetOwner PRIMARY KEY (PetOwnerID),
    CONSTRAINT FK_PetOwner_Person FOREIGN KEY (PetOwnerID)
        REFERENCES people.Person(PersonID)
        ON DELETE CASCADE,
    CONSTRAINT FK_PetOwner_Veterinarian FOREIGN KEY (VeterinarianID)
        REFERENCES people.Veterinarian(VeterinarianID)
        ON DELETE NO ACTION
);
GO

-- Volunteer table
CREATE TABLE people.Volunteer (
    VolunteerID UNIQUEIDENTIFIER NOT NULL,
    VolunteerPosition NVARCHAR(50) NOT NULL,
    StartDate DATE NOT NULL DEFAULT GETDATE(),
    EndDate DATE NULL,
    EmergencyContactName NVARCHAR(100) NULL,
    EmergencyContactPhone VARCHAR(20) NULL,
    IsActive BIT NOT NULL DEFAULT 1,
    CONSTRAINT PK_Volunteer PRIMARY KEY (VolunteerID),
    CONSTRAINT FK_Volunteer_Person FOREIGN KEY (VolunteerID)
        REFERENCES people.Person(PersonID)
        ON DELETE CASCADE,
    CONSTRAINT CK_Volunteer_EmergencyContact CHECK
        ((EmergencyContactName IS NULL AND EmergencyContactPhone IS NULL) OR
        (EmergencyContactName IS NOT NULL AND EmergencyContactPhone IS NOT NULL))
);
GO