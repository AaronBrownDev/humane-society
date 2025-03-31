# Stored Procedures

This document contains the SQL stored procedures used in the Humane Society of Northwest Louisiana Management System. These procedures provide encapsulated business logic for common database operations.

## Person Management

### INSERT Person

The base procedure for creating a new person record.

```sql
-- Inserts the person into the person table 
CREATE PROCEDURE INSERTPerson 
    @PersonID UNIQUEIDENTIFIER,
    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE, 
    @PhysicalAddress NVARCHAR(225),
    @MailingAddress NVARCHAR (225),
    @EmailAddress NVARCHAR(100),
    @PhoneNumber NVARCHAR(20) 

AS 
BEGIN 
    SET NOCOUNT ON;
    INSERT INTO people.Person (PersonID, FirstName, LastName, BirthDate, PhysicalAddress, MailingAddress, EmailAddress, PhoneNumber)
    VALUES (@PersonID, @FirstName, @LastName, @BirthDate, @PhysicalAddress, @MailingAddress, @EmailAddress, @PhoneNumber); 
                
END;
GO
```

### INSERT Veterinarian

Creates a new veterinarian record, including the base person record if needed.

```sql
CREATE PROCEDURE INSERTVeterinarian
    @VeterinarianID UNIQUEIDENTIFIER, 

    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE, 
    @PhysicalAddress NVARCHAR(225),
    @MailingAddress NVARCHAR (225),
    @EmailAddress NVARCHAR(100),
    @PhoneNumber NVARCHAR(20) 
    
AS 
BEGIN 
    SET NOCOUNT ON; 
        
    BEGIN TRY 
        BEGIN TRANSACTION;
        -- Create person record if not exists
        IF NOT EXISTS (SELECT 1 FROM people.Person WHERE PersonID = @VeterinarianID)
        BEGIN
            EXEC INSERTPerson
            @PersonID = @VeterinarianID, 
            @FirstName = @FirstName,
            @LastName = @LastName, 
            @BirthDate= @BirthDate,
            @PhysicalAddress = @PhysicalAddress,
            @MailingAddress = @MailingAddress, 
            @EmailAddress = @EmailAddress, 
            @PhoneNumber = @PhoneNumber;

            PRINT 'Person inserted successfully';
        END
        
        -- Create veterinarian record
        INSERT INTO people.Veterinarian(VeterinarianID)
        VALUES (@VeterinarianID)
        
        COMMIT TRANSACTION; 
    END TRY
    BEGIN CATCH 
        ROLLBACK TRANSACTION
        PRINT 'Error occured: '+ERROR_MESSAGE();
    END CATCH; 
    
END;
GO
```

### Insert Pet Owner

Creates a new pet owner record, including the base person record if needed.

```sql
-- Inserts Pet Owner
CREATE PROCEDURE InsertPetOwner
    @PetOwnerID UNIQUEIDENTIFIER,
    @VeterinarianID UNIQUEIDENTIFIER,
    @PetsSterilized BIT,
    @PetsVaccinated BIT,
    @HeartWormPreventionFromVet BIT,
    
    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE,
    @PhysicalAddress NVARCHAR(255),
    @MailingAddress NVARCHAR(255),
    @EmailAddress NVARCHAR(100),
    @PhoneNumber NVARCHAR(20)

AS 
BEGIN 
    SET NOCOUNT ON;

    BEGIN TRY
        BEGIN TRANSACTION;
        -- Ensures the vetID already exists
        IF NOT EXISTS (SELECT 1 FROM people.Veterinarian WHERE VeterinarianID = @VeterinarianID)
        BEGIN
            INSERT INTO people.Veterinarian (VeterinarianID)
            VALUES (@VeterinarianID);
        END;

        -- Ensure the Person record is inserted first
        IF NOT EXISTS (SELECT 1 FROM people.Person WHERE PersonID = @PetOwnerID)
        BEGIN
            EXEC INSERTPerson 
                @PersonID = @PetOwnerID, 
                @FirstName = @FirstName,
                @LastName = @LastName, 
                @BirthDate= @BirthDate,
                @PhysicalAddress = @PhysicalAddress,
                @MailingAddress = @MailingAddress, 
                @EmailAddress = @EmailAddress, 
                @PhoneNumber = @PhoneNumber;
            PRINT 'Person inserted successfully';
        END

        -- Now insert into PetOwner (which references PersonID)
        INSERT INTO people.PetOwner (PetOwnerID, VeterinarianID, HasSterilizedPets, HasVaccinatedPets, UsesVetHeartWormPrevention)
        VALUES (@PetOwnerID, @VeterinarianID, @PetsSterilized, @PetsVaccinated, @HeartWormPreventionFromVet);

        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH
        ROLLBACK TRANSACTION;
        PRINT 'Error occurred: ' + ERROR_MESSAGE();
    END CATCH;
END; 
GO
```

### INSERT Adopter

Creates a new adopter record, including the base person record if needed.

```sql
-- Inserts into the adopter table. If the AdopterID is not in the Person table it executes the insertPerson procedure and 
-- inserts into person table 
CREATE PROCEDURE INSERTADOPTER 
    
    @AdopterID UNIQUEIDENTIFIER,
    @PetAllergies BIT, 
    @HaveSurrendered BIT, 
    @HomeStatus VARCHAR (20),
    
    @FirstName NVARCHAR(50),
    @LastName NVARCHAR(50),
    @BirthDate DATE, 
    @PhysicalAddress NVARCHAR(225),
    @MailingAddress NVARCHAR (225),
    @EmailAddress NVARCHAR(100),
    @PhoneNumber NVARCHAR(20)
    
AS
BEGIN 
    SET NOCOUNT ON; 
    
    BEGIN TRY 
        BEGIN TRANSACTION; 
        --Ensures the person doesn't already exist 
        IF NOT EXISTS (SELECT 1 FROM people.Person WHERE PersonID = @AdopterID)
        BEGIN 
            EXEC INSERTPerson 
            @PersonID = @AdopterID, 
            @FirstName = @FirstName,
            @LastName = @LastName, 
            @BirthDate= @BirthDate,
            @PhysicalAddress = @PhysicalAddress,
            @MailingAddress = @MailingAddress, 
            @EmailAddress = @EmailAddress, 
            @PhoneNumber = @PhoneNumber;
        END
        --Inserts the adopter 
        INSERT INTO people.Adopter(AdopterID, HasPetAllergies, HasSurrenderedPets, HomeStatus)
        VALUES(@AdopterID, @PetAllergies, @HaveSurrendered, @HomeStatus);
        
        COMMIT TRANSACTION; 
        
    END TRY 
    BEGIN CATCH 
        ROLLBACK TRANSACTION;
        PRINT 'Error occurred: '+ ERROR_MESSAGE();
    END CATCH; 
    
END; 
GO
```

### INSERT Volunteer

Creates a new volunteer record, including the base person record if needed.

```sql
-- Inserts the Volunteer
CREATE PROCEDURE INSERTVOLUNTEER
    @VolunteerID UNIQUEIDENTIFIER, 
    @VolunteerPositon NVARCHAR(50),
    @StartDate DATE,
    @EndDate DATE, 
    @EmergencyContactName NVARCHAR(100),
    @EmergencyContactPhone NVARCHAR(20), 
    @IsActive BIT,
    
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
    
    BEGIN TRY 
        BEGIN TRANSACTION; 
        --Ensures the person doesnt already exist in the person table 
        IF NOT EXISTS (SELECT 1 FROM people.Person WHERE PersonID = @VolunteerID)
        BEGIN 
            EXEC INSERTPerson 
            @PersonID = @VolunteerID, 
            @FirstName = @FirstName,
            @LastName = @LastName, 
            @BirthDate= @BirthDate,
            @PhysicalAddress = @PhysicalAddress,
            @MailingAddress = @MailingAddress, 
            @EmailAddress = @EmailAddress, 
            @PhoneNumber = @PhoneNumber;
        END
        -- Inserts in the volunteer table     
        INSERT INTO people.Volunteer(VolunteerID, VolunteerPosition, StartDate, EndDate, EmergencyContactName, EmergencyContactPhone, IsActive)
        VALUES (@VolunteerID, @VolunteerPositon, @StartDate, @EndDate, @EmergencyContactName, @EmergencyContactPhone, @IsActive);
        
        COMMIT TRANSACTION;
    END TRY
    BEGIN CATCH 
        ROLLBACK TRANSACTION;
        PRINT 'Error occured: '+ ERROR_MESSAGE();
    END CATCH;
        
END;
GO
```

## Notes on Stored Procedures Usage

These stored procedures provide a convenient and consistent way to create new records in the system. They handle the complexity of creating related records in a single transaction, ensuring data integrity. Some key benefits include:

1. **Simplified Data Entry**: These procedures allow for creating complex related records with a single procedure call.

2. **Transactional Integrity**: All operations are wrapped in transactions to ensure either all operations succeed or none do.

3. **Business Logic Enforcement**: The procedures can enforce business rules and data validation.

4. **Error Handling**: Built-in error handling provides better diagnostics when issues occur.

## Usage Examples

### Adding a New Volunteer

```sql
EXEC INSERTVOLUNTEER
    @VolunteerID = 'A3F2E1D0-B9C8-4A7B-5E6F-7D8E9F0A1B2C', 
    @VolunteerPositon = 'Dog Walker',
    @StartDate = '2023-06-01',
    @EndDate = NULL, 
    @EmergencyContactName = 'Jane Smith',
    @EmergencyContactPhone = '555-123-4567', 
    @IsActive = 1,
    
    @FirstName = 'John',
    @LastName = 'Doe',
    @BirthDate = '1990-05-15', 
    @PhysicalAddress = '123 Main St, Shreveport, LA 71101',
    @MailingAddress = '123 Main St, Shreveport, LA 71101',
    @EmailAddress = 'john.doe@example.com',
    @PhoneNumber = '555-987-6543';
```

### Adding a New Adopter

```sql
EXEC INSERTADOPTER 
    @AdopterID = 'B4C5D6E7-F8A9-0B1C-2D3E-4F5A6B7C8D9E',
    @PetAllergies = 0, 
    @HaveSurrendered = 0, 
    @HomeStatus = 'Pending',
    
    @FirstName = 'Sarah',
    @LastName = 'Johnson',
    @BirthDate = '1985-08-22', 
    @PhysicalAddress = '456 Oak Ave, Bossier City, LA 71111',
    @MailingAddress = '456 Oak Ave, Bossier City, LA 71111',
    @EmailAddress = 'sarah.johnson@example.com',
    @PhoneNumber = '555-456-7890';
```