package businesslogic

import (
	"errors"
	"fmt"
	"time"
)

// CompetitionEntry is the entry for a competition (not including events).
// Athlete does not have to have a partner to enter a competition (depending on the rule)
// CompetitionEntry helps with
// - finding attendance of competition
// - reducing duplicate entries
type CompetitionEntry struct {
	ID                 int
	CompetitionID      int
	AthleteID          int  // account id
	CheckedIn          bool // only organizer can check in athlete
	PaymentReceivedIND bool
	PaymentDateTime    time.Time
	CheckInDateTime    *time.Time
	CreateUserID       int
	DateTimeCreated    time.Time
	UpdateUserID       int
	DateTimeUpdated    time.Time
}

// SearchCompetitionEntryCriteria provides parameters to ICompetitionEntryRepository to search competition entry
type SearchCompetitionEntryCriteria struct {
	ID            int
	CompetitionID int
	AthleteID     int
}

// ICompetitionEntryRepository specifies the interface that data source should implement
// to perform CRUD operations on CompetitionEntry
type ICompetitionEntryRepository interface {
	CreateCompetitionEntry(entry *CompetitionEntry) error
	UpdateCompetitionEntry(entry CompetitionEntry) error
	DeleteCompetitionEntry(entry CompetitionEntry) error
	SearchCompetitionEntry(criteria SearchCompetitionEntryCriteria) ([]CompetitionEntry, error)
}

// CompetitionTBAEntry provides the entry for dancers who do not have a partner
// but still would like to compete. Athlete who enters competition as TBA
// will also enter the match-making queue and DAS shall provides a list of dancers
// who satisfy the searching criteria the TBA dancer
type CompetitionTBAEntry struct {
	ID              int
	AccountID       int
	CompetitionID   int
	ContactEmail    string // optional
	ContactPhone    string // optional
	Message         string // use this message to specify level and style of dance to enter
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

// CreateCompetitionEntry will check if current entry exists in the repository. If yes, an error will be returned,
// if not, a competition entry will be created for this athlete.
// Competition must be during open registration stage.
func (entry *CompetitionEntry) CreateCompetitionEntry(competitionRepo ICompetitionRepository, entryRepo ICompetitionEntryRepository) error {

	// check if competition still accept entries
	compSearchResults, searchCompErr := competitionRepo.SearchCompetition(SearchCompetitionCriteria{ID: entry.CompetitionID, StatusID: COMPETITION_STATUS_OPEN_REGISTRATION})
	if searchCompErr != nil {
		return searchCompErr
	}
	if len(compSearchResults) != 1 {
		return errors.New("competition does not exist or it no longer accept new entries")
	}

	criteria := SearchCompetitionEntryCriteria{
		AthleteID:     entry.AthleteID,
		CompetitionID: entry.CompetitionID,
	}

	searchResults, err := entryRepo.SearchCompetitionEntry(criteria)
	if err != nil {
		return err
	}

	if len(searchResults) == 0 {
		return entryRepo.CreateCompetitionEntry(entry)
	}

	if len(searchResults) > 0 {
		return errors.New(fmt.Sprintf("competition entry for athlete %d is already created", entry.AthleteID))
	}

	return errors.New("cannot create competition entry for this athlete")
}
