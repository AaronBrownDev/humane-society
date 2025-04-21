ALTER TABLE auth.UserRole
ADD CONSTRAINT UQ_UserRole_UserID UNIQUE (UserID);
-- Only want users to have a single role