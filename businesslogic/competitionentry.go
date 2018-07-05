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
	ID            int `schema:"id"`
	AthleteID     int `schema:"athlete"`
	CompetitionID int `schema:"competition"`
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

// SearchPartnershipCompetitionEntryCriteria specifies parameters that can be used to search the Competition Entry
// of a Partnership
type SearchPartnershipCompetitionEntryCriteria struct {
	Partnership int `schema:"partnership"`
	Competition int `schema:"competition"`
}

// IPartnershipCompetitionEntryRepository specifies functions that should be implemented to
// provide CRUD operations on PartnershipCompetitionEntry
type IPartnershipCompetitionEntryRepository interface {
	CreatePartnershipCompetitionEntry(entry *PartnershipCompetitionEntry) error
	DeletePartnershipCompetitionEntry(entry PartnershipCompetitionEntry) error
	SearchPartnershipCompetitionEntry(criteria SearchPartnershipCompetitionEntryCriteria) ([]PartnershipCompetitionEntry, error)
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
func (entry *AthleteCompetitionEntry) createAthleteCompetitionEntry(competitionRepo ICompetitionRepository, athleteCompEntryRepo IAthleteCompetitionEntryRepository) error {
	// check if competition still accept entries
	compSearchResults, searchCompErr := competitionRepo.SearchCompetition(
		SearchCompetitionCriteria{
			ID:       entry.CompetitionEntry.CompetitionID,
			StatusID: CompetitionStatusOpenRegistration,
		})
	if searchCompErr != nil {
		return searchCompErr
	}
	if len(compSearchResults) != 1 {
		return errors.New("competition does not exist or it no longer accept new entries")
	}

	criteria := SearchAthleteCompetitionEntryCriteria{
		AthleteID:     entry.AthleteID,
		CompetitionID: entry.CompetitionEntry.CompetitionID,
	}

	searchResults, err := athleteCompEntryRepo.SearchAthleteCompetitionEntry(criteria)
	if err != nil {
		return err
	}

	if len(searchResults) == 0 {
		return athleteCompEntryRepo.CreateAthleteCompetitionEntry(entry)
	}

	if len(searchResults) > 0 {
		return errors.New(fmt.Sprintf("competition entry for athlete %d is already created", entry.AthleteID))
	}

	return errors.New("cannot create competition entry for this athlete")
}

func (entry *PartnershipCompetitionEntry) createPartnershipCompetitionEntry(compRepo ICompetitionRepository, entryRepo IPartnershipCompetitionEntryRepository) error {
	// check if competition still accepts new entries
	competition, findCompErr := GetCompetitionByID(entry.CompetitionEntry.CompetitionID, compRepo)
	if findCompErr != nil {
		return findCompErr
	}
	if competition.GetStatus() != CompetitionStatusOpenRegistration {
		return errors.New("this competition no longer accepts new entries")
	}

	return nil
}
