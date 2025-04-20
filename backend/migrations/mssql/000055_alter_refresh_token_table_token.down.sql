-- Add the Token column back
ALTER TABLE auth.RefreshToken
ADD Token NVARCHAR(255) NULL;

-- Update Token values to match TokenID for existing records
UPDATE auth.RefreshToken
SET Token = CAST(TokenID AS NVARCHAR(255));

-- Make Token NOT NULL after populating data
ALTER TABLE auth.RefreshToken
ALTER COLUMN Token NVARCHAR(255) NOT NULL;