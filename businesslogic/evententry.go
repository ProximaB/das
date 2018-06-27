// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// EventEntry is event-wise. It indicates the participation of partnership in a competitive ballroom event
// The owner of the
type EventEntry struct {
	ID              int
	EventID         int
	PartnershipID   int
	CheckInTime     time.Time
	leadAge         int
	followAge       int
	CompetitorTag   int // TODO: think about it. this does not support tags like A/B/C/D/E for team matches
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type PartnershipEventEntry struct {
	ID            int
	EventEntry    EventEntry
	PartnershipID int
	CheckInTime   time.Time
}

type SearchPartnershipEventEntryCriteria struct {
	PartnershipID int
	EventID       int
}

type IPartnershipEventEntryRepository interface {
	CreatePartnershipEventEntry(entry *PartnershipEventEntry) error
	DeletePartnershipEventEntry(entry PartnershipEventEntry) error
	SearchPartnershipEventEntry(criteria SearchPartnershipEventEntryCriteria) ([]PartnershipEventEntry, error)
	UpdatePartnershipEventEntry(entry PartnershipEventEntry) error
}

type AdjudicatorEventEntry struct {
	ID            int
	EventEntry    EventEntry
	AdjudicatorID int
}

type SearchAdjudicatorEventEntryCriteria struct {
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

func createEventEntry(entry PartnershipEventEntry, entryRepo IPartnershipEventEntryRepository) error {
	// check if entries were already created
	searchCriteria := SearchPartnershipEventEntryCriteria{
		///PartnershipID:              entry.PartnershipID,
		//CompetitiveBallroomEventID: entry.CompetitiveBallroomEventID,
	}
	existingEntries, _ := entryRepo.SearchPartnershipEventEntry(searchCriteria)

	if len(existingEntries) == 1 {
		return errors.New("event is already added")
	} else if len(existingEntries) > 1 {
		log.Println(errors.New(fmt.Sprintf("more than 1 entry has been added: %v", entry)))
	}

	return entryRepo.CreatePartnershipEventEntry(&entry)
}
