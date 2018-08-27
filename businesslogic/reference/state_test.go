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

package reference_test

import (
	"errors"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/mock/businesslogic/reference"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestState_GetCities(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_reference.NewMockICityRepository(mockCtrl)

	// behavior 1
	mockRepo.EXPECT().SearchCity(reference.SearchCityCriteria{StateID: 1}).Return([]reference.City{
		{ID: 1, Name: "City of ID 1", StateID: 1},
		{ID: 2, Name: "City of ID 2", StateID: 1},
	}, nil)

	// behavior 2
	mockRepo.EXPECT().SearchCity(reference.SearchCityCriteria{StateID: 2}).Return(nil,
		errors.New("state does not exist"))

	state_1 := reference.State{ID: 1}
	cities_1, err_1 := state_1.GetCities(mockRepo)
	assert.EqualValues(t, 2, len(cities_1))
	assert.Nil(t, err_1)

	state_2 := reference.State{ID: 2}
	cities_2, err_2 := state_2.GetCities(mockRepo)
	assert.Nil(t, cities_2)
	assert.NotNil(t, err_2)

	cities_3, err_3 := state_1.GetCities(nil)
	assert.Nil(t, cities_3)
	assert.NotNil(t, err_3)

}
