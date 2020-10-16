package app

import (
	"os"

	"github.com/urfave/cli"
)

// App is the main struct for setting up the cli
type App struct {
	app   *cli.App
	flags []cli.Flag
}

// this function initializes the cli.App commands
func (a *App) initCommands(c []*cli.Command) {
	a.app.Commands = c
	return
}

// New is a method to initialize the cli app
func (a *App) New(c []*cli.Command) {
	a.app = cli.NewApp()

	a.app.Name = "Website Lookup CLI"
	a.app.Usage = "Query IPs, CNAMES, MX records and Name Servers!"

	a.initCommands(c)
	return
}

// Run is a method to start up the cli app
func (a App) Run() error {
	return a.app.Run(os.Args)
}
