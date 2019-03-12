package businesslogic

import (
	"errors"
	"fmt"
	"time"
)

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
