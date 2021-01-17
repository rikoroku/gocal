package main

import (
	"fmt"
	"gocal/cmd"
	"os"

	"github.com/mkideal/cli"
)

var help = cli.HelpCommand("Display help of command information.")

type rootT struct {
	cli.Helper
	Version bool `cli:"v,version" usage:"Display version"`
}

func (r *rootT) AutoHelp() bool {
	return r.Help
}

var root = &cli.Command{
	Desc: "\ngocal is a schedule management tool of Google Calendar",
	Argv: func() interface{} { return new(rootT) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*rootT)
		if argv.Version {
			ctx.String("%v\n", appVersion)
			return nil
		}
		return nil
	},
}

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(cmd.Setup),
		cli.Tree(cmd.List),
		cli.Tree(cmd.Insert),
		cli.Tree(cmd.Delete),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
