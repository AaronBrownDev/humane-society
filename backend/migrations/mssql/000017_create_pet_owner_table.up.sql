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