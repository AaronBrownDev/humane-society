package database

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

var db *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found. Using system environment variables.")
	}
}

func InitializeDB() (err error) {

	driverName := "sqlserver"
	connectionString := getConnectionString()

	db, err = sql.Open(driverName, connectionString)
	if err != nil {
		return fmt.Errorf("error opening database: %s", err)
	}

	configureConnectionPool()

	if err = db.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %s", err)
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() (err error) {
	if db == nil {
		return fmt.Errorf("no existing database connection to close")
	}
	err = db.Close()
	if err != nil {
		return fmt.Errorf("error closing database: %s", err)
	}
	return

}

func getConnectionString() string {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")

	return fmt.Sprintf("server=%s;port=%s;user id=%s;password=%s;database=%s;trustservercertificate=true", host, port, user, password, databaseName)
}

func configureConnectionPool() {
	connMaxLifetime := getEnvAsInt("DB_CONN_MAX_LIFE_TIME", 5)
	maxIdleConns := getEnvAsInt("DB_MAX_IDLE_CONNS", 10)
	maxOpenConns := getEnvAsInt("DB_MAX_OPEN_CONNS", 25)

	db.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Minute)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(valueStr)
	if err != nil {
		log.Printf("Error parsing %s as int: %v", key, err)
		return defaultValue
	}
	return value
}
