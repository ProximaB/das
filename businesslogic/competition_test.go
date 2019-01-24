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
	"testing"
	"time"

	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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
		{ID: 3, OrganizerID: 1, Available: 3, Hosted: 7},
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

func TestCompetition_GetStatus(t *testing.T) {
	comp := businesslogic.Competition{}
	comp.UpdateStatus(businesslogic.CompetitionStatusCancelled)

	assert.Equal(t, businesslogic.CompetitionStatusCancelled, comp.GetStatus())
}

// GetCompetitionByID test helpers
type getCompetitionByIDResult struct {
	comp businesslogic.Competition
	err  error
}

func twoValueReturnHandler(c businesslogic.Competition, e error) getCompetitionByIDResult {
	result := getCompetitionByIDResult{comp: c, err: e}

	return result
}

func getCompetitionByIDMockHandler(m *gomock.Controller, id int, r []businesslogic.Competition,
	e error) businesslogic.ICompetitionRepository {
	searchComp := businesslogic.SearchCompetitionCriteria{ID: id}
	competitionRepo := mock_businesslogic.NewMockICompetitionRepository(m)
	competitionRepo.EXPECT().SearchCompetition(searchComp).Return(r, e).MaxTimes(2)

	return competitionRepo
}

func getCompetitionByIDAssertNilHandler(t *testing.T, competitionRepo businesslogic.ICompetitionRepository) {
	assert.Equal(
		t,
		twoValueReturnHandler(businesslogic.Competition{}, errors.New("Return an error")).comp,
		twoValueReturnHandler(businesslogic.GetCompetitionByID(2, competitionRepo)).comp,
	)
	assert.Nil(t, twoValueReturnHandler(businesslogic.GetCompetitionByID(2, competitionRepo)).err)
}

// GetCompetitionByID tests
func TestCompetition_GetCompetitionByID_ErrorNotNil(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	competitionRepo := getCompetitionByIDMockHandler(
		mockCtrl,
		2,
		[]businesslogic.Competition{},
		errors.New("Return empty competitions and a database error."),
	)
	assert.Equal(
		t,
		twoValueReturnHandler(businesslogic.Competition{}, errors.New("Return an error")).comp,
		twoValueReturnHandler(businesslogic.GetCompetitionByID(2, competitionRepo)).comp,
	)
	assert.Error(t, twoValueReturnHandler(businesslogic.GetCompetitionByID(2, competitionRepo)).err)

}

func TestCompetition_GetCompetitionByID_SearchResultNil(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	competitionRepo := getCompetitionByIDMockHandler(mockCtrl, 2, nil, nil)
	getCompetitionByIDAssertNilHandler(t, competitionRepo)
}

func TestCompetition_GetCompetitionByID_SearchResultLengthNotOne(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	competitionRepo := getCompetitionByIDMockHandler(mockCtrl, 2, make([]businesslogic.Competition, 2), nil)
	getCompetitionByIDAssertNilHandler(t, competitionRepo)
}

func TestCompetition_GetCompetitionByID_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	competitionRepo := getCompetitionByIDMockHandler(mockCtrl, 2, []businesslogic.Competition{}, nil)
	getCompetitionByIDAssertNilHandler(t, competitionRepo)
}

func TestCompetition_UpdateCompetition_ErrorNotNil(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user := businesslogic.Account{ID: 2}
	competition := businesslogic.OrganizerUpdateCompetition{
		CompetitionID: 2,
		Status:        businesslogic.CompetitionStatusInProgress,
		Name:          "The Great American Ball",
		Website:       "www.tgab.com",
		StartDate:     time.Date(2018, time.November, 18, 9, 0, 0, 0, time.UTC),
		EndDate:       time.Date(2018, time.November, 19, 22, 0, 0, 0, time.UTC),
	}
	searchComp := businesslogic.SearchCompetitionCriteria{ID: competition.CompetitionID}
	competitionRepo := mock_businesslogic.NewMockICompetitionRepository(mockCtrl)
	competitionRepo.EXPECT().SearchCompetition(searchComp).Return(
		[]businesslogic.Competition{},
		errors.New("Return an error"),
	)

	assert.Error(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}
