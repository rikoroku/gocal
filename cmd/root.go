package cmd

import (
	"gocal/googlecalendar"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gocal",
	Short: "gocal is a golang application that allows you to access your Google Calendar from a command line. It's easy to manage your events.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() error {
	rootCmd.AddCommand(agendaCmd(googlecalendar.NewService()))
	return rootCmd.Execute()
}
