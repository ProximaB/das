package businesslogic

import (
	"time"
)

type OrganizerProvision struct {
	ID              int
	OrganizerID     int
	Available       int
	Hosted          int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type SearchOrganizerProvisionCriteria struct {
	ID          int `schema:"organizer"`
	OrganizerID int `schema:"organizer"`
}

type IOrganizerProvisionRepository interface {
	CreateOrganizerProvision(provision OrganizerProvision) error
	UpdateOrganizerProvision(provision OrganizerProvision) error
	DeleteOrganizerProvision(provision OrganizerProvision) error
	SearchOrganizerProvision(criteria *SearchOrganizerProvisionCriteria) ([]OrganizerProvision, error)
}

func (provision OrganizerProvision) updateForCreateCompetition(competition Competition) OrganizerProvision {
	newProvision := provision
	newProvision.Available = provision.Available - 1
	newProvision.Hosted = provision.Hosted + 1
	newProvision.UpdateUserID = competition.CreateUserID
	newProvision.DateTimeUpdated = time.Now()
	return newProvision
}

func initializeOrganizerProvision(accountID int) (OrganizerProvision, OrganizerProvisionHistoryEntry) {
	provision := OrganizerProvision{
		OrganizerID:     accountID,
		Available:       0,
		CreateUserID:    accountID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    accountID,
		DateTimeUpdated: time.Now(),
	}
	history := OrganizerProvisionHistoryEntry{
		OrganizerID:     accountID,
		Amount:          0,
		Note:            "initialize organizer organizer",
		CreateUserID:    accountID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    accountID,
		DateTimeUpdated: time.Now(),
	}
	return provision, history
}

func updateOrganizerProvision(provision OrganizerProvision, history OrganizerProvisionHistoryEntry,
	organizerRepository IOrganizerProvisionRepository, historyRepository IOrganizerProvisionHistoryRepository) {
	historyRepository.CreateOrganizerProvisionHistory(history)
	organizerRepository.UpdateOrganizerProvision(provision)
}
