-- Revert back to ReplacedByToken as NVARCHAR
ALTER TABLE auth.RefreshToken
DROP CONSTRAINT FK_RefreshToken_ReplacedBy;

ALTER TABLE auth.RefreshToken
DROP COLUMN ReplacedByTokenID;

ALTER TABLE auth.RefreshToken
ADD ReplacedByToken NVARCHAR(255) NULL;