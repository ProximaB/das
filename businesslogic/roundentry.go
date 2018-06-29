// Copyright 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import (
	"errors"
	"time"
)

// RoundEntry is the base entry for a round that can be used for adjudicator and partnership.
type RoundEntry struct {
	RoundID         int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

// PartnershipRoundEntry is the Round Entry of a Partnership
type PartnershipRoundEntry struct {
	ID            int
	RoundEntry    RoundEntry
	PartnershipID int
}

// SearchPartnershipRoundEntryCriteria specifies the parameters that can be used to search the Round Entry of Partnership
type SearchPartnershipRoundEntryCriteria struct {
	ID            int `schema:"entry"`
	RoundID       int `schema:"round"`
	PartnershipID int `schema:"partnership"`
	EventID       int `schema:"event"`
}

// IPartnershipRoundEntryRepository specifies the functions that need to be implemented to perform CRUD operations
type IPartnershipRoundEntryRepository interface {
	CreatePartnershipRoundEntry(entry *PartnershipRoundEntry) error
	DeletePartnershipRoundEntry(entry PartnershipRoundEntry) error
	SearchPartnershipRoundEntry(criteria SearchPartnershipRoundEntryCriteria) ([]PartnershipRoundEntry, error)
	UpdatePartnershipRoundEntry(entry PartnershipRoundEntry) error
}

// AdjudicatorRoundEntry defines the Round Entry of an Adjudicator
type AdjudicatorRoundEntry struct {
	ID                 int
	AdjudicatorEntryID int
	RoundEntry         RoundEntry
}

// SearchAdjudicatorRoundEntryCriteria specifies the parameters that can be used to search the Round Entry of Adjudicator
type SearchAdjudicatorRoundEntryCriteria struct {
}

// IAdjudicatorRoundEntryRepository specifies the functions that need to be implemented to perform CRUD operations
type IAdjudicatorRoundEntryRepository interface {
	CreateAdjudicatorRoundEntry(entry *AdjudicatorRoundEntry) error
	DeleteAdjudicatorRoundEntry(entry AdjudicatorRoundEntry) error
	SearchAdjudicatorRoundEntry(criteria SearchAdjudicatorRoundEntryCriteria) ([]AdjudicatorRoundEntry, error)
	UpdateAdjudicatorRoundEntry(entry AdjudicatorRoundEntry) error
}

// CreatePartnershipRoundEntry checks if provided entry already exists in the repository, and create a new entry if
// the specified partnership has no entry at the specified round.
func CreatePartnershipRoundEntry(entry *PartnershipRoundEntry, repo IPartnershipRoundEntryRepository) error {
	// check if there is a duplicate entry
	if searchedResults, err := repo.SearchPartnershipRoundEntry(SearchPartnershipRoundEntryCriteria{
		PartnershipID: entry.PartnershipID,
		RoundID:       entry.RoundEntry.RoundID,
	}); err != nil {
		return err
	} else if searchedResults != nil && len(searchedResults) > 1 {
		return errors.New("round entry for this partnership is created")
	}

	// create the round entry
	return repo.CreatePartnershipRoundEntry(entry)
}
