package cli

import (
	"fmt"

	cgCli "github.com/codegangsta/cli"
	"github.com/toorop/tmail/api"
)

var RelayIP = &cgCli.Command{
	Name:  "relayip",
	Usage: "commands to authorise IP to relay through tmail",
	Subcommands: []*cgCli.Command{
		// Add an authorized IP
		{
			Name:        "add",
			Usage:       "Add an authorized IP",
			Description: "tmail relayip add IP",
			Action: func(c *cgCli.Context) (err error) {
				if c.NArg() == 0 {
					cliDieBadArgs(c)
				}
				cliHandleErr(api.RelayIpAdd(c.Args().First()))
				return nil
			},
		},
		// List authorized IPs
		{
			Name:        "list",
			Usage:       "List authorized IP",
			Description: "tmail relayip list",
			Action: func(c *cgCli.Context) error {
				ips, err := api.RelayIpGetAll()
				cliHandleErr(err)
				if len(ips) == 0 {
					println("There no athorized IP.")
				} else {
					for _, ip := range ips {
						fmt.Println(fmt.Sprintf("%d %s", ip.Id, ip.Ip))
					}
				}
				return nil
			},
		},

		// Delete relayip
		{
			Name:        "del",
			Usage:       "Delete an authorized IP",
			Description: "tmail relayip del IP",
			Action: func(c *cgCli.Context) error {
				if c.NArg() == 0 {
					cliDieBadArgs(c)
				}
				err := api.RelayIpDel(c.Args().First())
				cliHandleErr(err)
				return nil
			},
		},
	},
}
