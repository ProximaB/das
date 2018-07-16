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

	// specify behaviors
	accountRepo.EXPECT().SearchAccount(gomock.Any()).Return([]businesslogic.Account{
		{ID: 12},
	}, nil)
	accountRepo.EXPECT().SearchAccount(gomock.Any()).Return([]businesslogic.Account{
		{ID: 33},
	}, nil)
	blacklistRepo.EXPECT().SearchPartnershipRequestBlacklist(gomock.Any()).Return([]businesslogic.PartnershipRequestBlacklistEntry{}, nil)
	partnershipRepo.EXPECT().SearchPartnership(gomock.Any()).Return([]businesslogic.Partnership{}, nil)
	requestRepo.EXPECT().SearchPartnershipRequest(gomock.Any()).Return([]businesslogic.PartnershipRequest{}, nil)
	requestRepo.EXPECT().CreatePartnershipRequest(gomock.Any()).Return(nil)

	err := businesslogic.CreatePartnershipRequest(request, partnershipRepo, requestRepo, accountRepo, blacklistRepo)
	assert.Nil(t, err, "should not throw an error if every step is working correctly")
}
