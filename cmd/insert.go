package cmd

import "github.com/mkideal/cli"

type insertT struct {
	Start string `cli:"*start,s" usage:"start time"`
	End   string `cli:"*e" usage:"end time"`
	Title string `cli:"*title, t" usage:"schedule's title"`
}

var Insert = &cli.Command{
	Name:    "insert",
	Desc:    "creates a new schedule",
	Aliases: []string{"i"},
	Argv:    func() interface{} { return new(insertT) },
	Fn: func(_ *cli.Context) error {
		return nil
	},
}
