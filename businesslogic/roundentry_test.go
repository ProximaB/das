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

func TestCreatePartnershipRoundEntry_NoError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mock_businesslogic.NewMockIPartnershipRoundEntryRepository(mockCtrl)

	repo.EXPECT().SearchPartnershipRoundEntry(gomock.Any()).Return([]businesslogic.PartnershipRoundEntry{}, nil)
	repo.EXPECT().CreatePartnershipRoundEntry(gomock.Any()).Return(nil)

	result := businesslogic.CreatePartnershipRoundEntry(&businesslogic.PartnershipRoundEntry{
		PartnershipID: 13,
		RoundEntry:    businesslogic.RoundEntry{RoundID: 33},
	}, repo)

	assert.Nil(t, result, "should not return error when search and create entry does not return error")
}
