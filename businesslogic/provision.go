// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
	historyRepository.CreateOrganizerProvisionHistory(&history)
	organizerRepository.UpdateOrganizerProvision(provision)
}
