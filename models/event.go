package models

import "google.golang.org/api/calendar/v3"

// Event is ...
type Event struct {
	ID          string                  `json:"id,omitempty"`
	Summary     string                  `json:"summary,omitempty"`
	Description string                  `json:"description,omitempty"`
	Kind        string                  `json:"kind,omitempty"`
	Start       *calendar.EventDateTime `json:"start,omitempty"`
	End         *calendar.EventDateTime `json:"end,omitempty"`
}
