package businesslogic

import (
	"errors"
	"time"
)

type AthleteEventEntry struct {
	ID                int
	Athlete           Account
	Competition       Competition
	Event             Event
	CompetitorTag     int
	CheckedIn         bool
	DateTimeCheckedIn bool
	Placement         int // default should be 0: unplaced
	CreateUserID      int
	DateTimeCreated   time.Time
	UpdateUserID      int
	DateTimeUpdated   time.Time
}

type SearchAthleteEventEntryCriteria struct {
	ID            int
	AthleteID     int
	CompetitionID int
	EventID       int
}

type IAthleteEventEntryRepository interface {
	CreateAthleteEventEntry(entry *AthleteEventEntry) error
	DeleteAthleteEventEntry(entry AthleteEventEntry) error
	SearchAthleteEventEntry(criteria SearchAthleteEventEntryCriteria) ([]AthleteEventEntry, error)
	UpdateAthleteEventEntry(entry AthleteEventEntry) error
}

// AthleteEventEntryService imposes requirements of DAS on AthleteEventEntry management.
type AthleteEventEntryService struct {
}

func (service AthleteEventEntryService) CreateAthleteEventEntry(currentUser Account, entry AthleteEventEntry) error {
	return errors.New("not implemented")
}

func (service AthleteEventEntryService) DeleteAthleteEventEntry(currentUser Account, entry AthleteEventEntry) error {
	return errors.New("not implemented")
}

func (service AthleteEventEntryService) SearchAthleteEventEntry(criteria SearchAthleteEventEntryCriteria) ([]AthleteEventEntry, error) {
	return make([]AthleteEventEntry, 0), errors.New("not implemented")
}

func (service AthleteEventEntryService) UpdateAthleteEventEntry(currentUser Account, entry AthleteEventEntry) error {
	return errors.New("not implemented")
}
