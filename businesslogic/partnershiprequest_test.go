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
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePartnershipRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	request := businesslogic.PartnershipRequest{
		SenderID:      12,
		RecipientID:   33,
		SenderRole:    businesslogic.PartnershipRoleLead,
		RecipientRole: businesslogic.PartnershipRoleFollow,
		Message:       "Hi, can you add me please?",
		Status:        businesslogic.PartnershipRequestStatusPending,
	}

	accountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	requestRepo := mock_businesslogic.NewMockIPartnershipRequestRepository(mockCtrl)
	partnershipRepo := mock_businesslogic.NewMockIPartnershipRepository(mockCtrl)
	blacklistRepo := mock_businesslogic.NewMockIPartnershipRequestBlacklistRepository(mockCtrl)

	rolesOfLeadAccount := []businesslogic.AccountRole{
		{ID: 1, AccountID: 1, AccountTypeID: businesslogic.AccountTypeOrganizer},
		{ID: 2, AccountID: 1, AccountTypeID: businesslogic.AccountTypeAthlete},
	}
	leadAccount := businesslogic.Account{
		ID: 1,
	}
	leadAccount.SetRoles(rolesOfLeadAccount)

	rolesOfFollowAccount := []businesslogic.AccountRole{
		{ID: 3, AccountID: 2, AccountTypeID: businesslogic.AccountTypeAdjudicator},
		{ID: 4, AccountID: 2, AccountTypeID: businesslogic.AccountTypeDeckCaptain},
		{ID: 5, AccountID: 2, AccountTypeID: businesslogic.AccountTypeAthlete},
	}
	followAccount := businesslogic.Account{
		ID: 2,
	}
	followAccount.SetRoles(rolesOfFollowAccount)

	// specify behaviors
	accountRepo.EXPECT().SearchAccount(gomock.Any()).Return([]businesslogic.Account{
		leadAccount,
	}, nil)
	accountRepo.EXPECT().SearchAccount(gomock.Any()).Return([]businesslogic.Account{
		followAccount,
	}, nil)
	blacklistRepo.EXPECT().SearchPartnershipRequestBlacklist(gomock.Any()).Return([]businesslogic.PartnershipRequestBlacklistEntry{}, nil)
	partnershipRepo.EXPECT().SearchPartnership(gomock.Any()).Return([]businesslogic.Partnership{}, nil)
	requestRepo.EXPECT().SearchPartnershipRequest(gomock.Any()).Return([]businesslogic.PartnershipRequest{}, nil)
	requestRepo.EXPECT().CreatePartnershipRequest(gomock.Any()).Return(nil)

	err := businesslogic.CreatePartnershipRequest(request, partnershipRepo, requestRepo, accountRepo, blacklistRepo)
	assert.Nil(t, err, "should not throw an error if every step is working correctly")
}
