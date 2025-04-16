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