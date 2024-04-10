package helpers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"
)

func MigrateFromFile(db *sql.DB, filepath string) error {
	// Check for nil database connection
	if db == nil {
		return fmt.Errorf("nil database connection provided")
	}

	// Read the SQL file content
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err) // Wrap error for context
	}

	// Split the SQL statements, adjusting for delimiter if needed
	statements := getStatements(string(data)) // Using a separate function for clarity

	// Transaction for atomicity (optional)
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer func() {
		if err := recover(); err != nil {
			// Rollback on panic
			_ = tx.Rollback()
			panic(err) // Re-throw panic
		} else if err != nil {
			// Rollback on errors
			_ = tx.Rollback()
		} else {
			// Commit on success
			err = tx.Commit()
		}
	}()

	for _, stmt := range statements {
		// Execute each statement within the transaction
		_, err := tx.Exec(stmt)
		if err != nil {
			// Handle specific error types (optional)
			if isIgnorableError(err) {
				fmt.Printf("Warning: Statement '%s' might not have affected any rows.\n", stmt)
				continue
			}

			return fmt.Errorf("failed to execute SQL statement: %w", err) // Wrap error
		}
	}

	return nil // Commit will happen in the deferred function
}

// Helper functions for clarity and customization

// getStatements extracts statements, adjusting delimiter if needed
func getStatements(sql string) []string {
	// Replace with logic for custom delimiter handling if required
	return strings.Split(sql, ";")
}

// isIgnorableError determines if an error can be ignored
func isIgnorableError(err error) bool {
	// Customize criteria for ignoring errors (e.g., sql.ErrNoRows)
	return err == sql.ErrNoRows
}

// MigrateAll creates the "users" table and a trigger function for automatic timestamp updates
//func MigrateAll(db *sql.DB) error {
//	createTableSQL := `
//CREATE TABLE IF NOT EXISTS users (
//  id SERIAL PRIMARY KEY,
//  username VARCHAR(50) NOT NULL UNIQUE,
//  email VARCHAR(100) NOT NULL UNIQUE,
//  password_hash VARCHAR(255) NOT NULL,
//  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- No need for separate default here
//  first_name VARCHAR(50) DEFAULT NULL,
//  last_name VARCHAR(50) DEFAULT NULL,
//  profile_pic_url VARCHAR(255) DEFAULT NULL,
//  verified BOOLEAN DEFAULT FALSE,
//  UNIQUE (username, email)
//);
//`
//
//	createTriggerSQL := `
//CREATE OR REPLACE FUNCTION update_timestamp()
//RETURNS TRIGGER AS $$
//BEGIN
//  IF NEW.updated_at IS NULL THEN  -- Update only if not explicitly set
//    NEW.updated_at = CURRENT_TIMESTAMP;
//  END IF;
//  RETURN NEW;
//END;
//$$ LANGUAGE plpgsql;
//
//CREATE TRIGGER update_users_timestamp
//BEFORE UPDATE ON users
//FOR EACH ROW
//EXECUTE PROCEDURE update_timestamp();  -- Use PROCEDURE instead of FUNCTION
//`
//
//	_, err := db.Exec(createTableSQL)
//	if err != nil {
//		logging.Log(err)
//		return fmt.Errorf("failed to create 'users' table: %w", err)
//	}
//
//	_, err = db.Exec(createTriggerSQL)
//	if err != nil {
//		logging.Log(err)
//		return fmt.Errorf("failed to create trigger: %w", err)
//	}
//
//	fmt.Println("Users table and trigger migrated successfully!")
//	return nil
//}
