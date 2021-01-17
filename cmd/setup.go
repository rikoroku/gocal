package cmd

import (
	"gocal/googlecalendar"

	"github.com/mkideal/cli"
)

type setupT struct{}

var Setup = &cli.Command{
	Name: "setup",
	Desc: "To enable Google Calendar API only the first time.",
	Fn: func(_ *cli.Context) error {
		googlecalendar.Setup()
		return nil
	},
}
