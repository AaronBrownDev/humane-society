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