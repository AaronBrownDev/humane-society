CREATE TABLE people.Veterinarian (
    VeterinarianID UNIQUEIDENTIFIER NOT NULL,
    CONSTRAINT PK_Veterinarian PRIMARY KEY (VeterinarianID),
    CONSTRAINT FK_Veterinarian_Person FOREIGN KEY (VeterinarianID)
        REFERENCES people.Person(PersonID)
        ON DELETE CASCADE
);