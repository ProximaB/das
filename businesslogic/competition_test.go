package businesslogic_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var competitions = []businesslogic.Competition{
	{Name: "Test Comp 1"},
	{Name: "Test Comp 2"},
}

func TestCreateCompetition(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	competitionRepo := mock_businesslogic.NewMockICompetitionRepository(mockCtrl)
	provisionRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	provisionHistoryRepo := mock_businesslogic.NewMockIOrganizerProvisionHistoryRepository(mockCtrl)

	comp := businesslogic.Competition{}

	businesslogic.CreateCompetition(comp, competitionRepo, provisionRepo, provisionHistoryRepo)
}

func TestCompetition_UpdateStatus(t *testing.T) {
	comp := businesslogic.Competition{}

	err_1 := comp.UpdateStatus(businesslogic.COMPETITION_STATUS_PRE_REGISTRATION)
	assert.Nil(t, err_1, "change the status of newly instantiated competition should not result in error")

	err_2 := comp.UpdateStatus(businesslogic.COMPETITION_STATUS_IN_PROGRESS)
	assert.Nil(t, err_2, "change the status of competition from pre-registration to in-progress should not result in error")

	err_3 := comp.UpdateStatus(businesslogic.COMPETITION_STATUS_OPEN_REGISTRATION)
	assert.NotNil(t, err_3, "cannot revert competition status from in-progress to open-registration")

}
