-- ReadOnly role gets read-only access
GRANT SELECT ON SCHEMA::shelter TO HumaneSociety_ReadOnly;
GRANT SELECT ON SCHEMA::people TO HumaneSociety_ReadOnly;
GRANT SELECT ON SCHEMA::medical TO HumaneSociety_ReadOnly;