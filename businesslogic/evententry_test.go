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
)

func TestCreatePartnershipEventEntry_NoError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mock_businesslogic.NewMockIPartnershipEventEntryRepository(mockCtrl)

	// defines expected behavior
	repo.EXPECT().SearchPartnershipEventEntry(gomock.Any()).Return([]businesslogic.PartnershipEventEntry{}, nil)
	repo.EXPECT().CreatePartnershipEventEntry(gomock.Any()).Return(nil)

	entry := businesslogic.PartnershipEventEntry{
		PartnershipID: 3,
		EventEntry: businesslogic.EventEntry{
			EventID: 12,
		},
	}

	result := businesslogic.CreatePartnershipEventEntry(entry, repo)
	assert.Nil(t, result, "should not return error if no entry exists")
}

func TestCreatePartnershipEventEntry_WithError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mock_businesslogic.NewMockIPartnershipEventEntryRepository(mockCtrl)

	// defines expected behavior
	repo.EXPECT().SearchPartnershipEventEntry(gomock.Any()).Return([]businesslogic.PartnershipEventEntry{}, errors.New("a random error"))

	entry := businesslogic.PartnershipEventEntry{
		PartnershipID: 3,
		EventEntry: businesslogic.EventEntry{
			EventID: 12,
		},
	}

	result := businesslogic.CreatePartnershipEventEntry(entry, repo)
	assert.NotNil(t, result, "should return error if searching process has error")
}

func TestCreatePartnershipEventEntry_ExistingEntry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mock_businesslogic.NewMockIPartnershipEventEntryRepository(mockCtrl)

	// defines expected behavior
	repo.EXPECT().SearchPartnershipEventEntry(gomock.Any()).Return([]businesslogic.PartnershipEventEntry{
		{PartnershipID: 3, EventEntry: businesslogic.EventEntry{
			EventID: 12,
		}},
	}, nil)

	entry := businesslogic.PartnershipEventEntry{
		PartnershipID: 3,
		EventEntry: businesslogic.EventEntry{
			EventID: 12,
		},
	}

	result := businesslogic.CreatePartnershipEventEntry(entry, repo)
	assert.NotNil(t, result, "should return error if a matching entry is found")
}
