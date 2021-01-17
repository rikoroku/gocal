package cmd

import (
	"fmt"
	"gocal/googlecalendar"

	"github.com/mkideal/cli"
)

type listT struct{}

var List = &cli.Command{
	Name:    "list",
	Desc:    "display today's schedules",
	Aliases: []string{"l"},
	Fn: func(_ *cli.Context) error {
		events := googlecalendar.GetEvents()
		for _, item := range events.Items {
			if item.Start.DateTime != "" {
				fmt.Printf("id: %v %v (%v - %v)\n", item.Id, item.Summary, item.Start.DateTime, item.End.DateTime)
			}
		}
		return nil
	},
}
