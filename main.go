package main

import (
	"en-CRYPT/pkg/cli"
	"en-CRYPT/pkg/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// load variables from .env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	// connect and ping the database to check healthiness
	if db, err := postgres.CheckDBConnection(); err != nil {
		log.Fatal(err)
		return
	} else {
		// wrapper class to create and populate a new cli app
		cli.NewApp(db)

		//defer a graceful shutdown
		defer db.Close()
	}
}
