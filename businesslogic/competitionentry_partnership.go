package businesslogic

import (
	"errors"
	"time"
)

// PartnershipCompetitionEntry defines a partnership's participation of a competition
type PartnershipCompetitionEntry struct {
	ID                int
	Couple            Partnership
	Competition       Competition
	CheckedIn         bool
	DateTimeCheckedIn time.Time
	CreateUserID      int
	DateTimeCreated   time.Time
	UpdateUserID      int
	DateTimeUpdated   time.Time
}

// SearchPartnershipCompetitionEntryCriteria specifies parameters that can be used to search the Competition Entry
// of a Partnership
type SearchPartnershipCompetitionEntryCriteria struct {
	ID            int `schema:"id"`
	PartnershipID int `schema:"partnership"`
	CompetitionID int `schema:"competition"`
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
	competition, findCompErr := GetCompetitionByID(entry.Competition.ID, compRepo)
	if findCompErr != nil {
		return findCompErr
	}
	if competition.GetStatus() != CompetitionStatusOpenRegistration {
		return errors.New("this competition no longer accepts new entries")
	}

	return nil
}

type PartnershipCompetitionEntryService struct {
}

func (service PartnershipCompetitionEntryService) CreatePartnershipCompetitionEntry() error {
	return errors.New("not implemented")
}

func (service PartnershipCompetitionEntryService) DeletePartnershipCompetitionEntry() error {
	return errors.New("not implemented")
}

func (service PartnershipCompetitionEntryService) SearchPartnershipCompetitionEntry() error {
	return errors.New("not implemented")
}
