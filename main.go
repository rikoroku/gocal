package main

import (
	"fmt"
	"os"

	"github.com/mkideal/cli"
)

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(lsit),
		cli.Tree(insert),
		cli.Tree(delete),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var help = cli.HelpCommand("display help of command information")

type argT struct {
	cli.Helper
	Version bool `cli:"v,version" usage:"display version"`
}

func (a *argT) AutoHelp() bool {
	return a.Help
}

var root = &cli.Command{
	Desc: "\nGoedule is a schedule management tool of Google Calendar",
	Argv: func() interface{} { return new(argT) },
	Fn: func(_ *cli.Context) error {
		return nil
	},
}

type listT struct{}

var lsit = &cli.Command{
	Name:    "list",
	Desc:    "display today's schedules",
	Aliases: []string{"l"},
	Fn: func(_ *cli.Context) error {
		return nil
	},
}

type insertT struct {
	Start string `cli:"*start,s" usage:"start time"`
	End   string `cli:"*e" usage:"end time"`
	Title string `cli:"*title, t" usage:"schedule's title"`
}

var insert = &cli.Command{
	Name:    "insert",
	Desc:    "creates a new schedule",
	Aliases: []string{"i"},
	Argv:    func() interface{} { return new(insertT) },
	Fn: func(_ *cli.Context) error {
		return nil
	},
}

type deleteT struct {
	Id string `cli:"*id" usage:"target id to delete"`
}

var delete = &cli.Command{
	Name:    "delete",
	Desc:    "deletes a schedule of specified id",
	Aliases: []string{"d"},
	Argv:    func() interface{} { return new(deleteT) },
	Fn: func(_ *cli.Context) error {
		return nil
	},
}
