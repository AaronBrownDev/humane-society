CREATE UNIQUE INDEX UQ_DogPrescription_DogMedicine
    ON medical.DogPrescription(DogID, MedicineID, StartDate);