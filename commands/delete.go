package commands

import "github.com/mkideal/cli"

type deleteT struct {
	Id string `cli:"*id" usage:"target id to delete"`
}

var Delete = &cli.Command{
	Name:    "delete",
	Desc:    "deletes a schedule of specified id",
	Aliases: []string{"d"},
	Argv:    func() interface{} { return new(deleteT) },
	Fn: func(_ *cli.Context) error {
		return nil
	},
}
