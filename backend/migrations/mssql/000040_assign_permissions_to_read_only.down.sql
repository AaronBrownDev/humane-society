-- Revoke permissions from ReadOnly role
REVOKE SELECT ON SCHEMA::shelter FROM HumaneSociety_ReadOnly;
REVOKE SELECT ON SCHEMA::people FROM HumaneSociety_ReadOnly;
REVOKE SELECT ON SCHEMA::medical FROM HumaneSociety_ReadOnly;