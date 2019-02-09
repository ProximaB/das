package businesslogic

import "time"

type OrganizerProvisionHistoryEntry struct {
	ID              int
	OrganizerRoleID int
	Amount          int
	Note            string
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type SearchOrganizerProvisionHistoryCriteria struct {
	OrganizerID int
}

func newProvisionHistoryEntry(provision OrganizerProvision, competition Competition) OrganizerProvisionHistoryEntry {
	historyEntry := OrganizerProvisionHistoryEntry{
		OrganizerRoleID: provision.OrganizerRoleID,
		Amount:          -1,
		Note:            "created competition " + competition.Name,
		CreateUserID:    competition.CreateUserID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    competition.CreateUserID,
		DateTimeUpdated: time.Now(),
	}
	return historyEntry
}

type IOrganizerProvisionHistoryRepository interface {
	SearchOrganizerProvisionHistory(criteria SearchOrganizerProvisionHistoryCriteria) ([]OrganizerProvisionHistoryEntry, error)
	UpdateOrganizerProvisionHistory(history OrganizerProvisionHistoryEntry) error
	DeleteOrganizerProvisionHistory(history OrganizerProvisionHistoryEntry) error
	CreateOrganizerProvisionHistory(history *OrganizerProvisionHistoryEntry) error
}
