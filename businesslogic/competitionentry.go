// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import (
	"errors"
	"fmt"
	"time"
)

// CompetitionEntry is the entry for a competition (not including events) and is a
// base entry for more specific entry such as AthleteCompetitionEntry, PartnershipCompetitionEntry,
// and AdjudicatorCompetitionEntry
type CompetitionEntry struct {
	CompetitionID    int
	CheckInIndicator bool
	DateTimeCheckIn  time.Time
	CreateUserID     int
	DateTimeCreated  time.Time
	UpdateUserID     int
	DateTimeUpdated  time.Time
}

// ICompetitionEntryRepository specifies the interface that data source should implement
// to perform CRUD operations on CompetitionEntry
type ICompetitionEntryRepository interface {
	CreateCompetitionEntry(entry *CompetitionEntry) error
	UpdateCompetitionEntry(entry CompetitionEntry) error
	DeleteCompetitionEntry(entry CompetitionEntry) error
	SearchCompetitionEntry(criteria SearchCompetitionEntryCriteria) ([]CompetitionEntry, error)
}

// SearchCompetitionEntryCriteria provides parameters to ICompetitionEntryRepository to search competition entry
type SearchCompetitionEntryCriteria struct {
	ID            int
	CompetitionID int
	AthleteID     int
}

// AthleteCompetitionEntry wraps CompetitionEntry and adds additional data to manage payment status for Athletes. It
// also allows quick indexing of competition attendance
type AthleteCompetitionEntry struct {
	ID                       int
	CompetitionEntry         CompetitionEntry
	AthleteID                int
	PaymentReceivedIndicator bool
	DateTimeOfPayment        time.Time
}

// SearchAthleteCompetitionEntryCriteria specifies the parameters that can be used
// to search Athlete Competition Entries in DAS
type SearchAthleteCompetitionEntryCriteria struct {
	ID int `schema:"id"`
}

// IAthleteCompetitionEntryRepository specifies the interface that data source should implement
// to perform CRUD operations for AthleteCompetitionEntry
type IAthleteCompetitionEntryRepository interface {
	CreateAthleteCompetitionEntry(entry *AthleteCompetitionEntry) error
	DeleteAthleteCompetitionEntry(entry AthleteCompetitionEntry) error
	SearchAthleteCompetitionEntry(criteria SearchAthleteCompetitionEntryCriteria) ([]AthleteCompetitionEntry, error)
	UpdateAthleteCompetitionEntry(entry AthleteCompetitionEntry) error
}

// PartnershipCompetitionEntry defines a partnership's participation of a competition
type PartnershipCompetitionEntry struct {
	ID               int
	CompetitionEntry CompetitionEntry
	PartnershipID    int
}

// SearchPartnershipCompetitionEntryCriteria specifies
type SearchPartnershipCompetitionEntryCriteria struct {
}

// IPartnershipCompetitionEntryRepository specifies functions that should be implemented to
// provide CRUD operations on PartnershipCompetitionEntry
type IPartnershipCompetitionEntryRepository interface {
	CreatePartnershipCompetitionEntry(entry *PartnershipCompetitionEntry) error
	DeletePartnershipCompetitionEntry(entry PartnershipCompetitionEntry) error
	SearchPartnershipCompetitionEntry(criteria SearchPartnershipCompetitionEntryCriteria) error
	UpdatePartnershipCompetitionEntry(entry PartnershipCompetitionEntry) error
}

// AdjudicatorCompetitionEntry defines the presence of an Adjudicator at a Competition
type AdjudicatorCompetitionEntry struct {
	ID               int
	CompetitionEntry CompetitionEntry
	AdjudicatorID    int
}

// SearchAdjudicatorCompetitionEntryCriteria specifies the parameters that can be used to search Adjudicator's
// participation at competitions
type SearchAdjudicatorCompetitionEntryCriteria struct {
}

// IAdjudicatorCompetitionEntryRepository specifies the methods that should be
// implemented to provide repository function for businesslogic
type IAdjudicatorCompetitionEntryRepository interface {
	CreateAdjudicatorCompetitionEntry(entry *AdjudicatorCompetitionEntry) error
	DeleteAdjudicatorCompetitionEntry(entry AdjudicatorCompetitionEntry) error
	SearchAdjudicatorCompetitionEntry(criteria SearchAdjudicatorCompetitionEntryCriteria) ([]AdjudicatorCompetitionEntry, error)
	UpdateAdjudicatorCompetitionEntry(entry AdjudicatorCompetitionEntry) error
}

// AthleteCompetitionTBAEntry provides the entry for dancers who do not have a partner
// but still would like to compete. Athlete who enters competition as TBA
// will also enter the match-making queue and DAS shall provides a list of dancers
// who satisfy the searching criteria the TBA dancer
type AthleteCompetitionTBAEntry struct {
	ID              int
	AccountID       int
	CompetitionID   int
	ContactEmail    string // optional
	ContactPhone    string // optional
	Message         string // use this message to specify level and style of dance to enter
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

// CreateAthleteCompetitionEntry will check if current entry exists in the repository. If yes, an error will be returned,
// if not, a competition entry will be created for this athlete.
// Competition must be during open registration stage.
func (entry * AthleteCompetitionEntry) CreateAthleteCompetitionEntry(competitionRepo ICompetitionRepository, entryRepo IAthleteCompetitionEntryRepository) error {

	// check if competition still accept entries
	compSearchResults, searchCompErr := competitionRepo.SearchCompetition(
		SearchCompetitionCriteria {
			ID: entry.CompetitionEntry.CompetitionID, 
			StatusID: COMPETITION_STATUS_OPEN_REGISTRATION,
		})
	if searchCompErr != nil {
		return searchCompErr
	}
	if len(compSearchResults) != 1 {
		return errors.New("competition does not exist or it no longer accept new entries")
	}

	criteria := SearchAthleteCompetitionEntryCriteria{
		//AthleteID:     entry.AthleteID,
		//CompetitionID: entry.CompetitionID,
	}

	searchResults, err := entryRepo.SearchAthleteCompetitionEntry(criteria)
	if err != nil {
		return err
	}

	if len(searchResults) == 0 {
		return entryRepo.CreateAthleteCompetitionEntry(entry)
	}

	if len(searchResults) > 0 {
		return errors.New(fmt.Sprintf("competition entry for athlete %d is already created", entry.AthleteID))
	}

	return errors.New("cannot create competition entry for this athlete")
}
