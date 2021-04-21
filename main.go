package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = "5432"
	user = "postgres"
	dbname = "passwords"
	sslmode = "disable"
)

func main() {
	// load variables from .env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	if err := checkDBConnection(); err != nil {
		log.Fatal(err)
	}
}

// opens and verifies a connection to our database, logs any errors
func checkDBConnection() error {
	db, err := sql.Open("postgres", "host="+host+" port="+port+" user="+user+" "+
		"password="+getPwd()+" dbname="+dbname+" sslmode="+sslmode)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// verify postgres is able to accept connections
	err = db.Ping()
	if err != nil {
		return err
	}

	log.Println("Successfully connected!")

	return nil
}

// factory function to return postgres password
func getPwd() string {
	pwd := os.Getenv("POSTGRES_PASSWORD")
	return pwd
}