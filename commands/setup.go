package commands

import "github.com/mkideal/cli"

type setupT struct{}

var Setup = &cli.Command{
	Name: "setup",
	Desc: "do only the first time. enable Google Calendar API",
	Fn: func(_ *cli.Context) error {
		return nil
	},
}
