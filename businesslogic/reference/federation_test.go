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

package referencebll_test

import (
	"errors"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	mock_reference "github.com/DancesportSoftware/das/mock/businesslogic/reference"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFederation_GetDivisions(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_reference.NewMockIDivisionRepository(mockCtrl)

	// behavior 1
	mockRepo.EXPECT().SearchDivision(referencebll.SearchDivisionCriteria{FederationID: 1}).Return([]referencebll.Division{
		{ID: 1, Name: "Correct Division 1", FederationID: 1},
		{ID: 2, Name: "Correct Division 2", FederationID: 2},
	}, nil)

	// behavior 2
	mockRepo.EXPECT().SearchDivision(referencebll.SearchDivisionCriteria{FederationID: 2}).Return(nil, errors.New("invalid search"))

	federation_1 := referencebll.Federation{ID: 1}
	federation_2 := referencebll.Federation{ID: 2}

	result_1, err_1 := federation_1.GetDivisions(mockRepo)
	assert.EqualValues(t, 2, len(result_1))
	assert.Nil(t, err_1)

	result_2, err_2 := federation_2.GetDivisions(mockRepo)
	assert.Nil(t, result_2)
	assert.NotNil(t, err_2)

	result_3, err_3 := federation_1.GetDivisions(nil)
	assert.Nil(t, result_3)
	assert.NotNil(t, err_3)
}
