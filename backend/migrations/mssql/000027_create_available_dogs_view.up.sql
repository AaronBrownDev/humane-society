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
    d.CageNumber,
    d.IsAdopted
FROM
    shelter.Dog AS d
WHERE
    d.IsAdopted = 0;
