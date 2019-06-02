package commands

import "github.com/mkideal/cli"

type listT struct{}

var List = &cli.Command{
	Name:    "list",
	Desc:    "display today's schedules",
	Aliases: []string{"l"},
	Fn: func(_ *cli.Context) error {
		return nil
	},
}
