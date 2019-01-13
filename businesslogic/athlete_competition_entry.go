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

// BaseCompetitionEntry is the entry for a competition (not including events) and is a
// base entry for more specific entry such as AthleteCompetitionEntry, PartnershipCompetitionEntry,
// and AdjudicatorCompetitionEntry
type BaseCompetitionEntry struct {
	CompetitionID    int
	CheckInIndicator bool
	DateTimeCheckIn  time.Time
	CreateUserID     int
	DateTimeCreated  time.Time
	UpdateUserID     int
	DateTimeUpdated  time.Time
}

// AthleteCompetitionEntry wraps BaseCompetitionEntry and adds additional data to manage payment status for Athletes. It
// also allows quick indexing of competition attendance
type AthleteCompetitionEntry struct {
	ID                       int
	CompetitionEntry         BaseCompetitionEntry
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
	CreateEntry(entry *AthleteCompetitionEntry) error
	DeleteEntry(entry AthleteCompetitionEntry) error
	SearchEntry(criteria SearchAthleteCompetitionEntryCriteria) ([]AthleteCompetitionEntry, error)
	UpdateEntry(entry AthleteCompetitionEntry) error
}

type AthleteCompetitionEntryService struct {
	accountRepo          IAccountRepository
	competitionRepo      ICompetitionRepository
	athleteCompEntryRepo IAthleteCompetitionEntryRepository
}

func NewAthleteCompetitionEntryService(accountRepo IAccountRepository, competitionRepo ICompetitionRepository, athleteCompEntryRepo IAthleteCompetitionEntryRepository) AthleteCompetitionEntryService {
	return AthleteCompetitionEntryService{
		accountRepo,
		competitionRepo,
		athleteCompEntryRepo,
	}
}

// CreateAthleteCompetitionEntry checks if current entry exists in the repository. If yes, an error will be returned,
// if not, a competition entry will be created for this athlete.
// Competition must be during open registration stage.
func (service AthleteCompetitionEntryService) CreateAthleteCompetitionEntry(entry *AthleteCompetitionEntry) error {
	// check if competition still accept entries
	compSearchResults, searchCompErr := service.competitionRepo.SearchCompetition(
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

	searchResults, err := service.athleteCompEntryRepo.SearchEntry(criteria)
	if err != nil {
		return err
	}

	if len(searchResults) == 0 {
		return service.athleteCompEntryRepo.CreateEntry(entry)
	}

	if len(searchResults) > 0 {
		return errors.New(fmt.Sprintf("competition entry for athlete %d is already created", entry.AthleteID))
	}

	return errors.New("cannot create competition entry for this athlete")
}
