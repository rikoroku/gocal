package cmd

import (
	"gocal/googlecalendar"

	"github.com/spf13/cobra"
)

func agendaCmd(service googlecalendar.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "agenda",
		Short: "Get events for a time period from start-time(default is 12am today) to end-time(default is 5 days from start).",
		Run: func(cmd *cobra.Command, args []string) {
			for _, event := range service.GetEvents() {
				cmd.Println(event.ID)
			}
		},
	}
}
