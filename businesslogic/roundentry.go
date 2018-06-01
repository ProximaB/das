package businesslogic

import "time"

type RoundEntry struct {
	ID              int
	RoundID         int
	EventEntryID    int
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}
