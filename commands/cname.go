package commands

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// CNAME is a cli action
func CNAME(flags []cli.Flag) *cli.Command {

	// declare a command variable to return a pointer to it
	command := cli.Command{
		Name:  "cname",
		Usage: "Looks up CNAME for a particulare host",
		Flags: flags,
		Action: func(c *cli.Context) error {
			cname, err := net.LookupCNAME(c.String("host"))
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println(cname)
			return nil
		},
	}

	// return the address of command
	return &command
}
