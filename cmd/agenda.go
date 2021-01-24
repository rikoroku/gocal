package cmd

import (
	"gocal/googlecalendar"
	"time"

	"github.com/spf13/cobra"
)

func agendaCmd(service googlecalendar.Service) *cobra.Command {
	return &cobra.Command{
		Use:   "agenda",
		Short: "Get events for a time period from start-time(default is 12am today) to end-time(default is 5 days from start).",
		Run: func(cmd *cobra.Command, args []string) {
			fromDate, toDate := generateFromToDates()
			for _, event := range service.GetEvents(fromDate, toDate) {
				cmd.Println(event.ID)
			}
		},
	}
}

func generateFromToDates() (time.Time, time.Time) {
	t := time.Now()
	timezone, _ := time.LoadLocation("Local")
	fromDate := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, timezone)
	toDate := time.Date(t.Year(), t.Month(), t.Day()+5, 23, 59, 59, 59, timezone)
	return fromDate, toDate
}
