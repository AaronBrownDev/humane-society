CREATE TABLE people.PetSurrenderer (
    SurrendererID UNIQUEIDENTIFIER NOT NULL,
    CONSTRAINT PK_PetSurrenderer PRIMARY KEY (SurrendererID),
    CONSTRAINT FK_PetSurrenderer_Person FOREIGN KEY (SurrendererID)
        REFERENCES people.Person(PersonID)
        ON DELETE CASCADE
);