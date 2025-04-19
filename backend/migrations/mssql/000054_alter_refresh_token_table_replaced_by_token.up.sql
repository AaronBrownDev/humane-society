-- Change ReplacedByToken to ReplacedByTokenID of type UNIQUEIDENTIFIER
ALTER TABLE auth.RefreshToken
DROP COLUMN ReplacedByToken;

ALTER TABLE auth.RefreshToken
ADD ReplacedByTokenID UNIQUEIDENTIFIER NULL;

-- Add self-referencing foreign key constraint
ALTER TABLE auth.RefreshToken
ADD CONSTRAINT FK_RefreshToken_ReplacedBy
FOREIGN KEY (ReplacedByTokenID)
REFERENCES auth.RefreshToken(TokenID);