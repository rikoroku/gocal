package cmd

import (
	"gocal/googlecalendar"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setupCmd)
}

var setupCmd = &cobra.Command{
	Use:   "setup",
	Short: "Enable Google Calendar API only the first time.",
	Run: func(cmd *cobra.Command, args []string) {
		googlecalendar.Setup()
	},
}
