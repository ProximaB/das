// Copyright 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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
	Mask            int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

// PartnershipEventEntry defines the Entry of a Partnership at an Event
type PartnershipEventEntry struct {
	ID            int
	EventEntry    EventEntry
	PartnershipID int
	leadAge       int
	followAge     int
	CheckInTime   time.Time
}

// SearchPartnershipEventEntryCriteria specifies the parameters that can be used to search the Event Entry of a Partnership
type SearchPartnershipEventEntryCriteria struct {
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

// PartnershipEventEntryList contains the ID of an event and the Partnerships that are competing in this event
type PartnershipEventEntryList struct {
	EventID   int
	EntryList []PartnershipEventEntry
}

// AdjudicatorEventEntryList contains the ID of an event and the Adjudicators that are assigned to this event
type AdjudicatorEventEntryList struct {
	EventID   int
	EntryList []AdjudicatorEventEntry
}

// CreatePartnershipEventEntry checks if an entry for the specified Partnership already exists in the specified Event. If
// not, a new PartnershipEventEntry will be created for the specified event in the provided repository
func CreatePartnershipEventEntry(entry PartnershipEventEntry, entryRepo IPartnershipEventEntryRepository) error {
	// check if entries were already created
	if searchedResults, err := entryRepo.SearchPartnershipEventEntry(SearchPartnershipEventEntryCriteria{
		PartnershipID: entry.PartnershipID,
		EventID:       entry.EventEntry.EventID,
	}); err != nil {
		return err
	} else if len(searchedResults) > 0 {
		return errors.New(fmt.Sprintf("entry for partnership %d already exists for event %d", entry.PartnershipID, entry.EventEntry.EventID))
	}

	// entry does not exist, create the entry
	return entryRepo.CreatePartnershipEventEntry(&entry)
}
