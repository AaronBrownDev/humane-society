CREATE TABLE people.PetOwnerPets (
    PetID INT IDENTITY(1,1) NOT NULL,
    PetOwnerID UNIQUEIDENTIFIER NOT NULL,
    Name NVARCHAR(50) NOT NULL,
    Type NVARCHAR(20) DEFAULT 'Dog',
    Breed NVARCHAR(50) NOT NULL,
    Sex VARCHAR(8) NOT NULL,
    OwnershipDate DATE NOT NULL,
    LivingEnvironment VARCHAR(7) NOT NULL,
    CONSTRAINT PK_PetOwnerPets PRIMARY KEY (PetID),
    CONSTRAINT FK_PetOwnerPets_PetOwner FOREIGN KEY (PetOwnerID)
        REFERENCES people.PetOwner(PetOwnerID)
        ON DELETE CASCADE,
    CONSTRAINT CK_PetOwnerPets_LivingEnvironment CHECK (LivingEnvironment IN ('Indoor', 'Outdoor', 'Both')),
    CONSTRAINT CK_PetOwnerPets_Sex CHECK (Sex IN ('Male', 'Female', 'Intersex'))
);