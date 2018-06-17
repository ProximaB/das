package businesslogic

import (
	"time"
)

// EventDance represents the many-to-many relationship between competition event and dance references.
type EventDance struct {
	ID              int
	EventID         int
	DanceID         int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type SearchEventDanceCriteria struct {
	EventDanceID  int
	CompetitionID int
	EventID       int
}

type IEventDanceRepository interface {
	SearchEventDance(criteria SearchEventDanceCriteria) ([]EventDance, error)
	CreateEventDance(eventDance *EventDance) error
	DeleteEventDance(eventDance EventDance) error
	UpdateEventDance(eventDance EventDance) error
}

func NewEventDance(event Event, danceID int) *EventDance {
	return &EventDance{
		EventID:         event.ID,
		DanceID:         danceID,
		CreateUserID:    event.CreateUserID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    event.UpdateUserID,
		DateTimeUpdated: time.Now(),
	}
}
