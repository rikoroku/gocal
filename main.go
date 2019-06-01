package main

import (
	"os"

	"github.com/mkideal/cli"
)

type argT struct {
	cli.Helper
}

func main() {
	os.Exit(cli.Run(new(argT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argT)
		ctx.String("%v", argv)
		return nil
	}))
}
