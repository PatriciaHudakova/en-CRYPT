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
	if err := postgres.CheckDBConnection(); err != nil {
		log.Fatal(err)
	}

	// wrapper class to create and populate a new cli app
	cli.NewApp()
}
