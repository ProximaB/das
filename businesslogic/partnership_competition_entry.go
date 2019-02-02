package businesslogic

import (
	"errors"
)

// PartnershipCompetitionEntry defines a partnership's participation of a competition
type PartnershipCompetitionEntry struct {
	ID               int
	PartnershipID    int
	Partnership      Partnership
	CompetitionEntry BaseCompetitionEntry
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
	CreateEntry(entry *PartnershipCompetitionEntry) error
	DeleteEntry(entry PartnershipCompetitionEntry) error
	SearchEntry(criteria SearchPartnershipCompetitionEntryCriteria) ([]PartnershipCompetitionEntry, error)
	UpdateEntry(entry PartnershipCompetitionEntry) error
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
