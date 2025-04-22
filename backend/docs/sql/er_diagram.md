```mermaid
erDiagram
    Person ||--o{ Adopter : "is a"
    Person ||--o{ Volunteer : "is a"
    Person ||--o{ PetOwner : "is a"
    Person ||--o{ PetSurrenderer : "is a"
    Person ||--o{ Veterinarian : "is a"
    Person ||--o{ UserAccount : "has"

    UserAccount ||--o{ UserRole : "has"
    UserAccount ||--o{ RefreshToken : "has"
    Role ||--o{ UserRole : "assigned to"
    RefreshToken ||--o| RefreshToken : "replaced by"

    Dog ||--o{ DogPrescription : "receives"
    Medicine ||--o{ DogPrescription : "used in"
    Veterinarian ||--o{ DogPrescription : "prescribes"

    ItemCatalog ||--o{ Supply : "categorizes"

    PetOwner ||--o{ PetOwnerPet : "owns"
    Veterinarian ||--o{ PetOwner : "treats pets of"

    PetSurrenderer ||--o{ SurrenderForm : "submits"
    Volunteer ||--o{ SurrenderForm : "processes"
    Veterinarian ||--o{ SurrenderForm : "provides history for"
    Dog ||--o{ SurrenderForm : "results from"

    Adopter ||--o{ AdoptionForm : "submits"
    Dog ||--o{ AdoptionForm : "is subject of"
    Volunteer ||--o{ AdoptionForm : "processes"

    Person ||--o{ VolunteerForm : "submits"
    Volunteer ||--o{ VolunteerForm : "processes"

    Volunteer ||--o{ VolunteerSchedule : "is scheduled for"

    ChangeLog }|--|| Dog : "tracks changes to"

    Person {
uuid PersonID PK
string FirstName
string LastName
date BirthDate NULL
        string PhysicalAddress
string MailingAddress
string EmailAddress
string PhoneNumber NULL
}

Adopter {
uuid AdopterID PK, FK
bool HasPetAllergies
bool HasSurrenderedPets
string HomeStatus
}

Volunteer {
uuid VolunteerID PK, FK
string VolunteerPosition
date StartDate
date EndDate NULL
string EmergencyContactName NULL
string EmergencyContactPhone NULL
bool IsActive
}

PetOwner {
uuid PetOwnerID PK, FK
uuid VeterinarianID FK NULL
bool HasSterilizedPets
bool HasVaccinatedPets
bool UsesVetHeartWormPrevention
    }

PetSurrenderer {
uuid SurrendererID PK, FK
}

Veterinarian {
uuid VeterinarianID PK, FK
}

Dog {
uuid DogID PK
string Name
date IntakeDate
date EstimatedBirthDate
string Breed
string Sex
string Color
int CageNumber
bool IsAdopted
}

Medicine {
int MedicineID PK
string Name
string Manufacturer
string Description NULL
string DosageUnit NULL
}

DogPrescription {
int PrescriptionID PK
uuid DogID FK
int MedicineID FK
decimal Dosage
string Frequency NULL
date StartDate
date EndDate NULL
string Notes NULL
uuid VetPrescriberID FK NULL
}

ItemCatalog {
uuid ItemID PK
string Name
string Category
string Description NULL
int MinimumQuantity
bool IsActive
    }

Supply {
int SupplyID PK
uuid ItemID FK
int Quantity
string StorageLocation NULL
date ExpirationDate NULL
string BatchNumber NULL
date AcquisitionDate NULL
}

PetOwnerPet {
int PetID PK
uuid PetOwnerID FK
string Name
string Type
string Breed
string Sex
date OwnershipDate
string LivingEnvironment
}

SurrenderForm {
int SurrenderFormID PK
uuid SurrendererID FK
datetime SubmissionDate
string DogName
int DogAge
decimal WeightInPounds
string Sex
string Breed NULL
string Color NULL
string LivingEnvironment
date OwnershipDate
uuid VeterinarianID FK NULL
date LastVetVisitDate NULL
bool IsGoodWithChildren
bool IsGoodWithDogs
bool IsGoodWithCats
bool IsGoodWithStrangers
bool IsHouseTrained
bool IsSterilized
string MicroChipNumber NULL
string MedicalProblems NULL
string BiteHistory NULL
string SurrenderReason
uuid ProcessedByVolunteerID FK NULL
datetime ProcessingDate NULL
uuid ResultingDogID FK NULL
string Status
}

AdoptionForm {
int AdoptionFormID PK
uuid AdopterID FK
uuid DogID FK
datetime SubmissionDate
uuid ProcessedByVolunteerID FK NULL
datetime ProcessingDate NULL
string Status
string RejectionReason NULL
}

VolunteerForm {
int VolunteerFormID PK
uuid ApplicantID FK
datetime SubmissionDate
bool SupportsAnimalWelfareEducation
string AvailableShifts NULL
bool SupportsResponsibleBreeding
bool AcceptsCleaningDuties
bool AcceptsDogCare
bool HasDogAllergies
bool HasPhysicalLimitations
bool IsForCommunityService
int RequiredServiceHours NULL
string ReferralSource
string CommentsAndQuestions NULL
uuid ProcessedByVolunteerID FK NULL
datetime ProcessingDate NULL
string Status
string RejectionReason NULL
}

VolunteerSchedule {
int ScheduleID PK
uuid VolunteerID FK
date ScheduleDate
time StartTime
time EndTime
string TaskDescription NULL
string Status
}

UserAccount {
uuid UserID PK, FK
string PasswordHash
datetime LastLogin NULL
bool IsActive
int FailedLoginAttempts
bool IsLocked
datetime LockoutEnd NULL
datetime CreatedAt
}

Role {
int RoleID PK
string Name
string Description NULL
}

UserRole {
uuid UserID PK, FK
int RoleID PK, FK
datetime AssignedAt
}

RefreshToken {
uuid TokenID PK
uuid UserID FK
datetime Expires
datetime CreatedAt
datetime RevokedAt NULL
uuid ReplacedByTokenID FK NULL
}

ChangeLog {
int LogID PK
string TableName
string PrimaryKeyColumn
string PrimaryKeyValue
string ColumnName
string OldValue NULL
string NewValue NULL
datetime ChangeDate
string ChangedBy
char AuditActionType
}
```