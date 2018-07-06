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

func TestAccount_GetAllPartnerships(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_businesslogic.NewMockIPartnershipRepository(mockCtrl)
	mockRepo.EXPECT().SearchPartnership(businesslogic.SearchPartnershipCriteria{
		LeadID: 9,
	}).Return([]businesslogic.Partnership{
		{ID: 1, LeadID: 9, FollowID: 8},
		{ID: 2, LeadID: 9, FollowID: 3},
	}, nil)
	mockRepo.EXPECT().SearchPartnership(businesslogic.SearchPartnershipCriteria{
		FollowID: 9,
	}).Return([]businesslogic.Partnership{
		{ID: 7, LeadID: 33, FollowID: 9},
	}, nil)

	athlete := businesslogic.Account{ID: 9}
	partnerships, err := athlete.GetAllPartnerships(mockRepo)

	assert.Nil(t, err, "should get all partnerships without error")
	assert.EqualValues(t, 3, len(partnerships), "should get all partnerships as lead and follow")
}
