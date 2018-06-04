package businesslogic

import (
	"time"
)

// this entry is competition-wise. athlete does not need to have a partner to enter a competition
// competition entry helps with
// - attendance of competition
// - billing organizer for each unique dancer
type CompetitionEntry struct {
	ID                 int
	CompetitionID      int
	AthleteID          int  // account id
	CheckedIn          bool // only organzier can check in athlete
	PaymentReceivedIND bool
	PaymentDateTime    time.Time
	CheckInDateTime    *time.Time
	CreateUserID       int
	DateTimeCreated    time.Time
	UpdateUserID       int
	DateTimeUpdated    time.Time
}

type ICompetitionEntryRepository interface {
	CreateCompetitionEntry(entry CompetitionEntry) error
	UpdateCompetitionEntry(entry CompetitionEntry) error
	DeleteCompetitionEntry(entry CompetitionEntry) error
	SearchCompetitionEntry(criteria SearchCompetitionEntryCriteria) ([]CompetitionEntry, error)
}

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

func CreateCompetitionEntry(user *Account,
	registration *CompetitiveBallroomEventRegistration,
	repo ICompetitionEntryRepository,
	accountRepo IAccountRepository,
	partnershipRepo IPartnershipRepository) error {
	partnerships, _ := partnershipRepo.SearchPartnership(SearchPartnershipCriteria{PartnershipID: registration.PartnershipID})
	leadAccount := GetAccountByID(partnerships[0].LeadID, accountRepo)
	followAccount := GetAccountByID(partnerships[0].FollowID, accountRepo)

	// check if entry has been created
	leadEntry, _ := repo.SearchCompetitionEntry(SearchCompetitionEntryCriteria{
		CompetitionID: registration.CompetitionID,
		AthleteID:     leadAccount.ID,
	})

	followEntry, _ := repo.SearchCompetitionEntry(SearchCompetitionEntryCriteria{
		CompetitionID: registration.CompetitionID,
		AthleteID:     followAccount.ID,
	})

	if len(leadEntry) == 0 {
		// entry does not exist, create entry
		if createErr := repo.CreateCompetitionEntry(CompetitionEntry{
			CompetitionID:   registration.CompetitionID,
			AthleteID:       leadAccount.ID,
			CreateUserID:    user.ID,
			DateTimeCreated: time.Now(),
			UpdateUserID:    user.ID,
			DateTimeUpdated: time.Now(),
		}); createErr != nil {
			return createErr
		}
	}

	if len(followEntry) == 0 {
		if createErr := repo.CreateCompetitionEntry(CompetitionEntry{
			CompetitionID:   registration.CompetitionID,
			AthleteID:       followAccount.ID,
			CreateUserID:    user.ID,
			DateTimeCreated: time.Now(),
			UpdateUserID:    user.ID,
			DateTimeUpdated: time.Now(),
		}); createErr != nil {
			return createErr
		}
	}

	//updateCompetitionAttendance(registration.ID)
	return nil
}
