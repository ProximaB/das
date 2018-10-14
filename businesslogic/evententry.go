// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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

type IAdjudicatorEventEntryRepository interface {
	CreateEventEntry(entry *AdjudicatorEventEntry) error
	DeleteEventEntry(entry AdjudicatorEventEntry) error
	SearchEventEntry(criteria SearchAdjudicatorEventEntryCriteria) ([]AdjudicatorEventEntry, error)
	UpdateEventEntry(entry AdjudicatorEventEntry) error
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
	searchedResults, err := entryRepo.SearchPartnershipEventEntry(SearchPartnershipEventEntryCriteria{
		PartnershipID: entry.PartnershipID,
		EventID:       entry.EventEntry.EventID,
	})
	if err != nil {
		return err
	}
	if len(searchedResults) > 0 {
		return errors.New(fmt.Sprintf("entry for partnership %d already exists for event %d", entry.PartnershipID, entry.EventEntry.EventID))
	}

	// entry does not exist, create the entry
	return entryRepo.CreatePartnershipEventEntry(&entry)
}
