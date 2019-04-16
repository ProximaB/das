package businesslogic

import (
	"errors"
	"fmt"
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

type AthleteEventEntry struct {
	ID                int
	Athlete           Account
	Competition       Competition
	Event             Event
	CompetitorTag     int
	CheckedIn         bool
	DateTimeCheckedIn bool
	Placement         int // default should be 0: unplaced
	CreateUserID      int
	DateTimeCreated   time.Time
	UpdateUserID      int
	DateTimeUpdated   time.Time
}

type SearchAthleteEventEntryCriteria struct {
	ID            int
	AthleteID     int
	CompetitionID int
	EventID       int
}

type IAthleteEventEntryRepository interface {
	CreateAthleteEventEntry(entry *AthleteEventEntry) error
	DeleteAthleteEventEntry(entry AthleteEventEntry) error
	SearchAthleteEventEntry(criteria SearchAthleteEventEntryCriteria) ([]AthleteEventEntry, error)
	UpdateAthleteEventEntry(entry AthleteEventEntry) error
}

// AthleteEventEntryService imposes requirements of DAS on AthleteEventEntry management.
type AthleteEventEntryService struct {
}

func (service AthleteEventEntryService) CreateAthleteEventEntry(currentUser Account, entry AthleteEventEntry) error {
	return errors.New("not implemented")
}

func (service AthleteEventEntryService) DeleteAthleteEventEntry(currentUser Account, entry AthleteEventEntry) error {
	return errors.New("not implemented")
}

func (service AthleteEventEntryService) SearchAthleteEventEntry(criteria SearchAthleteEventEntryCriteria) ([]AthleteEventEntry, error) {
	return make([]AthleteEventEntry, 0), errors.New("not implemented")
}

func (service AthleteEventEntryService) UpdateAthleteEventEntry(currentUser Account, entry AthleteEventEntry) error {
	return errors.New("not implemented")
}

// PartnershipEventEntry defines the Entry of a PartnershipID at an Event
type PartnershipEventEntry struct {
	ID                int
	Couple            Partnership
	Competition       Competition
	Event             Event
	CompetitorTag     int
	CheckedIn         bool
	DateTimeCheckedIn *time.Time
	Placement         int
	CreateUserID      int
	DateTimeCreated   time.Time
	UpdateUserID      int
	DateTimeUpdated   time.Time
}

// SearchPartnershipEventEntryCriteria specifies the parameters that can be used to search the Event Entry of a PartnershipID
type SearchPartnershipEventEntryCriteria struct {
	CompetitionID int
	PartnershipID int
	EventID       int
}

// IPartnershipEventEntryRepository defines the functions that need to be implemented to perform CRUD function
// for businesslogic to use
type IPartnershipEventEntryRepository interface {
	CreatePartnershipEventEntry(entry *PartnershipEventEntry) error
	DeletePartnershipEventEntry(entry PartnershipEventEntry) error
	SearchPartnershipEventEntry(criteria SearchPartnershipEventEntryCriteria) ([]PartnershipEventEntry, error)
	UpdatePartnershipEventEntry(entry PartnershipEventEntry) error
}

type PartnershipEventEntryService struct{}

// CreatePartnershipEventEntry checks if an entry for the specified PartnershipID already exists in the specified Event. If
// not, a new PartnershipEventEntry will be created for the specified event in the provided repository
func CreatePartnershipEventEntry(entry PartnershipEventEntry, entryRepo IPartnershipEventEntryRepository) error {
	// check if entries were already created
	searchedResults, err := entryRepo.SearchPartnershipEventEntry(SearchPartnershipEventEntryCriteria{
		PartnershipID: entry.Couple.ID,
		EventID:       entry.Event.ID,
	})
	if err != nil {
		return err
	}
	if len(searchedResults) > 0 {
		return errors.New(fmt.Sprintf("entry for partnership %d already exists for event %d", entry.Couple.ID, entry.Event.ID))
	}

	// entry does not exist, create the entry
	return entryRepo.CreatePartnershipEventEntry(&entry)
}
