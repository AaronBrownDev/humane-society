CREATE TABLE shelter.Dog (
    DogID UNIQUEIDENTIFIER NOT NULL,
    Name NVARCHAR(50) NOT NULL,
    IntakeDate DATE NOT NULL,
    EstimatedBirthDate DATE NOT NULL,
    Breed NVARCHAR(50) NOT NULL,
    Sex VARCHAR(8) NOT NULL,
    Color NVARCHAR(30) NOT NULL,
    CageNumber INT NOT NULL,
    IsAdopted BIT NOT NULL DEFAULT 0,
    CONSTRAINT PK_Dog PRIMARY KEY (DogID),
    CONSTRAINT CK_Dog_Sex CHECK (Sex IN ('Male', 'Female', 'Intersex'))
);