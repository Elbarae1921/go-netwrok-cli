package flags

import "github.com/urfave/cli"

// Host is a command flag
var Host cli.StringFlag = cli.StringFlag{
	Name:  "host",
	Value: "google.com",
}
