package businesslogic

import "time"

const (
	EVENT_STATUS_DRAFT    = 1
	EVENT_STATUS_OPEN     = 2
	EVENT_STATUS_RUNNING  = 3
	EVENT_STATUS_CLOSED   = 4
	EVENT_STATUS_CANCELED = 5
)

type EventStatus struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

type IEventStatusRepository interface {
	GetEventStatus() ([]EventStatus, error)
}
