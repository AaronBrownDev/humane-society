CREATE UNIQUE INDEX UQ_AdoptionForm_AdopterDog
    ON shelter.AdoptionForm(AdopterID, DogID)
    WHERE Status IN ('Pending', 'HomeVisitScheduled', 'Approved');