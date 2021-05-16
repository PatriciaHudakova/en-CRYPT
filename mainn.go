package main

import (
	"en-CRYPT/pkg/postgres"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
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

	// create and run a new cli app with default values
	app := cli.NewApp()
	info(app)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// info populates non nil default app attributes
func info(app *cli.App) {
	app.Name = "en_Crypt"
	app.Authors = []*cli.Author{{
		Name:  "Patricia Hudakova",
		Email: "phudakova7@gmail.com",
	}}
	app.Description = "en-Crypt is a simple password management store"
	app.Version = "1.0"
}
