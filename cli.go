package main

import (
	"log"

	"cli/app"
	"cli/commands"
	"cli/flags"

	"github.com/urfave/cli"
)

func main() {

	// setup the imported flags into an array
	myFlags := []cli.Flag{
		&flags.Host,
	}

	// create an instance of App
	myApp := app.App{}

	// initialize the app with the imported commands
	myApp.New([]*cli.Command{
		commands.NS(myFlags),
		commands.IP(myFlags),
		commands.CNAME(myFlags),
		commands.MX(myFlags),
	})

	err := myApp.Run()

	// in case of error
	if err != nil {
		log.Fatal(err)
	}
}
