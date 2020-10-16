package commands

import (
	"fmt"
	"net"

	"github.com/urfave/cli"
)

// IP is a cli action
func IP(flags []cli.Flag) *cli.Command {

	// declare a command variable to return a pointer to it
	command := cli.Command{
		Name:  "ip",
		Usage: "Looks up the IPs for a particular host",
		Flags: flags,
		Action: func(c *cli.Context) error {
			ip, err := net.LookupIP(c.String("host"))
			if err != nil {
				fmt.Println(err)
				return err
			}
			for _, i := range ip {
				fmt.Println(i.String())
			}
			return nil
		},
	}

	// return the address of command
	return &command
}
