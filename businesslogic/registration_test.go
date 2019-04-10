package businesslogic_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

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

	compRepo := mock_businesslogic.NewMockICompetitionRepository(mockCtrl)

	eventRepo := mock_businesslogic.NewMockIEventRepository(mockCtrl)
	athleteEntryRepo := mock_businesslogic.NewMockIAthleteCompetitionEntryRepository(mockCtrl)
	athleteEntryRepo.EXPECT().SearchEntry(gomock.Any()).Return([]businesslogic.AthleteCompetitionEntry{
		{ID: 3, Athlete: businesslogic.Account{ID: 12},
			Competition: businesslogic.Competition{ID: 44}},
	}, nil)
	athleteEntryRepo.EXPECT().SearchEntry(gomock.Any()).Return([]businesslogic.AthleteCompetitionEntry{
		{ID: 3, Athlete: businesslogic.Account{ID: 12},
			Competition: businesslogic.Competition{ID: 44}},
	}, nil)

	athleteEventEntryRepo := mock_businesslogic.NewMockIAthleteEventEntryRepository(mockCtrl)

	partnershipCompEntryRepo := mock_businesslogic.NewMockIPartnershipCompetitionEntryRepository(mockCtrl)
	partnershipEventEntryRepo := mock_businesslogic.NewMockIPartnershipEventEntryRepository(mockCtrl)

	service := businesslogic.NewCompetitionRegistrationService(
		accountRepo,
		partnershipRepo,
		compRepo,
		eventRepo,
		athleteEntryRepo,
		athleteEventEntryRepo, partnershipCompEntryRepo, partnershipEventEntryRepo)

	registration := businesslogic.EventRegistrationForm{
		Couple:        businesslogic.Partnership{ID: 33},
		Competition:   businesslogic.Competition{ID: 127},
		EventsAdded:   make([]businesslogic.Event, 0),
		EventsDropped: make([]businesslogic.Event, 0),
	}
	currentUser := businesslogic.Account{}

	err := service.ValidateEventRegistration(currentUser, registration)
	assert.Nil(t, err, "should not return error when registration data is legit")
}
