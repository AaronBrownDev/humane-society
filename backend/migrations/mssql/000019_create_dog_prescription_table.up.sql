CREATE TABLE medical.DogPrescription (
    PrescriptionID INT IDENTITY(1, 1) NOT NULL,
    DogID UNIQUEIDENTIFIER NOT NULL,
    MedicineID INT NOT NULL,
    Dosage DECIMAL(5,2) NOT NULL,
    Frequency NVARCHAR(50) NULL,
    StartDate DATE NOT NULL DEFAULT GETDATE(),
    EndDate DATE NULL,
    Notes NVARCHAR(200) NULL,
    VetPrescriberID UNIQUEIDENTIFIER NULL,
    CONSTRAINT PK_DogPrescription PRIMARY KEY (PrescriptionID),
    CONSTRAINT FK_DogPrescription_Dog FOREIGN KEY (DogID)
        REFERENCES shelter.Dog(DogID)
        ON DELETE CASCADE,
    CONSTRAINT FK_DogPrescription_Medicine FOREIGN KEY (MedicineID)
        REFERENCES medical.Medicine(MedicineID)
        ON DELETE CASCADE,
    CONSTRAINT FK_DogPrescription_Veterinarian FOREIGN KEY (VetPrescriberID)
        REFERENCES people.Veterinarian(VeterinarianID)
        ON DELETE NO ACTION
);