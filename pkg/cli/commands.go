package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func Create() cli.ActionFunc {
	// TODO
	fmt.Println("Created a new entry!")
	return nil
}

func Delete() cli.ActionFunc {
	// TODO
	fmt.Println("Deleted an entry!")
	return nil
}

func Edit() cli.ActionFunc {
	// TODO
	fmt.Println("Edited an entry!")
	return nil
}

func GetAll() cli.ActionFunc {
	// TODO
	fmt.Println("Shows all entries")
	return nil
}

func Get() cli.ActionFunc {
	// TODO
	fmt.Println("Retrieved an entry!")
	return nil
}
