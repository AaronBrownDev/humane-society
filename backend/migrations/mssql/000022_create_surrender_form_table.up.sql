CREATE TABLE shelter.SurrenderForm (
    SurrenderFormID INT IDENTITY(1,1) NOT NULL,
    SurrendererID UNIQUEIDENTIFIER NOT NULL,
    SubmissionDate DATETIME2(0) NOT NULL DEFAULT GETDATE(),
    DogName NVARCHAR(50) NOT NULL,
    DogAge INT NOT NULL,
    WeightInPounds DECIMAL(5,2) NOT NULL,
    Sex VARCHAR(8) NOT NULL,
    Breed NVARCHAR(50) NULL,
    Color NVARCHAR(30) NULL,
    LivingEnvironment VARCHAR(7) NOT NULL,
    OwnershipDate DATE NOT NULL,
    VeterinarianID UNIQUEIDENTIFIER NULL,
    LastVetVisitDate DATE NULL,
    IsGoodWithChildren BIT NOT NULL DEFAULT 0,
    IsGoodWithDogs BIT NOT NULL DEFAULT 0,
    IsGoodWithCats BIT NOT NULL DEFAULT 0,
    IsGoodWithStrangers BIT NOT NULL DEFAULT 0,
    IsHouseTrained BIT NOT NULL DEFAULT 0,
    IsSterilized BIT NOT NULL DEFAULT 0,
    MicroChipNumber VARCHAR(15) NULL,
    MedicalProblems NVARCHAR(500) NULL,
    BiteHistory NVARCHAR(500) NULL,
    SurrenderReason NVARCHAR(500) NOT NULL,
    ProcessedByVolunteerID UNIQUEIDENTIFIER NULL,
    ProcessingDate DATETIME2(0) NULL,
    ResultingDogID UNIQUEIDENTIFIER NULL,
    Status VARCHAR(20) DEFAULT 'Pending',
    CONSTRAINT PK_SurrenderForm PRIMARY KEY (SurrenderFormID),
    CONSTRAINT FK_SurrenderForm_PetSurrenderer FOREIGN KEY (SurrendererID)
        REFERENCES people.PetSurrenderer(SurrendererID)
        ON DELETE CASCADE,
    CONSTRAINT FK_SurrenderForm_Veterinarian FOREIGN KEY (VeterinarianID)
        REFERENCES people.Veterinarian(VeterinarianID)
        ON DELETE NO ACTION,
    CONSTRAINT FK_SurrenderForm_Volunteer FOREIGN KEY (ProcessedByVolunteerID)
        REFERENCES people.Volunteer(VolunteerID)
        ON DELETE NO ACTION,
    CONSTRAINT FK_SurrenderForm_Dog FOREIGN KEY (ResultingDogID)
        REFERENCES shelter.Dog(DogID)
        ON DELETE NO ACTION,
    CONSTRAINT CK_SurrenderForm_Sex CHECK (Sex IN ('Male', 'Female', 'Intersex')),
    CONSTRAINT CK_SurrenderForm_LivingEnvironment CHECK (LivingEnvironment IN ('Indoor', 'Outdoor', 'Both')),
    CONSTRAINT CK_SurrenderForm_Status CHECK (Status IN ('Pending', 'Approved', 'Rejected', 'Completed'))
);