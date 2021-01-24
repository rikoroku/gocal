package googlecalendar

import (
	"context"
	"encoding/json"
	"gocal/models"
	"log"
	"time"

	"google.golang.org/api/calendar/v3"
)

// Service is ...
type Service interface {
	GetEvents(fromDate, toDate time.Time) []*models.Event
}

type calendarService struct {
	service *calendar.Service
}

// NewService is ...
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

func (cs *calendarService) GetEvents(fromDate, toDate time.Time) []*models.Event {
	fetchedEvents, err := cs.service.Events.List("primary").SingleEvents(true).TimeMin(fromDate.Format(time.RFC3339)).TimeMax(toDate.Format(time.RFC3339)).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve events: %v", err)
	}
	if len(fetchedEvents.Items) == 0 {
		log.Fatalf("No events found.")
	}

	var events []*models.Event
	for _, item := range fetchedEvents.Items {
		event, err := calendarEventToEvent(item)
		if err != nil {
			log.Fatalf("An unexpected error has occurred: %v", err)
		}
		events = append(events, event)
	}

	return events
}

func calendarEventToEvent(calendarEvent *calendar.Event) (*models.Event, error) {
	var event models.Event
	byte, err := calendarEvent.MarshalJSON()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byte, &event)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
