package helpers

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"warranty/logging"
)

var ErrMissingEnvVars = errors.New("missing required database environment variables, check variable names")

func ConnectMysql() *sql.DB {
	LoadGodotEnv()
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	fmt.Println(dbName)

	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbName == "" {
		return nil
		logging.Log("", ErrMissingEnvVars) // Return the pre-defined error
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logging.Log("", err)
		return nil
	}

	err = db.Ping()
	if err != nil {
		logging.Log("", err)
		return nil
	}

	fmt.Println("Database connection successful!")

	return db
}
