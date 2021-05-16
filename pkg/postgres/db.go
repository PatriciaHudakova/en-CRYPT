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
func CheckDBConnection() error {
	db, err := sql.Open("postgres", "host="+host+" port="+port+" user="+user+" "+
		"password="+getPwd()+" dbname="+dbname+" sslmode="+sslMode)
	if err != nil {
		panic(err)
	}

	// defer graceful shutdown
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("unable to gracefully close db connection: %s", err)
			return
		}
	}(db)

	// verify postgres is able to accept connections
	err = db.Ping()
	if err != nil {
		return err
	}

	log.Println("Successfully connected!")
	return nil
}

// getPwd is a factory function to return postgres password from env
func getPwd() string {
	pwd := os.Getenv("POSTGRES_PASSWORD")
	return pwd
}
