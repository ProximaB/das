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
	"errors"
	"time"
)

const (
	RoleApplicationStatusApproved = 1
	RoleApplicationStatusDenied   = 2
	RoleApplicationStatusPending  = 3
)

type SearchRoleApplicationCriteria struct {
	AccountID      int
	AppliedRoleID  int
	StatusID       int
	ApprovalUserID int
}

type RoleApplication struct {
	ID               int
	AccountID        int
	AppliedRoleID    int
	Description      string
	StatusID         int
	ApprovalUserID   *int
	DateTimeApproved time.Time
	CreateUserID     int
	DateTimeCreated  time.Time
	UpdateUserID     int
	DateTimeUpdated  time.Time
}

type IRoleApplicationRepository interface {
	CreateApplication(application *RoleApplication) error
	SearchApplication(criteria SearchRoleApplicationCriteria) ([]RoleApplication, error)
	UpdateApplication(application RoleApplication) error
}

type RoleProvisionService struct {
	accountRepo         IAccountRepository
	roleApplicationRepo IRoleApplicationRepository
}

func NewRoleProvisionService(accountRepo IAccountRepository, roleApplicationRepo IRoleApplicationRepository) *RoleProvisionService {
	service := RoleProvisionService{
		accountRepo:         accountRepo,
		roleApplicationRepo: roleApplicationRepo,
	}
	return &service
}

func (service RoleProvisionService) respondRoleApplication(currentUser Account, application *RoleApplication, action int) error {
	application.StatusID = action
	application.ApprovalUserID = &currentUser.ID
	application.DateTimeApproved = time.Now()
	return service.roleApplicationRepo.UpdateApplication(*application)
}

// UpdateApplication attempts to approve the Role application based on the privilege of current user.
// If current user is admin, any application can be approved
// If current user is organizer, only emcee and deck-captain can be approved
// If current user is other roles, current user will be prohibited from performing such action
func (service RoleProvisionService) UpdateApplication(currentUser Account, application *RoleApplication, action int) error {
	// check if action is valid
	if !(action == RoleApplicationStatusApproved || action == RoleApplicationStatusDenied) {
		return errors.New("invalid response to role application")
	}
	// Only an Admin or Organizer user ca update user's role application
	if !(currentUser.HasRole(AccountTypeOrganizer) || currentUser.HasRole(AccountTypeAdministrator)) {
		return errors.New("unauthorized")
	}
	// should not allow users to provision themselves other than Admin
	if currentUser.ID == application.AccountID && !currentUser.HasRole(AccountTypeAdministrator) {
		return errors.New("not authorized to provision your own role application")
	}
	switch application.AppliedRoleID {
	case AccountTypeAthlete:
		return nil // Athlete role does not need to be provisioned
	case AccountTypeAdjudicator:
		if !currentUser.HasRole(AccountTypeAdministrator) {
			return errors.New("not authorized to approve user's role application")
		}
	case AccountTypeScrutineer:
		if !currentUser.HasRole(AccountTypeAdministrator) {
			return errors.New("not authorized to approve user's role application")
		}
	case AccountTypeOrganizer:
		if !currentUser.HasRole(AccountTypeAdministrator) {
			return errors.New("not authorized to approve user's role application")
		}
	case AccountTypeDeckCaptain:
		if !(currentUser.HasRole(AccountTypeAdministrator) || currentUser.HasRole(AccountTypeOrganizer)) {
			return errors.New("not authorized to approve user's role application")
		}
	case AccountTypeEmcee:
		if !(currentUser.HasRole(AccountTypeAdministrator) || currentUser.HasRole(AccountTypeOrganizer)) {
			return errors.New("not authorized to approve user's role application")
		}
	default:
		return errors.New("invalid role application")
	}
	return service.respondRoleApplication(currentUser, application, action)
}

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
