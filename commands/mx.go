package commands

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// MX is a cli action
func MX(flags []cli.Flag) *cli.Command {

	// declare a command variable to return a pointer to it
	command := cli.Command{
		Name:  "mx",
		Usage: "Looks up MX records for a particular host",
		Flags: flags,
		Action: func(c *cli.Context) error {
			mx, err := net.LookupMX(c.String("host"))
			if err != nil {
				fmt.Println(err)
				return err
			}
			for _, i := range mx {
				fmt.Println(i.Host, i.Pref)
			}
			return nil
		},
	}

	// return the address of command
	return &command
}
