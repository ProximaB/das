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
	// RoleApplicationStatusApproved marks the role application as "approved"
	RoleApplicationStatusApproved = 1
	// RoleApplicationStatusDenied marks the role application as "denied"
	RoleApplicationStatusDenied = 2
	// RoleApplicationStatusPending marks the role application as "pending"
	RoleApplicationStatusPending = 3
)

// SearchRoleApplicationCriteria specifies the search criteria for role application
type SearchRoleApplicationCriteria struct {
	ID             int
	AccountID      int
	AppliedRoleID  int
	StatusID       int
	ApprovalUserID int
	Responded      bool
}

// RoleApplication is an application for restricted roles, including adjudicator, scrutineer, and organizer.
// Non-restrictive roles such as emcee and deck captain can be approved by competition organizers
type RoleApplication struct {
	ID               int
	AccountID        int
	Account          Account
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

// IRoleApplicationRepository specifies the interface that a Role Application Repository should implement
type IRoleApplicationRepository interface {
	CreateApplication(application *RoleApplication) error
	SearchApplication(criteria SearchRoleApplicationCriteria) ([]RoleApplication, error)
	UpdateApplication(application RoleApplication) error
}

// RoleProvisionService is a service that handles Role Application and provision
type RoleProvisionService struct {
	accountRepo         IAccountRepository
	roleApplicationRepo IRoleApplicationRepository
	roleRepo            IAccountRoleRepository
}

// NewRoleProvisionService create a service that serves Role Provision
func NewRoleProvisionService(accountRepo IAccountRepository, roleApplicationRepo IRoleApplicationRepository, roleRepo IAccountRoleRepository) *RoleProvisionService {
	service := RoleProvisionService{
		accountRepo:         accountRepo,
		roleApplicationRepo: roleApplicationRepo,
		roleRepo:            roleRepo,
	}
	return &service
}

// CreateRoleApplication check the validity of the role application and create it if it's valid
func (service RoleProvisionService) CreateRoleApplication(currentUser Account, application *RoleApplication) error {
	// check if current user has the role
	if currentUser.HasRole(application.AppliedRoleID) {
		return errors.New("current user already has the applied role")
	}

	// check if has a pending application
	searchResults, err := service.roleApplicationRepo.SearchApplication(SearchRoleApplicationCriteria{
		AccountID:     currentUser.ID,
		AppliedRoleID: application.AppliedRoleID,
		StatusID:      RoleApplicationStatusPending,
	})
	if err != nil {
		return err
	}
	if len(searchResults) != 0 {
		return errors.New("previous application has not been responded")
	}

	// check what role that user is applying for
	if application.AppliedRoleID == AccountTypeAthlete {
		return errors.New("athlete role should be granted when the account was created")
	}
	if application.AppliedRoleID > AccountTypeEmcee {
		return errors.New("invalid role")
	}

	return service.roleApplicationRepo.CreateApplication(application)
}

func (service RoleProvisionService) respondRoleApplication(currentUser Account, application *RoleApplication, action int) error {
	application.StatusID = action
	application.ApprovalUserID = &currentUser.ID
	application.DateTimeApproved = time.Now()
	if updateErr := service.roleApplicationRepo.UpdateApplication(*application); updateErr != nil {
		return updateErr
	}
	if action == RoleApplicationStatusApproved {
		role := AccountRole{
			AccountID:       application.AccountID,
			AccountTypeID:   application.AppliedRoleID,
			CreateUserID:    currentUser.ID,
			DateTimeCreated: time.Now(),
			UpdateUserID:    currentUser.ID,
			DateTimeUpdated: time.Now(),
		}
		return service.roleRepo.CreateAccountRole(&role)
	}
	return nil
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
	// check if application is pending
	if application.StatusID == RoleApplicationStatusApproved || application.StatusID == RoleApplicationStatusDenied {
		return errors.New("role application is already responded")
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

// SearchRoleApplication searches the available role application based on current user's privilege
// TODO: this is not working correctly!
func (service RoleProvisionService) SearchRoleApplication(currentUser Account, criteria SearchRoleApplicationCriteria) ([]RoleApplication, error) {
	return service.roleApplicationRepo.SearchApplication(criteria)
}

// OrganizerProvision provision organizer competition slots for creating and hosting competitions
type OrganizerProvision struct {
	ID              int
	OrganizerID     int
	Organizer       Account
	Available       int
	Hosted          int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

// SearchOrganizerProvisionCriteria specifies the search criteria of Organizer's provision information
type SearchOrganizerProvisionCriteria struct {
	ID          int `schema:"organizer"`
	OrganizerID int `schema:"organizer"`
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
