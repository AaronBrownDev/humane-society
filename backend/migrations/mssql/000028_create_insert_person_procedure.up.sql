CREATE OR ALTER PROCEDURE InsertPerson
    @PersonID UNIQUEIDENTIFIER,
    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE,
    @PhysicalAddress NVARCHAR(225),
    @MailingAddress NVARCHAR(225),
    @EmailAddress NVARCHAR(100),
    @PhoneNumber NVARCHAR(20)
AS
BEGIN
    SET NOCOUNT ON;
    INSERT INTO people.Person (PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, EmailAddress, PhoneNumber)
    VALUES (@PersonID, @FirstName, @LastName, @BirthDate, @PhysicalAddress, @MailingAddress, @EmailAddress, @PhoneNumber);
END;