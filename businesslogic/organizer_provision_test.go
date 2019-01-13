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

package businesslogic_test

import (
	"errors"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetOrganizerProvision(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	mockRepo.EXPECT().SearchOrganizerProvision(businesslogic.SearchOrganizerProvisionCriteria{
		OrganizerID: 1,
	}).Return([]businesslogic.OrganizerProvision{
		{ID: 1, OrganizerRoleID: 1, Available: 1, Hosted: 2},
	}, nil)

	res_1, err_1 := mockRepo.SearchOrganizerProvision(businesslogic.SearchOrganizerProvisionCriteria{
		OrganizerID: 1,
	})

	assert.Len(t, res_1, 1)
	assert.Nil(t, err_1)
}

func TestGetOrganizerProvision_Invalid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	mockRepo.EXPECT().SearchOrganizerProvision(businesslogic.SearchOrganizerProvisionCriteria{
		OrganizerID: 0,
	}).Return(nil, errors.New("invalid search"))

	res_2, err_2 := mockRepo.SearchOrganizerProvision(businesslogic.SearchOrganizerProvisionCriteria{
		OrganizerID: 0,
	})

	assert.Nil(t, res_2)
	assert.NotNil(t, err_2)
}

func TestNewRowProvisionService(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockRoleAppRepo := mock_businesslogic.NewMockIRoleApplicationRepository(mockCtrl)
	mockRoleRepo := mock_businesslogic.NewMockIAccountRoleRepository(mockCtrl)
	mockOrgProvRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	mockOrgProvHistRepo := mock_businesslogic.NewMockIOrganizerProvisionHistoryRepository(mockCtrl)
	service := businesslogic.NewRoleProvisionService(mockAccountRepo, mockRoleAppRepo, mockRoleRepo, mockOrgProvRepo, mockOrgProvHistRepo)

	assert.NotNil(t, service)
}

func TestRoleProvisionService_ApproveApplication(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockRoleAppRepo := mock_businesslogic.NewMockIRoleApplicationRepository(mockCtrl)
	mockRoleRepo := mock_businesslogic.NewMockIAccountRoleRepository(mockCtrl)
	mockOrgProvRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	mockOrgProvHistRepo := mock_businesslogic.NewMockIOrganizerProvisionHistoryRepository(mockCtrl)
	service := businesslogic.NewRoleProvisionService(mockAccountRepo, mockRoleAppRepo, mockRoleRepo, mockOrgProvRepo, mockOrgProvHistRepo)

	application := businesslogic.RoleApplication{
		AccountID:       33,
		AppliedRoleID:   businesslogic.AccountTypeAthlete,
		StatusID:        businesslogic.RoleApplicationStatusPending,
		DateTimeCreated: time.Now(),
		DateTimeUpdated: time.Now(),
	}

	currentUser := businesslogic.Account{
		ID: 18,
	}
	currentUser.SetRoles([]businesslogic.AccountRole{
		{ID: 3, AccountID: 18, AccountTypeID: businesslogic.AccountTypeAthlete},
		{ID: 3, AccountID: 18, AccountTypeID: businesslogic.AccountTypeOrganizer},
		{ID: 3, AccountID: 18, AccountTypeID: businesslogic.AccountTypeAdministrator},
	})

	err := service.UpdateApplication(currentUser, &application, businesslogic.RoleApplicationStatusApproved)
	assert.Nil(t, err, "should not throw error when 'approving' an Athlete role application")
}

func TestRoleProvisionService_ApproveApplication_ApplyToBeAdmin(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockRoleAppRepo := mock_businesslogic.NewMockIRoleApplicationRepository(mockCtrl)
	mockRoleRepo := mock_businesslogic.NewMockIAccountRoleRepository(mockCtrl)
	mockOrgProvRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	mockOrgProvHistRepo := mock_businesslogic.NewMockIOrganizerProvisionHistoryRepository(mockCtrl)
	service := businesslogic.NewRoleProvisionService(mockAccountRepo, mockRoleAppRepo, mockRoleRepo, mockOrgProvRepo, mockOrgProvHistRepo)

	application := businesslogic.RoleApplication{
		AccountID:       33,
		AppliedRoleID:   businesslogic.AccountTypeAdministrator,
		StatusID:        businesslogic.RoleApplicationStatusPending,
		DateTimeCreated: time.Now(),
		DateTimeUpdated: time.Now(),
	}

	currentUser := businesslogic.Account{
		ID: 18,
	}
	currentUser.SetRoles([]businesslogic.AccountRole{
		{ID: 3, AccountID: 18, AccountTypeID: businesslogic.AccountTypeAthlete},
		{ID: 4, AccountID: 18, AccountTypeID: businesslogic.AccountTypeAdjudicator},
		{ID: 5, AccountID: 18, AccountTypeID: businesslogic.AccountTypeScrutineer},
		{ID: 8, AccountID: 18, AccountTypeID: businesslogic.AccountTypeOrganizer},
		{ID: 6, AccountID: 18, AccountTypeID: businesslogic.AccountTypeEmcee},
		{ID: 7, AccountID: 18, AccountTypeID: businesslogic.AccountTypeDeckCaptain},
		{ID: 3, AccountID: 18, AccountTypeID: businesslogic.AccountTypeAdministrator},
	})

	err := service.UpdateApplication(currentUser, &application, businesslogic.RoleApplicationStatusApproved)
	assert.Error(t, err, "administrator role application should not be approved by another administrator")
}

func TestRoleProvisionService_ApproveApplication_LowPrivilegeAttemptsToProvisionHighPrivilege(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockRoleAppRepo := mock_businesslogic.NewMockIRoleApplicationRepository(mockCtrl)
	mockRoleRepo := mock_businesslogic.NewMockIAccountRoleRepository(mockCtrl)
	mockOrgProvRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	mockOrgProvHistRepo := mock_businesslogic.NewMockIOrganizerProvisionHistoryRepository(mockCtrl)
	service := businesslogic.NewRoleProvisionService(mockAccountRepo, mockRoleAppRepo, mockRoleRepo, mockOrgProvRepo, mockOrgProvHistRepo)

	application := businesslogic.RoleApplication{
		AccountID:       33,
		AppliedRoleID:   businesslogic.AccountTypeOrganizer,
		StatusID:        businesslogic.RoleApplicationStatusPending,
		DateTimeCreated: time.Now(),
		DateTimeUpdated: time.Now(),
	}

	currentUser := businesslogic.Account{
		ID: 18,
	}
	currentUser.SetRoles([]businesslogic.AccountRole{
		{ID: 3, AccountID: 18, AccountTypeID: businesslogic.AccountTypeAthlete},
		{ID: 4, AccountID: 18, AccountTypeID: businesslogic.AccountTypeAdjudicator},
		{ID: 5, AccountID: 18, AccountTypeID: businesslogic.AccountTypeScrutineer},
		{ID: 6, AccountID: 18, AccountTypeID: businesslogic.AccountTypeEmcee},
		{ID: 7, AccountID: 18, AccountTypeID: businesslogic.AccountTypeDeckCaptain},
	})

	err := service.UpdateApplication(currentUser, &application, businesslogic.RoleApplicationStatusApproved)
	assert.Error(t, err, "athlete cannot approve any positions")
}

func TestRoleProvisionService_ApproveApplication_SelfApproval_OrganizerApproveScrutineer(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockRoleAppRepo := mock_businesslogic.NewMockIRoleApplicationRepository(mockCtrl)
	mockRoleRepo := mock_businesslogic.NewMockIAccountRoleRepository(mockCtrl)
	mockOrgProvRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	mockOrgProvHistRepo := mock_businesslogic.NewMockIOrganizerProvisionHistoryRepository(mockCtrl)
	service := businesslogic.NewRoleProvisionService(mockAccountRepo, mockRoleAppRepo, mockRoleRepo, mockOrgProvRepo, mockOrgProvHistRepo)

	application := businesslogic.RoleApplication{
		AccountID:       33,
		AppliedRoleID:   businesslogic.AccountTypeScrutineer,
		StatusID:        businesslogic.RoleApplicationStatusPending,
		DateTimeCreated: time.Now(),
		DateTimeUpdated: time.Now(),
	}

	currentUser := businesslogic.Account{
		ID: 33,
	}
	currentUser.SetRoles([]businesslogic.AccountRole{
		{ID: 3, AccountID: 33, AccountTypeID: businesslogic.AccountTypeOrganizer},
	})

	err := service.UpdateApplication(currentUser, &application, businesslogic.RoleApplicationStatusApproved)
	assert.Error(t, err, "organizer cannot approve a scrutineer role application")
}

func TestRoleProvisionService_UpdateApplication_ValidApplication(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockRoleAppRepo := mock_businesslogic.NewMockIRoleApplicationRepository(mockCtrl)
	mockRoleRepo := mock_businesslogic.NewMockIAccountRoleRepository(mockCtrl)
	mockOrgProvRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	mockOrgProvHistRepo := mock_businesslogic.NewMockIOrganizerProvisionHistoryRepository(mockCtrl)
	service := businesslogic.NewRoleProvisionService(mockAccountRepo, mockRoleAppRepo, mockRoleRepo, mockOrgProvRepo, mockOrgProvHistRepo)

	currentUser := businesslogic.Account{
		ID: 31,
	}
	currentUser.SetRoles([]businesslogic.AccountRole{
		{ID: 31, AccountID: 33, AccountTypeID: businesslogic.AccountTypeAthlete},
		{ID: 31, AccountID: 33, AccountTypeID: businesslogic.AccountTypeOrganizer},
		{ID: 31, AccountID: 33, AccountTypeID: businesslogic.AccountTypeAdministrator},
	})

	application := businesslogic.RoleApplication{
		ID:            3,
		AccountID:     7,
		AppliedRoleID: businesslogic.AccountTypeOrganizer,
		StatusID:      businesslogic.RoleApplicationStatusPending,
	}

	mockRoleRepo.EXPECT().SearchAccountRole(gomock.Any()).Return([]businesslogic.AccountRole{
		{ID: 22, AccountID: 7, AccountTypeID: businesslogic.AccountTypeAthlete},
	}, nil)
	mockRoleRepo.EXPECT().SearchAccountRole(gomock.Any()).Return([]businesslogic.AccountRole{
		{ID: 22, AccountID: 7, AccountTypeID: businesslogic.AccountTypeAthlete},
	}, nil)
	mockRoleAppRepo.EXPECT().UpdateApplication(gomock.Any()).Return(nil)
	mockRoleRepo.EXPECT().CreateAccountRole(gomock.Any()).Return(nil)
	mockOrgProvRepo.EXPECT().CreateOrganizerProvision(gomock.Any()).Return(nil)
	mockOrgProvHistRepo.EXPECT().CreateOrganizerProvisionHistory(gomock.Any()).Return(nil)

	err := service.UpdateApplication(currentUser, &application, businesslogic.RoleApplicationStatusApproved)
	assert.Nil(t, err, "should not throw an error if the application is legit and current user has access")
}
