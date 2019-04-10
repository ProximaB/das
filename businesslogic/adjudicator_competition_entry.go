package businesslogic

import "time"

// AdjudicatorCompetitionEntry defines the presence of an Adjudicator at a Competition
type AdjudicatorCompetitionEntry struct {
	ID              int
	AdjudicatorID   int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

// SearchAdjudicatorCompetitionEntryCriteria specifies the parameters that can be used to search Adjudicator's
// participation at competitions
type SearchAdjudicatorCompetitionEntryCriteria struct {
	ID            int
	AdjudicatorID int
	CompetitionID int
}

// IAdjudicatorCompetitionEntryRepository specifies the methods that should be
// implemented to provide repository function for businesslogic
type IAdjudicatorCompetitionEntryRepository interface {
	CreateEntry(entry *AdjudicatorCompetitionEntry) error
	DeleteEntry(entry AdjudicatorCompetitionEntry) error
	SearchEntry(criteria SearchAdjudicatorCompetitionEntryCriteria) ([]AdjudicatorCompetitionEntry, error)
	UpdateEntry(entry AdjudicatorCompetitionEntry) error
}
