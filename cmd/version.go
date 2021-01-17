package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gocal",
	Long:  `All software has versions. This is gocal's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gocal v1.0 -- HEAD")
	},
}
