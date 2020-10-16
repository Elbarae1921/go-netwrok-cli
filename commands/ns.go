package commands

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// NS is a cli command
func NS(flags []cli.Flag) *cli.Command {

	// declare a command variable to return a pointer to it
	command := cli.Command{
		Name:  "ns",
		Usage: "Looks up the name servers for a particular host",
		Flags: flags,
		Action: func(c *cli.Context) error {
			ns, err := net.LookupNS(c.String("host"))
			if err != nil {
				fmt.Println(err)
				return err
			}
			for _, i := range ns {
				fmt.Println(i.Host)
			}
			return nil
		},
	}

	// return the address of command
	return &command
}
