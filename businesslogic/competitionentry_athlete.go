package businesslogic

import (
	"errors"
	"fmt"
	"time"
)

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

// AthleteCompetitionEntryService encapsulates the data flow of Athlete's Competition Entry, including data validation
// and sanitization.
type AthleteCompetitionEntryService struct {
	accountRepo          IAccountRepository
	competitionRepo      ICompetitionRepository
	athleteCompEntryRepo IAthleteCompetitionEntryRepository
}

// NewAthleteCompetitionEntryService instantiates a new AthleteCompetitionEntryService.
func NewAthleteCompetitionEntryService(accountRepo IAccountRepository, competitionRepo ICompetitionRepository, athleteCompEntryRepo IAthleteCompetitionEntryRepository) AthleteCompetitionEntryService {
	return AthleteCompetitionEntryService{
		accountRepo:          accountRepo,
		competitionRepo:      competitionRepo,
		athleteCompEntryRepo: athleteCompEntryRepo,
	}
}

// CreateAthleteCompetitionEntry attempts to create competition for an athlete if following checks pass:
// - If the create user is authorized
//		- If the create user is the athlete: proceed
//		- If the create user is an organizer or scrutineer of this competition
// - If current entry exists in the repository:
// 		- yes, return error
//		- no: proceed
// - If Competition is in open registration stage:
//		- yes: proceed
//		- no: return error
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

func (service AthleteCompetitionEntryService) DeleteAthleteCompetitionEntry(entry AthleteCompetitionEntry) error {
	return errors.New("not implemented")
}

func (service AthleteCompetitionEntryService) SearchAthleteCompetitionEntry(currentUser Account, criteria SearchAthleteCompetitionEntryCriteria) ([]AthleteCompetitionEntry, error) {
	return make([]AthleteCompetitionEntry, 0), errors.New("not implemented")
}