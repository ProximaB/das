package businesslogic_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventRegistration_Validate(t *testing.T) {
	registration := businesslogic.EventRegistrationForm{}
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
	athleteEntryRepo.EXPECT().SearchEntry(gomock.Any()).Return([]businesslogic.AthleteCompetitionEntry{
		{ID: 3, AthleteID: 12,
			CompetitionEntry: businesslogic.BaseCompetitionEntry{CompetitionID: 44}},
	}, nil)
	athleteEntryRepo.EXPECT().SearchEntry(gomock.Any()).Return([]businesslogic.AthleteCompetitionEntry{
		{ID: 3, AthleteID: 12,
			CompetitionEntry: businesslogic.BaseCompetitionEntry{CompetitionID: 44}},
	}, nil)

	athleteEventEntryRepo := mock_businesslogic.NewMockIAthleteEventEntryRepository(mockCtrl)

	// partnershipCompEntryRepo := mock_businesslogic.NewMockIPartnershipCompetitionEntryRepository(mockCtrl)
	// partnershipEventEntryRepo := mock_businesslogic.NewMockIPartnershipEventEntryRepository(mockCtrl)

	service := businesslogic.NewCompetitionRegistrationService(
		accountRepo,
		partnershipRepo,
		compRepo,
		eventRepo,
		athleteEntryRepo,
		athleteEventEntryRepo)

	registration := businesslogic.EventRegistrationForm{
		PartnershipID: 33,
		CompetitionID: 127,
		EventsAdded:   []int{},
		EventsDropped: []int{},
	}
	currentUser := businesslogic.Account{}

	err := service.ValidateEventRegistration(currentUser, registration)
	assert.Nil(t, err, "should not return error when registration data is legit")
}
