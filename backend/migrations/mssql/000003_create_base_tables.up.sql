USE HumaneSociety;
GO

-- Person Table
CREATE TABLE people.Person (
    PersonID UNIQUEIDENTIFIER NOT NULL,
    FirstName NVARCHAR(50) NOT NULL,
    LastName NVARCHAR(50) NOT NULL,
    BirthDate DATE NOT NULL,
    PhysicalAddress NVARCHAR(150) NOT NULL,
    MailingAddress NVARCHAR(150) NOT NULL,
    EmailAddress VARCHAR(100) NULL,
    PhoneNumber VARCHAR(20) NULL,
    CONSTRAINT PK_Person PRIMARY KEY (PersonID)
);
GO

-- Index on LastName, FirstName for name searches
CREATE INDEX IX_Person_Name ON people.Person(LastName, FirstName);
GO

-- Dog Table
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
GO

-- Index for looking up available dogs
CREATE INDEX IX_Dog_Adoption ON shelter.Dog(IsAdopted);
GO

-- Medicine table
CREATE TABLE medical.Medicine (
    MedicineID INT IDENTITY(1,1) NOT NULL,
    Name NVARCHAR(50) NOT NULL,
    Manufacturer NVARCHAR(50) NOT NULL,
    Description NVARCHAR(200) NULL,
    DosageUnit NVARCHAR(20) NULL,
    CONSTRAINT PK_Medicine PRIMARY KEY (MedicineID)
);
GO

-- Item Catalog
CREATE TABLE shelter.ItemCatalog (
    ItemID UNIQUEIDENTIFIER NOT NULL,
    Name NVARCHAR(50) NOT NULL,
    Category NVARCHAR(30) NOT NULL,
    Description NVARCHAR(200) NULL,
    MinimumQuantity INT NOT NULL DEFAULT 0,
    IsActive BIT NOT NULL DEFAULT 1,
    CONSTRAINT PK_ItemCatalog PRIMARY KEY (ItemID),
    CONSTRAINT UK_ItemCatalog_Name UNIQUE (Name)
);
GO

-- Supply inventory
CREATE TABLE shelter.Supply (
    SupplyID INT IDENTITY(1,1) NOT NULL,
    ItemID UNIQUEIDENTIFIER NOT NULL,
    Quantity INT NOT NULL,
    StorageLocation NVARCHAR(50) NULL,
    ExpirationDate DATE NULL,
    BatchNumber NVARCHAR(50) NULL,
    AcquisitionDate DATE NULL DEFAULT GETDATE(),
    CONSTRAINT PK_Supply PRIMARY KEY (SupplyID),
    CONSTRAINT FK_Supply_ItemCatalog FOREIGN KEY (ItemID)
        REFERENCES shelter.ItemCatalog(ItemID)
        ON DELETE CASCADE,
    CONSTRAINT CK_Supply_Quantity CHECK (Quantity >= 0)
);
GO

-- Index for finding items by catalog ID
CREATE INDEX IX_Supply_ItemID ON shelter.Supply(ItemID);
GO