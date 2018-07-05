// Dancesport Application System (DAS)
// Copyright (C) 2018 Yubing Hou
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

func TestEventRegistration_Validate(t *testing.T) {
	registration := businesslogic.EventRegistration{}
	assert.NotNil(t, registration.Validate(), "should throw an error if Competition is not specified")

	registration.CompetitionID = 12
	assert.NotNil(t, registration.Validate(), "should throw an error if Partnership is not specified")

	registration.PartnershipID = 51
	assert.Nil(t, registration.Validate(), "registration data should be valid if Competition and Partnership is provided")
}

func TestCompetitionRegistrationService_ValidateEventRegistration_LegitimateData(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	competition := businesslogic.Competition{ID: 44, Name: "Awesome Competition"}
	competition.UpdateStatus(businesslogic.CompetitionStatusOpenRegistration)

	accountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	accountRepo.EXPECT().SearchAccount(gomock.Any()).Return([]businesslogic.Account{
		{ID: 1, FirstName: "Alice"},
	}, nil)
	partnershipRepo := mock_businesslogic.NewMockIPartnershipRepository(mockCtrl)
	partnershipRepo.EXPECT().SearchPartnership(gomock.Any()).Return([]businesslogic.Partnership{
		{ID: 837},
	}, nil)
	compRepo := mock_businesslogic.NewMockICompetitionRepository(mockCtrl)
	compRepo.EXPECT().SearchCompetition(gomock.Any()).Return([]businesslogic.Competition{
		competition,
	}, nil)
	eventRepo := mock_businesslogic.NewMockIEventRepository(mockCtrl)
	athleteEntryRepo := mock_businesslogic.NewMockIAthleteCompetitionEntryRepository(mockCtrl)
	athleteEntryRepo.EXPECT().SearchAthleteCompetitionEntry(gomock.Any()).Return([]businesslogic.AthleteCompetitionEntry{
		{ID: 3, AthleteID: 12,
			CompetitionEntry: businesslogic.CompetitionEntry{CompetitionID: 44}},
	}, nil)
	athleteEntryRepo.EXPECT().SearchAthleteCompetitionEntry(gomock.Any()).Return([]businesslogic.AthleteCompetitionEntry{
		{ID: 3, AthleteID: 12,
			CompetitionEntry: businesslogic.CompetitionEntry{CompetitionID: 44}},
	}, nil)
	partnershipCompEntryRepo := mock_businesslogic.NewMockIPartnershipCompetitionEntryRepository(mockCtrl)
	partnershipEventEntryRepo := mock_businesslogic.NewMockIPartnershipEventEntryRepository(mockCtrl)

	service := businesslogic.CompetitionRegistrationService{
		accountRepo,
		partnershipRepo,
		compRepo,
		eventRepo,
		athleteEntryRepo,
		partnershipCompEntryRepo,
		partnershipEventEntryRepo,
	}

	registration := businesslogic.EventRegistration{
		PartnershipID: 33,
		CompetitionID: 127,
	}
	currentUser := businesslogic.Account{}

	err := service.ValidateEventRegistration(currentUser, registration)
	assert.Nil(t, err, "should not return error when registration data is legit")
}
