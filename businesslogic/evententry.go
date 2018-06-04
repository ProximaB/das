package businesslogic

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type IEventEntryRepository interface {
	SearchEventEntry(criteria SearchEventEntryCriteria) ([]EventEntry, error)
	CreateEventEntry(entry *EventEntry) error
	DeleteEventEntry(entry EventEntry) error
	UpdateEventEntry(entry EventEntry) error
}
type SearchCompetitionEntryCriteria struct {
	ID            int
	CompetitionID int
	AthleteID     int
}

// EventEntry is event-wise. It indicates the participation of partnership in a competitive ballroom event
// The owner of the
type EventEntry struct {
	ID                int
	EventID           int
	PartnershipID     int
	CheckInTime       time.Time
	leadAge           int
	followAge         int
	leadSkillRating   float64
	followSkillRating float64
	CompetitorTag     int // TODO: think about it. this does not support tags like A/B/C/D/E for team matches
	CreateUserID      int
	DateTimeCreated   time.Time
	UpdateUserID      int
	DateTimeUpdated   time.Time
}

type EventEntryPublicView struct {
	CompetitiveBallroomEventEntryID int
	CompetitiveBallroomEventID      int
	EventID                         int
	CompetitionID                   int
	PartnershipID                   int
	LeadID                          int
	LeadFirstName                   string
	LeadLastName                    string
	FollowID                        int
	FollowFirstName                 string
	FollowLastName                  string
	CountryRepresented              string
	StateRepresented                string
	SchoolRepresented               string
	StudioRepresented               string
}

type SearchEventEntryCriteria struct {
	CompetitionID int `schema:"competition"`
	EventID       int
	PartnershipID int
	Federation    int `schema:"federation"`
	Division      int `schema:"division"`
	Age           int `schema:"age"`
	Proficiency   int `schema:"proficiency"`
	Style         int `schema:"style"`
}

type EventEntryList struct {
	EventID   int
	EntryList []EventEntryPublicView
}

func createEventEntry(entry EventEntry, entryRepo IEventEntryRepository) error {
	// check if entries were already created
	searchCriteria := SearchEventEntryCriteria{
		///PartnershipID:              entry.PartnershipID,
		//CompetitiveBallroomEventID: entry.CompetitiveBallroomEventID,
	}
	existingEntries, _ := entryRepo.SearchEventEntry(searchCriteria)

	if len(existingEntries) == 1 {
		return errors.New("event is already added")
	} else if len(existingEntries) > 1 {
		log.Println(errors.New(fmt.Sprintf("more than 1 entry has been added: %v", entry)))
	}

	return entryRepo.CreateEventEntry(&entry)
}
