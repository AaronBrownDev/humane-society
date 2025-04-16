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