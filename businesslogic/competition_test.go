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
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateCompetition(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	competitionRepo := mock_businesslogic.NewMockICompetitionRepository(mockCtrl)
	provisionRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	provisionHistoryRepo := mock_businesslogic.NewMockIOrganizerProvisionHistoryRepository(mockCtrl)

	comp := businesslogic.Competition{
		Name:          "Intergalactic Competition",
		Website:       "http://www.example.com",
		FederationID:  1,
		StartDateTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+2, 1, 1, 1, 1, time.UTC),
		EndDateTime:   time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+4, 1, 1, 1, 1, time.UTC),
		ContactName:   "James Bond",
		ContactEmail:  "james.bond@email.com",
		ContactPhone:  "2290092292",
		City:          reference.City{ID: 26},
		State:         reference.State{ID: 17},
		Country:       reference.Country{ID: 19},
		CreateUserID:  1,
		UpdateUserID:  1,
	}

	competitionRepo.EXPECT().CreateCompetition(&comp).Return(nil)
	provisionRepo.EXPECT().SearchOrganizerProvision(businesslogic.SearchOrganizerProvisionCriteria{
		OrganizerID: 1,
	}).Return([]businesslogic.OrganizerProvision{
		{ID: 3, OrganizerRoleID: 1, Available: 3, Hosted: 7},
	}, nil)
	provisionRepo.EXPECT().UpdateOrganizerProvision(gomock.Any()).Return(nil)
	provisionHistoryRepo.EXPECT().CreateOrganizerProvisionHistory(gomock.Any()).Return(nil)

	err := businesslogic.CreateCompetition(comp, competitionRepo, provisionRepo, provisionHistoryRepo)
	assert.Nil(t, err, "should create competition if competition data is correct and organizer has sufficient provision")
}

func TestCompetition_UpdateStatus(t *testing.T) {
	comp := businesslogic.Competition{}

	err_1 := comp.UpdateStatus(businesslogic.CompetitionStatusPreRegistration)
	assert.Nil(t, err_1, "change the status of newly instantiated competition should not result in error")

	err_2 := comp.UpdateStatus(businesslogic.CompetitionStatusInProgress)
	assert.Nil(t, err_2, "change the status of competition from pre-registration to in-progress should not result in error")

	err_3 := comp.UpdateStatus(businesslogic.CompetitionStatusOpenRegistration)
	assert.NotNil(t, err_3, "cannot revert competition status from in-progress to open-registration")

}
