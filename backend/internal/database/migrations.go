package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// RunMigrations executes all database migrations
func RunMigrations() error {
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	// Get migration directory from environment or use default
	migrationDir := os.Getenv("MIGRATION_DIR")
	if migrationDir == "" {
		// Default to migrations/mssql relative to working directory
		workDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get working directory: %w", err)
		}
		migrationDir = filepath.Join(workDir, "migrations", "mssql")
	}

	log.Printf("Running migrations from: %s", migrationDir)

	// Create driver instance for the migrate package
	driver, err := sqlserver.WithInstance(db, &sqlserver.Config{})
	if err != nil {
		return fmt.Errorf("failed to create migration driver: %w", err)
	}

	// Create migration instance
	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", migrationDir),
		"sqlserver", driver)
	if err != nil {
		return fmt.Errorf("failed to create migration instance: %w", err)
	}

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	log.Println("Migrations completed successfully")
	return nil
}

// EnsureDatabaseExists creates the database if it doesn't exist
func EnsureDatabaseExists() error {
	// Connect to master database
	connectionString := getMasterConnectionString()
	masterDB, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		return fmt.Errorf("error connecting to master db: %w", err)
	}
	defer masterDB.Close()

	dbName := os.Getenv("DB_NAME")
	// Check if database exists and create if needed
	query := fmt.Sprintf("IF NOT EXISTS (SELECT * FROM sys.databases WHERE name = '%s') CREATE DATABASE [%s]", dbName, dbName)
	_, err = masterDB.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating database: %w", err)
	}

	return nil
}

// Helper function to connect to master database
func getMasterConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	return fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=master;TrustServerCertificate=true",
		host, port, user, password)
}
