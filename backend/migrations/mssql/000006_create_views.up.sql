USE HumaneSociety;
GO

-- Create a view for available dogs
CREATE VIEW shelter.AvailableDogs AS
SELECT
    d.DogID,
    d.Name,
    d.IntakeDate,
    d.EstimatedBirthDate,
    DATEDIFF(YEAR, d.EstimatedBirthDate, GETDATE()) AS AgeInYears,
    d.Breed,
    d.Sex,
    d.Color,
    d.CageNumber
FROM
    shelter.Dog AS d
WHERE
    d.IsAdopted = 0;
GO