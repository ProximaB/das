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

func TestAccount_GetBlacklistedAccounts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user_17 := businesslogic.Account{ID: 17}

	accountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	blacklistRepo := mock_businesslogic.NewMockIPartnershipRequestBlacklistRepository(mockCtrl)

	blacklistRepo.EXPECT().SearchPartnershipRequestBlacklist(businesslogic.SearchPartnershipRequestBlacklistCriteria{ReporterID: 17, Whitelisted: false}).Return([]businesslogic.PartnershipRequestBlacklistEntry{
		{ID: 21, BlockedUserID: 33, ReporterID: 17},
	}, nil)
	accountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{ID: 33}).Return([]businesslogic.Account{
		{ID: 33, FirstName: "Sharp", LastName: "Shark"},
	}, nil)
	blacklist, err := user_17.GetBlacklistedAccounts(accountRepo, blacklistRepo)

	assert.Nil(t, err, "should retrieve blacklisted accounts without error")
	assert.EqualValues(t, 1, len(blacklist), "should get all blacklisted account without error")
}
