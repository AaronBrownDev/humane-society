CREATE TABLE medical.Medicine (
    MedicineID INT IDENTITY(1,1) NOT NULL,
    Name NVARCHAR(50) NOT NULL,
    Manufacturer NVARCHAR(50) NOT NULL,
    Description NVARCHAR(200) NULL,
    DosageUnit NVARCHAR(20) NULL,
    CONSTRAINT PK_Medicine PRIMARY KEY (MedicineID)
);