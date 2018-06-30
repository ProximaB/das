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

func TestAthleteCompetitionEntry_CreateAthleteCompetitionEntry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	entryRepo := mock_businesslogic.NewMockIAthleteCompetitionEntryRepository(mockCtrl)
	entryRepo.EXPECT().SearchAthleteCompetitionEntry(businesslogic.SearchAthleteCompetitionEntryCriteria{
		AthleteID:     12,
		CompetitionID: 44,
	}).Return([]businesslogic.AthleteCompetitionEntry{
		{ID: 3, AthleteID: 12,
			CompetitionEntry: businesslogic.CompetitionEntry{CompetitionID: 44}},
	}, nil)

	entry := businesslogic.AthleteCompetitionEntry{
		AthleteID:        12,
		CompetitionEntry: businesslogic.CompetitionEntry{CompetitionID: 44},
	}
	competition := businesslogic.Competition{ID: 44, Name: "Awesome Competition"}
	competition.UpdateStatus(businesslogic.CompetitionStatusOpenRegistration)

	compRepo := mock_businesslogic.NewMockICompetitionRepository(mockCtrl)
	compRepo.EXPECT().SearchCompetition(gomock.Any()).Return(
		[]businesslogic.Competition{
			competition,
		}, nil)

	err := entry.CreateAthleteCompetitionEntry(compRepo, entryRepo)
	assert.NotNil(t, err, "should create duplicate competition entry with error")

	entryRepo.EXPECT().SearchAthleteCompetitionEntry(businesslogic.SearchAthleteCompetitionEntryCriteria{
		AthleteID:     12,
		CompetitionID: 44,
	}).Return([]businesslogic.AthleteCompetitionEntry{}, nil)
	entryRepo.EXPECT().CreateAthleteCompetitionEntry(gomock.Any()).Return(nil)
	compRepo.EXPECT().SearchCompetition(gomock.Any()).Return(
		[]businesslogic.Competition{
			competition,
		}, nil)
	err = entry.CreateAthleteCompetitionEntry(compRepo, entryRepo)
	assert.Nil(t, err, "should create new competition entry without error")
}
