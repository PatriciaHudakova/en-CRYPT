package cli

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func NewApp() {
	app := cli.NewApp()
	info(app)
	commands(app)

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

func commands(app *cli.App) {
	app.Commands = []*cli.Command{
		{
			Name:    "Create",
			Aliases: []string{"c"},
			Usage:   "Creates a new credential entry in the database",
			Action:  Create(),
		},
		{
			Name:    "Delete",
			Aliases: []string{"d"},
			Usage:   "Deletes credential entry from the database",
			Action:  Delete(),
		},
		{
			Name:    "Edit",
			Aliases: []string{"e"},
			Usage:   "Edits an existing credential entry in the database",
			Action:  Edit(),
		},
		{
			Name:    "Get All",
			Aliases: []string{"ga"},
			Usage:   "Displays all credentials from the database",
			Action:  GetAll(),
		},
		{
			Name:    "Get",
			Aliases: []string{"g"},
			Usage:   "Gets a specific entry from the database by domain",
			Action:  Get(),
		},
	}
}
