package postgres

import (
	"database/sql"
	"log"
	"os"
)

const (
	host    = "localhost"
	port    = "5432"
	user    = "postgres"
	dbname  = "encrypt"
	sslMode = "disable"
)

// CheckDBConnection opens and verifies a connection to our database, logs any errors
func CheckDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host="+host+" port="+port+" user="+user+" "+
		"password="+getPwd()+" dbname="+dbname+" sslmode="+sslMode)
	if err != nil {
		panic(err)
	}

	// health check for db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Successfully connected!")
	return db, nil
}

// getPwd is a factory function to return postgres password from env
func getPwd() string {
	pwd := os.Getenv("POSTGRES_PASSWORD")
	return pwd
}
