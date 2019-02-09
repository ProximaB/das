package businesslogic_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCountry_GetStates(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_businesslogic.NewMockIStateRepository(mockCtrl)
	mockRepo.EXPECT().SearchState(businesslogic.SearchStateCriteria{}).Return([]businesslogic.State{
		{ID: 1, Name: "Alaska", Abbreviation: "AK"},
		{ID: 2, Name: "Michigan", Abbreviation: "MI"},
	}, nil)

	country := businesslogic.Country{}

	states, err := country.GetStates(mockRepo)
	assert.Nil(t, err, "search states of a Country should not return errors")
	assert.EqualValues(t, len(states), 2, "should return all states when search with empty criteria")
}

func TestCountry_GetFederations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFederationRepo := mock_businesslogic.NewMockIFederationRepository(ctrl)
	mockFederationRepo.EXPECT().SearchFederation(businesslogic.SearchFederationCriteria{}).Return(
		[]businesslogic.Federation{
			{ID: 1, Name: "WDSF"},
			{ID: 2, Name: "WDC"},
		}, nil,
	)

	country := businesslogic.Country{}
	federations, err := country.GetFederations(mockFederationRepo)

	assert.Nil(t, err)
	assert.EqualValues(t, len(federations), 2, "search federation with empty criteria should return all federations")
}
