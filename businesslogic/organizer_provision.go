package businesslogic

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// OrganizerProvision provision organizer competition slots for creating and hosting competitions
type OrganizerProvision struct {
	ID              int
	AccountID       int
	OrganizerRoleID int
	Organizer       Account
	Available       int
	Hosted          int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type UpdateOrganizerProvision struct {
	OrganizerID   int
	Amount        int
	Note          string
	CurrentUserID int
}

// SearchOrganizerProvisionCriteria specifies the search criteria of Organizer's provision information
type SearchOrganizerProvisionCriteria struct {
	ID           int    `schema:"id"`
	OrganizerID  int    `schema:"organizerID"`  // organizer's account ID, not type-account id
	OrganizerUID string `schema:"organizerUID"` // Organizer's UID,
}

// IOrganizerProvisionRepository specifies the interface that a repository should implement for Organizer Provision
type IOrganizerProvisionRepository interface {
	CreateOrganizerProvision(provision *OrganizerProvision) error
	UpdateOrganizerProvision(provision OrganizerProvision) error
	DeleteOrganizerProvision(provision OrganizerProvision) error
	SearchOrganizerProvision(criteria SearchOrganizerProvisionCriteria) ([]OrganizerProvision, error)
}

func (provision OrganizerProvision) updateForCreateCompetition(competition Competition) OrganizerProvision {
	newProvision := provision
	newProvision.Available = provision.Available - 1
	newProvision.Hosted = provision.Hosted + 1
	newProvision.UpdateUserID = competition.CreateUserID
	newProvision.DateTimeUpdated = time.Now()
	return newProvision
}

func initializeOrganizerProvision(accountRoleID, currentUserID int) (OrganizerProvision, OrganizerProvisionHistoryEntry) {
	provision := OrganizerProvision{
		OrganizerRoleID: accountRoleID,
		Available:       0,
		CreateUserID:    currentUserID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    currentUserID,
		DateTimeUpdated: time.Now(),
	}
	history := OrganizerProvisionHistoryEntry{
		OrganizerRoleID: accountRoleID,
		Amount:          0,
		Note:            "initialize organizer organizer",
		CreateUserID:    currentUserID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    currentUserID,
		DateTimeUpdated: time.Now(),
	}
	return provision, history
}

func updateOrganizerProvision(provision OrganizerProvision, history OrganizerProvisionHistoryEntry,
	organizerRepository IOrganizerProvisionRepository, historyRepository IOrganizerProvisionHistoryRepository) {
	historyRepository.CreateOrganizerProvisionHistory(&history)
	organizerRepository.UpdateOrganizerProvision(provision)
}

// OrganizerProvisionServices provides functions that allows provisioning Organizer's Competition, including updating
// and querying Organizer's Competition Provision.
type OrganizerProvisionService struct {
	accountRepo                   IAccountRepository
	accountRoleRepo               IAccountRoleRepository
	organizerProvisionRepo        IOrganizerProvisionRepository
	organizerProvisionHistoryRepo IOrganizerProvisionHistoryRepository
}

func NewOrganizerProvisionService(
	accountRepo IAccountRepository,
	accountRoleRepo IAccountRoleRepository,
	organizerProvisionRepo IOrganizerProvisionRepository,
	organizerProvisionHistoryRepo IOrganizerProvisionHistoryRepository) OrganizerProvisionService {
	return OrganizerProvisionService{
		accountRepo:                   accountRepo,
		accountRoleRepo:               accountRoleRepo,
		organizerProvisionRepo:        organizerProvisionRepo,
		organizerProvisionHistoryRepo: organizerProvisionHistoryRepo,
	}
}

func (service OrganizerProvisionService) NewOrganizerProvision(accountRoleID, currentUserID int) (OrganizerProvision, error) {
	if service.accountRoleRepo == nil {
		return OrganizerProvision{}, errors.New("AccountRoleRepository is not initialized")
	}
	results, searchErr := service.accountRoleRepo.SearchAccountRole(SearchAccountRoleCriteria{
		ID: accountRoleID,
	})
	if searchErr != nil {
		log.Printf("error in searching account role: %v", searchErr)
		return OrganizerProvision{}, searchErr
	}
	if len(results) != 1 {
		log.Printf("cannot find account role with ID %v", accountRoleID)
		return OrganizerProvision{}, errors.New(fmt.Sprintf("Cannot find account role with ID %d", accountRoleID))
	}
	organizerRole := results[0]
	entry := OrganizerProvision{
		AccountID:       organizerRole.AccountID,
		OrganizerRoleID: organizerRole.ID,
		Available:       0,
		Hosted:          0,
		CreateUserID:    currentUserID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    currentUserID,
		DateTimeUpdated: time.Now(),
	}
	return entry, nil
}

// NewInitialOrganizerProvisionHistoryEntry initialize the first provision history for an Organizer.
func (service OrganizerProvisionService) NewInitialOrganizerProvisionHistoryEntry(accountRoleID, currentUserID int) OrganizerProvisionHistoryEntry {
	return OrganizerProvisionHistoryEntry{
		OrganizerRoleID: accountRoleID,
		Amount:          0,
		Note:            "initialize organizer organizer",
		CreateUserID:    currentUserID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    currentUserID,
		DateTimeUpdated: time.Now(),
	}
}

func (service OrganizerProvisionService) SearchOrganizerProvision(criteria SearchOrganizerProvisionCriteria) ([]OrganizerProvision, error) {
	return service.organizerProvisionRepo.SearchOrganizerProvision(criteria)
}

func (service OrganizerProvisionService) UpdateOrganizerCompetitionProvision(update UpdateOrganizerProvision) error {
	provisions, searchErr := service.organizerProvisionRepo.SearchOrganizerProvision(SearchOrganizerProvisionCriteria{OrganizerID: update.OrganizerID})
	if searchErr != nil {
		return searchErr
	}
	if len(provisions) != 1 {
		return errors.New("cannot find organizer's competition provision information")
	}

	provision := provisions[0]
	provision.Available = provision.Available + update.Amount
	provision.UpdateUserID = update.CurrentUserID
	provision.DateTimeUpdated = time.Now()

	history := OrganizerProvisionHistoryEntry{
		OrganizerRoleID: update.OrganizerID,
		Amount:          update.Amount,
		Note:            update.Note,
		CreateUserID:    update.CurrentUserID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    update.CurrentUserID,
		DateTimeUpdated: time.Now(),
	}

	updateErr := service.organizerProvisionRepo.UpdateOrganizerProvision(provision)
	if updateErr != nil {
		return updateErr
	}
	createErr := service.organizerProvisionHistoryRepo.CreateOrganizerProvisionHistory(&history)
	if createErr != nil {
		return createErr
	}
	return nil
}
