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

func TestGetOrganizerProvision(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	mockRepo.EXPECT().SearchOrganizerProvision(businesslogic.SearchOrganizerProvisionCriteria{
		OrganizerID: 1,
	}).Return([]businesslogic.OrganizerProvision{
		{ID: 1, OrganizerID: 1, Available: 1, Hosted: 2},
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
