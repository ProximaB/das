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
	Athlete                  Account
	Competition              Competition
	IsLead                   bool
	LeadTag                  int
	CheckedIn                bool
	DateTimeCheckedIn        time.Time
	OrganizerNote            string // Organizer may need to add notes to the entry
	PaymentReceivedIndicator bool
	DateTimeOfPayment        time.Time
	CreateUserID             int
	DateTimeCreated          time.Time
	UpdateUserID             int
	DateTimeUpdated          time.Time
}

// SearchAthleteCompetitionEntryCriteria specifies the parameters that can be used
// to search Athlete Competition Entries in DAS
type SearchAthleteCompetitionEntryCriteria struct {
	ID            int  `schema:"id"`
	AthleteID     int  `schema:"athlete"`
	CompetitionID int  `schema:"competition"`
	IsLead        bool `schema:"isLead"`
	Tag           int  `schema:"leadTag"`
}

// IAthleteCompetitionEntryRepository specifies the interface that data source should implement
// to perform CRUD operations for AthleteCompetitionEntry
type IAthleteCompetitionEntryRepository interface {
	CreateEntry(entry *AthleteCompetitionEntry) error
	DeleteEntry(entry AthleteCompetitionEntry) error
	SearchEntry(criteria SearchAthleteCompetitionEntryCriteria) ([]AthleteCompetitionEntry, error)
	UpdateEntry(entry AthleteCompetitionEntry) error
	NextAvailableLeadTag(competition Competition) (int, error)
	GetEntriesByCompetition(competitionId int) ([]AthleteCompetitionEntry, error)
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
			ID:       entry.Competition.ID,
			StatusID: CompetitionStatusOpenRegistration,
		})
	if searchCompErr != nil {
		return searchCompErr
	}
	if len(compSearchResults) != 1 {
		return errors.New("competition does not exist or it no longer accept new entries")
	}

	criteria := SearchAthleteCompetitionEntryCriteria{
		AthleteID:     entry.Athlete.ID,
		CompetitionID: entry.Competition.ID,
	}

	searchResults, err := service.athleteCompEntryRepo.SearchEntry(criteria)
	if err != nil {
		return err
	}

	if len(searchResults) == 0 {
		return service.athleteCompEntryRepo.CreateEntry(entry)
	}

	if len(searchResults) > 0 {
		return errors.New(fmt.Sprintf("competition entry for athlete %v is already created", entry.Athlete.ID))
	}

	return errors.New("cannot create competition entry for this athlete")
}

func (service AthleteCompetitionEntryService) DeleteAthleteCompetitionEntry(entry AthleteCompetitionEntry) error {
	return errors.New("not implemented")
}

func (service AthleteCompetitionEntryService) SearchAthleteCompetitionEntry(currentUser Account, criteria SearchAthleteCompetitionEntryCriteria) ([]AthleteCompetitionEntry, error) {
	return make([]AthleteCompetitionEntry, 0), errors.New("not implemented")
}

type CompetitionEntryList struct {
	Competition    Competition
	AthleteEntries []AthleteCompetitionEntry
	CoupleEntries  []PartnershipCompetitionEntry
}

// CompetitionLeadTag maps a competition with a lead and that lead's number tag
type CompetitionLeadTag struct {
	ID              int
	CompetitionID   int
	LeadID          int
	Tag             int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type CompetitionLeadTagCollection []CompetitionLeadTag

func (tags CompetitionLeadTagCollection) Len() int {
	return len(tags)
}

func (tags CompetitionLeadTagCollection) Less(i, j int) bool {
	if tags[i].Tag > tags[j].Tag {
		return true
	}
	return false
}

func (tags CompetitionLeadTagCollection) Swap(i, j int) {
	tags[i], tags[j] = tags[j], tags[i]
}

// SearchCompetitionLeadTagCriteria defines the parameters that can be used to search lead's tags at competitions
type SearchCompetitionLeadTagCriteria struct {
	ID            int
	CompetitionID int
	LeadID        int
	Tag           int
	CreateUserID  int
}

// ICompetitionLeadTagRepository defines the interface that a lead tag repository should implement.
type ICompetitionLeadTagRepository interface {
	CreateCompetitionLeadTag(tag *CompetitionLeadTag) error
	DeleteCompetitionLeadTag(tag CompetitionLeadTag) error
	SearchCompetitionLeadTag(criteria SearchCompetitionLeadTagCriteria) ([]CompetitionLeadTag, error)
	UpdateCompetitionLeadTag(tag CompetitionLeadTag) error
}

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
	athleteCompEntryRepo IAthleteCompetitionEntryRepository
	partnershipEntryRepo IPartnershipCompetitionEntryRepository
}

func NewPartnershipCompetitionEntryService(athleteEntryRepo IAthleteCompetitionEntryRepository, partnershipEntryRepo IPartnershipCompetitionEntryRepository) PartnershipCompetitionEntryService {
	return PartnershipCompetitionEntryService{
		athleteCompEntryRepo: athleteEntryRepo,
		partnershipEntryRepo: partnershipEntryRepo,
	}
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

// GetAllLeadEntries returns all the unique leads at the specified competition, if any
func (service PartnershipCompetitionEntryService) GetAllLeadEntries(competition Competition) ([]AthleteCompetitionEntry, error) {
	partnerships, err := service.partnershipEntryRepo.SearchEntry(SearchPartnershipCompetitionEntryCriteria{CompetitionID: competition.ID})
	uniqueLeads := make(map[int]Account)

	for _, each := range partnerships {
		eachLead := each.Couple.Lead
		if _, hasLead := uniqueLeads[eachLead.ID]; !hasLead {
			uniqueLeads[eachLead.ID] = eachLead
		}
	}

	// extract the unique leads from the map
	leadEntries := make([]AthleteCompetitionEntry, 0)
	for _, each := range uniqueLeads {
		entryResults, searchErr := service.athleteCompEntryRepo.SearchEntry(SearchAthleteCompetitionEntryCriteria{CompetitionID: competition.ID, AthleteID: each.ID})
		if searchErr != nil {
			return leadEntries, searchErr
		}
		if len(entryResults) != 1 {
			return leadEntries, errors.New(fmt.Sprintf("cannot find %v AthleteCompetitionEntry at Competitition %v ", each.FullName(), competition.Name))
		}
		leadEntries = append(leadEntries, entryResults[0])
	}

	return leadEntries, err
}

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
