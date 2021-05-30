package cli

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/shomali11/xsql"
	"github.com/urfave/cli/v2"
	"log"
)

func Create() cli.ActionFunc {
	// TODO
	return nil
}

func Delete() cli.ActionFunc {
	// TODO
	return nil
}

func Edit() cli.ActionFunc {
	// TODO
	return nil
}

func GetAll(db *sql.DB) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		// initialised timeout to prevent request from hanging
		ctx, cancel := context.WithTimeout(c.Context, commandTimeout)
		defer cancel()

		// build and execute query
		rows, err := db.QueryContext(ctx, "SELECT * FROM credentials;")
		if err != nil {
			log.Fatalf("unable to return queried rows: %v", err)
			return err
		}
		defer rows.Close()

		// using the library to display all rows in a simulated table
		entries, err := xsql.Pretty(rows)
		if err != nil {
			log.Fatalf("unable to display results: %v", err)
			return err
		}

		fmt.Println(entries)
		return nil
	}
}

func Get() cli.ActionFunc {
	// TODO
	return nil
}
