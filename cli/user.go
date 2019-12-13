package cli

import (
	cgCli "github.com/codegangsta/cli"
	"github.com/toorop/tmail/api"
)

var user = &cgCli.Command{
	Name:  "user",
	Usage: "commands to manage users of mailserver",
	Subcommands: []*cgCli.Command{
		// users
		{
			Name:        "add",
			Usage:       "Add an user",
			Description: "tmail user add [--mailbox] [--relay] [--quota BYTES] [--catchall] USER CLEAR_PASSWD ",
			Flags: []cgCli.Flag{
				&cgCli.BoolFlag{
					Name:  "mailbox",
					Usage: "Create a mailbox for this user.",
				},
				&cgCli.BoolFlag{
					Name:  "relay",
					Usage: "Authorise user to use server as SMTP relay.",
				},
				&cgCli.BoolFlag{
					Name:  "catchall",
					Usage: "Set this user as catchall for domain",
				},
				&cgCli.StringFlag{
					Name:  "quota",
					Value: "1G",
					Usage: "Mailbox quota in bytes (not bits). You can use K,M,G as unit. Eg: 10G mean a quota of 10GB",
				},
			},
			Action: func(c *cgCli.Context) error {
				var err error
				if c.NArg() < 2 {
					cliDieBadArgs(c)
				}
				err = api.UserAdd(c.Args().First(), c.Args().Get(1), c.String("quota"), c.Bool("mailbox"), c.Bool("relay"), c.Bool("catchall"))
				cliHandleErr(err)
				cliDieOk()
				return nil
			},
		},
		{
			Name:        "del",
			Usage:       "Delete an user",
			Description: "tmail user del USER",
			Action: func(c *cgCli.Context) error {
				var err error
				if c.NArg() != 1 {
					cliDieBadArgs(c)
				}
				err = api.UserDel(c.Args().First())
				cliHandleErr(err)
				cliDieOk()
				return nil
			},
		},
		// Update to change proprieties of an user
		// for now only password change is handled
		{
			Name:        "update",
			Usage:       "change proprieties of an user",
			Description: "tmail user update USER -p NEW_PASSWORD",
			Flags: []cgCli.Flag{
				&cgCli.StringFlag{
					Name:  "password, p",
					Usage: "update user password",
				},
			},
			Action: func(c *cgCli.Context) error {
				if c.NArg() != 1 {
					cliDieBadArgs(c)
				}
				if c.String("p") != "" {
					cliHandleErr(api.UserChangePassword(c.Args().First(), c.String("p")))
					cliDieOk()
				}
				cliDieBadArgs(c)
				return nil
			},
		},
		{
			Name:        "list",
			Usage:       "Return a list of users",
			Description: "",
			Action: func(c *cgCli.Context) error {
				users, err := api.UserGetAll()
				cliHandleErr(err)
				if len(users) == 0 {
					println("There is no users yet.")
					return err
				}
				for _, user := range users {
					line := user.Login + " - authrelay: "
					if user.AuthRelay {
						line += "yes"
					} else {
						line += "no"
					}
					line += " - have mailbox: "
					if user.HaveMailbox {
						line += "yes - home: " + user.Home
					} else {
						line += "no"
					}
					if user.Active == "Y" {
						line += " - active: yes"
					} else {
						line += " - active: no"
					}
					if user.IsCatchall {
						line += " - catchall: yes"
					} else {
						line += " - catchall: no"
					}
					println(line)
				}
				return nil
			},
		},
	},
}
