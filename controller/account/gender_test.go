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

package account

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/mock/businesslogic/reference"
	"github.com/DancesportSoftware/das/viewmodel"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAccountGenderHandler_GetAccountGenderHandler(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedGenderRepo := mock_reference.NewMockIGenderRepository(mockCtrl)

	server := GenderServer{
		IGenderRepository: mockedGenderRepo,
	}

	req, _ := http.NewRequest(http.MethodGet, "/api/account/gender", nil)
	w := httptest.NewRecorder()

	// test with zero param
	mockedGenderRepo.EXPECT().GetAllGenders().Return([]reference.Gender{
		{ID: 1, Name: "Female"},
		{ID: 2, Name: "Male"},
	}, nil)
	server.GetAccountGenderHandler(w, req)
	genders := make([]viewmodel.Gender, 0)
	err := json.Unmarshal([]byte(w.Body.String()), &genders)
	assert.Nil(t, err, "should return a list of genders")
	w.Flush()

	// test with bad param should not result in error
	query := req.URL.Query()
	query.Add("badparam", "indeed")
	req.URL.RawQuery = query.Encode()
	// log.Printf("GET %s\n", req.URL.String())
	mockedGenderRepo.EXPECT().GetAllGenders().Return([]reference.Gender{
		{ID: 1, Name: "Female"},
		{ID: 2, Name: "Male"},
	}, nil)
	server.GetAccountGenderHandler(w, req)
	// log.Println(w.Body)
	assert.EqualValues(t, http.StatusOK, w.Code, "should not receive HTTP 400 when sending a bad request")
}
