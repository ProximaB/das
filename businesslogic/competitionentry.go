package businesslogic

import (
	"errors"
	"fmt"
	"time"
)

// CompetitionEntry is entry for a competition (not events).
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

func (entry *CompetitionEntry) CreateCompetitionEntry(entryRepo ICompetitionEntryRepository) error {
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
