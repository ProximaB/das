/*
To generate mock objects,
1. Open a terminal
2. Change directory to das root: $ cd $GOPATH/src/github.com/DancesportSoftware/das
3. Run command; $ mockgen -source=./businesslogic/referencedal/country.go > ./mock/businesslogic/referencedal/country.go
4. Use the test below as a template
5. If original file changes, chances are the mock file need to be regenerated as well
*/
package referencebll_test

import (
	"testing"

	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/mock/businesslogic/reference"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCountry_GetStates(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_reference.NewMockIStateRepository(mockCtrl)
	mockRepo.EXPECT().SearchState(referencebll.SearchStateCriteria{}).Return([]referencebll.State{
		{ID: 1, Name: "Alaska", Abbreviation: "AK"},
		{ID: 2, Name: "Michigan", Abbreviation: "MI"},
	}, nil)

	country := referencebll.Country{}

	states, err := country.GetStates(mockRepo)
	assert.Nil(t, err, "search states of a Country should not return errors")
	assert.EqualValues(t, len(states), 2, "should return all states when search with empty criteria")
}

func TestCountry_GetFederations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFederationRepo := mock_reference.NewMockIFederationRepository(ctrl)
	mockFederationRepo.EXPECT().SearchFederation(referencebll.SearchFederationCriteria{}).Return(
		[]referencebll.Federation{
			{ID: 1, Name: "WDSF"},
			{ID: 2, Name: "WDC"},
		}, nil,
	)

	country := referencebll.Country{}
	federations, err := country.GetFederations(mockFederationRepo)

	assert.Nil(t, err)
	assert.EqualValues(t, len(federations), 2, "search federation with empty criteria should return all federations")
}
