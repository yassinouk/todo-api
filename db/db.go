// db/db.go
package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // Importing the PostgreSQL driver package anonymously
)

// ConnectDB establishes a connection to a PostgreSQL database and returns the *sql.DB object.
// It takes database connection parameters from environment variables.
func ConnectDB() (*sql.DB, error) {
	// Retrieve database connection parameters from environment variables
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	// Create a connection string using the retrieved parameters
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a connection to the PostgreSQL database using the "postgres" driver
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		// Return an error if the connection cannot be established
		return nil, err
	}

	// Return the established database connection and nil error if successful
	return db, nil
}
