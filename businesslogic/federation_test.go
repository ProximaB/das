package businesslogic_test

import (
	"errors"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
