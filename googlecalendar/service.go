package googlecalendar

import (
	"context"
	"encoding/json"
	"gocal/models"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

type Service interface {
	GetEvents() []*models.Event
}

type calendarService struct {
	service *calendar.Service
}

func NewService() Service {
	tok, err := tokenFromTokFile()
	if err != nil {
		log.Fatalf("Couldn't get a refresh token. First, You must do setup command.")
	}

	config := getConfig()
	client := config.Client(context.Background(), tok)

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	return &calendarService{
		service: srv,
	}
}

func (cs *calendarService) GetEvents() []*models.Event {
	t := time.Now()
	timezone, _ := time.LoadLocation("Local")
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, timezone).Format(time.RFC3339)
	end := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 59, timezone).Format(time.RFC3339)

	fetchedEvents, err := cs.service.Events.List("primary").SingleEvents(true).TimeMin(start).TimeMax(end).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve today's user's events: %v", err)
	}
	if len(fetchedEvents.Items) == 0 {
		log.Fatalf("No today's events found.")
	}

	var events []*models.Event
	for _, item := range fetchedEvents.Items {
		var event models.Event
		byte, _ := item.MarshalJSON()
		err := json.Unmarshal(byte, &event)
		if err != nil {
			log.Fatalf("An unexpected error has occurred: %v", err)
		}
		events = append(events, &event)
	}
	// TODO: events(*calendar.Events) must be converted into models.Event

	return events
}
