package businesslogic

import (
	"time"
)

// EventEntry defines the base struct of an Entry at an Event
type EventEntry struct {
	ID              int
	EventID         int
	CheckInTime     time.Time
	CompetitorTag   int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

// AdjudicatorEventEntry defines the participation of an Adjudicator at an Event.
type AdjudicatorEventEntry struct {
	ID            int
	EventEntry    EventEntry
	AdjudicatorID int
}

// SearchAdjudicatorEventEntryCriteria specifies the parameters that can be used to search the Event Entry of a
// Adjudicator in DAS
type SearchAdjudicatorEventEntryCriteria struct {
	CompetitionID int `schema:"competition"`
	EventID       int
	PartnershipID int
	Federation    int `schema:"federation"`
	Division      int `schema:"division"`
	Age           int `schema:"age"`
	Proficiency   int `schema:"proficiency"`
	Style         int `schema:"style"`
}

type IAdjudicatorEventEntryRepository interface {
	CreateEventEntry(entry *AdjudicatorEventEntry) error
	DeleteEventEntry(entry AdjudicatorEventEntry) error
	SearchEventEntry(criteria SearchAdjudicatorEventEntryCriteria) ([]AdjudicatorEventEntry, error)
	UpdateEventEntry(entry AdjudicatorEventEntry) error
}

// EventEntryList contains the event, and the athletes and couples who are competing in this event
type EventEntryList struct {
	Event          Event
	AthleteEntries []AthleteEventEntry
	CoupleEntries  []PartnershipEventEntry
}

// AdjudicatorEventEntryList contains the ID of an event and the Adjudicators that are assigned to this event
type AdjudicatorEventEntryList struct {
	EventID   int
	EntryList []AdjudicatorEventEntry
}
