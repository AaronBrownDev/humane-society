USE HumaneSociety;
GO

-- Drop indexes first
DROP INDEX IF EXISTS IX_Supply_ItemID ON shelter.Supply;
DROP INDEX IF EXISTS IX_Dog_Adoption ON shelter.Dog;
DROP INDEX IF EXISTS IX_Person_Name ON people.Person;

-- Drop tables in reverse order of dependencies
DROP TABLE IF EXISTS shelter.Supply;
DROP TABLE IF EXISTS shelter.ItemCatalog;
DROP TABLE IF EXISTS medical.Medicine;
DROP TABLE IF EXISTS shelter.Dog;
DROP TABLE IF EXISTS people.Person;
GO