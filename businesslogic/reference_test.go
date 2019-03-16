package businesslogic_test

import (
	"errors"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestState_GetCities(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_businesslogic.NewMockICityRepository(mockCtrl)

	// behavior 1
	mockRepo.EXPECT().SearchCity(businesslogic.SearchCityCriteria{StateID: 1}).Return([]businesslogic.City{
		{ID: 1, Name: "City of ID 1", StateID: 1},
		{ID: 2, Name: "City of ID 2", StateID: 1},
	}, nil)

	// behavior 2
	mockRepo.EXPECT().SearchCity(businesslogic.SearchCityCriteria{StateID: 2}).Return(nil,
		errors.New("state does not exist"))

	state_1 := businesslogic.State{ID: 1}
	cities_1, err_1 := state_1.GetCities(mockRepo)
	assert.EqualValues(t, 2, len(cities_1))
	assert.Nil(t, err_1)

	state_2 := businesslogic.State{ID: 2}
	cities_2, err_2 := state_2.GetCities(mockRepo)
	assert.Nil(t, cities_2)
	assert.NotNil(t, err_2)

	cities_3, err_3 := state_1.GetCities(nil)
	assert.Nil(t, cities_3)
	assert.NotNil(t, err_3)
}

func TestFederation_GetDivisions(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_businesslogic.NewMockIDivisionRepository(mockCtrl)

	// behavior 1
	mockRepo.EXPECT().SearchDivision(businesslogic.SearchDivisionCriteria{FederationID: 1}).Return([]businesslogic.Division{
		{ID: 1, Name: "Correct Division 1", FederationID: 1},
		{ID: 2, Name: "Correct Division 2", FederationID: 2},
	}, nil)

	// behavior 2
	mockRepo.EXPECT().SearchDivision(businesslogic.SearchDivisionCriteria{FederationID: 2}).Return(nil, errors.New("invalid search"))

	federation_1 := businesslogic.Federation{ID: 1}
	federation_2 := businesslogic.Federation{ID: 2}

	result_1, err_1 := federation_1.GetDivisions(mockRepo)
	assert.EqualValues(t, 2, len(result_1))
	assert.Nil(t, err_1)

	result_2, err_2 := federation_2.GetDivisions(mockRepo)
	assert.Nil(t, result_2)
	assert.NotNil(t, err_2)

	result_3, err_3 := federation_1.GetDivisions(nil)
	assert.Nil(t, result_3)
	assert.NotNil(t, err_3)
}
