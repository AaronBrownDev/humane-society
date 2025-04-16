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
