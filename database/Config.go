package database

import (
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"warranty/helpers"
	"warranty/logging"
)

// ErrMissingEnvVars indicates missing required database environment variables.
var ErrMissingEnvVars = errors.New("missing required database environment variables, check variable names")
var gormDB *gorm.DB // Package-level variable for the DB connection

func Connect() (*sql.DB, error) {
	helpers.LoadGodotEnv()

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if dbUser == "" || dbName == "" || dbHost == "" || dbPort == "" {
		return nil, fmt.Errorf("missing environment variables: DB_USER, DB_PASSWORD, DB_NAME, DB_HOST, or DB_PORT")
	}

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logging.Log("", fmt.Errorf("failed to open database connection: %v", err))
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		logging.Log("", fmt.Errorf("failed to ping database: %v", err))
		return nil, err
	}

	// Consider logging connection success only if desired
	// logging.Log("", nil)

	return db, nil
}

// ErrMissingEnvVars indicates missing required database environment variables.

// DB holds the GORM database connection

func GormConnect() (*gorm.DB, error) {
	// If a connection already exists, no need to modify it

	// Load environment variables
	helpers.LoadGodotEnv()

	// Retrieve database configuration from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Check if required environment variables are set
	if dbUser == "" || dbName == "" || dbHost == "" || dbPort == "" {
		return nil, ErrMissingEnvVars
	}

	// Create the Data Source Name (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Establish the database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.Log("", fmt.Errorf("failed to connect to the database: %v", err))
		return nil, err
	}

	logging.Log("", nil) // Log successful connection
	return db, nil
}

func GetDB() *gorm.DB {
	if gormDB == nil {
		var err error
		gormDB, err = GormConnect()
		if err != nil {
			// Handle the error appropriately here
			// (e.g., log the error or panic depending on your application's needs)
			panic(err) // Example for demonstration
		}
	}
	return gormDB
}
