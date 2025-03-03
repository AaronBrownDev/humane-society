CREATE PROCEDURE InsertPetOwner
	@PetOwnerID UNIQUEIDENTIFIER,
   	@VetID UNIQUEIDENTIFIER,
    @PetsSterilized BIT ,
    @PetsVaccinated BIT,
    @HeartWormPreventionFromVet BIT,
	
	@FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE,
    @PhysicalAddress NVARCHAR(255),
   	@MailingAddress NVARCHAR(255),
    @Email NVARCHAR(100),
    @Phone NVARCHAR(20)


AS 
BEGIN 

	SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;
		 IF NOT EXISTS (SELECT 1 FROM people.Veterinarian WHERE VeterinarianID = @VetID)
        BEGIN
            INSERT INTO people.Veterinarian (VeterinarianID)
			VALUES (@VetID);
        END;

        -- Ensure the Person record is inserted first
        IF NOT EXISTS (SELECT 1 FROM people.Person WHERE PersonID = @PetOwnerID)
        BEGIN
            INSERT INTO people.Person (PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, Email, Phone)
            VALUES (@PetOwnerID, @FirstName, @LastName, @BirthDate, @PhysicalAddress, @MailingAddress, @Email, @Phone);
            
            PRINT 'Person inserted successfully';
        END

        -- Now insert into PetOwner (which references PersonID)
        INSERT INTO people.PetOwner (PetOwnerID, VetID, PetsSterilized, PetsVaccinated, HeartWormPreventionFromVet)
        VALUES (@PetOwnerID, @VetID, @PetsSterilized, @PetsVaccinated, @HeartWormPreventionFromVet);

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        PRINT 'Error occurred: ' + ERROR_MESSAGE();
    END CATCH;

END; 

DECLARE @NewPetOwnerID UNIQUEIDENTIFIER = NEWID();
DECLARE @NewVetID UNIQUEIDENTIFIER = NEWID();

INSERT INTO people.Veterinarian (VeterinarianID) 
VALUES (@NewVetID);


EXEC InsertPetOwner
    @PetOwnerID = @NewPetOwnerID,
    @VetID = @NewVetID, 
    @PetsSterilized = 1, 
    @PetsVaccinated = 1, 
    @HeartWormPreventionFromVet = 1, 

    @FirstName = 'John', 
    @LastName = 'Doe', 
    @BirthDate = '1990-05-15', 
    @PhysicalAddress = '123 Main St', 
    @MailingAddress = '123 Main St', 
    @Email = 'john.doe@email.com', 
    @Phone = '123-456-7890';

--SELECT * FROM people.Person WHERE FirstName = 'John';

--select * from people.Veterinarian;

--SELECT po.*
--FROM people.PetOwner po
--JOIN people.Person p ON po.PetOwnerID = p.PersonID
--WHERE p.FirstName = 'John';
