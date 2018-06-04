/*
To generate mock objects,
1. Open a terminal
2. Change directory to das root: $GOPATH/src/github.com/DancesportSoftware/das
3. Run command; $ mockgen -source./path/to/source/code > ./mock/path/to/mock/object
4. Use the test below as a template
5. If original file changes, chances are the mock file need to be regenerated as well
*/
package reference_test

import (
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/mock/businesslogic/reference"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountry_GetStates(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_reference.NewMockIStateRepository(mockCtrl)
	mockRepo.EXPECT().SearchState(&reference.SearchStateCriteria{}).Return([]reference.State{
		{ID: 1, Name: "Test"},
	}, nil)

	country := reference.Country{}

	states, err := country.GetStates(mockRepo)
	assert.Nil(t, err, "should not return an error")
	assert.NotZero(t, len(states), "should return some states")
}

func TestCountry_HasState(t *testing.T) {

}
