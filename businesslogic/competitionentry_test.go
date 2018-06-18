package businesslogic_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompetitionEntry_CreateCompetitionEntry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	entryRepo := mock_businesslogic.NewMockICompetitionEntryRepository(mockCtrl)
	entryRepo.EXPECT().SearchCompetitionEntry(businesslogic.SearchCompetitionEntryCriteria{
		AthleteID:     12,
		CompetitionID: 44,
	}).Return([]businesslogic.CompetitionEntry{
		{ID: 3, AthleteID: 12, CompetitionID: 44},
	}, nil)

	entry := businesslogic.CompetitionEntry{
		AthleteID:     12,
		CompetitionID: 44,
	}
	competition := businesslogic.Competition{ID: 44, Name: "Awesome Competition"}
	competition.UpdateStatus(businesslogic.COMPETITION_STATUS_OPEN_REGISTRATION)

	compRepo := mock_businesslogic.NewMockICompetitionRepository(mockCtrl)
	compRepo.EXPECT().SearchCompetition(gomock.Any()).Return(
		[]businesslogic.Competition{
			competition,
		}, nil)

	err := entry.CreateCompetitionEntry(compRepo, entryRepo)
	assert.NotNil(t, err, "should create duplicate competition entry with error")

	entryRepo.EXPECT().SearchCompetitionEntry(businesslogic.SearchCompetitionEntryCriteria{
		AthleteID:     12,
		CompetitionID: 44,
	}).Return([]businesslogic.CompetitionEntry{}, nil)
	entryRepo.EXPECT().CreateCompetitionEntry(gomock.Any()).Return(nil)
	compRepo.EXPECT().SearchCompetition(gomock.Any()).Return(
		[]businesslogic.Competition{
			competition,
		}, nil)
	err = entry.CreateCompetitionEntry(compRepo, entryRepo)
	assert.Nil(t, err, "should create new competition entry without error")
}
