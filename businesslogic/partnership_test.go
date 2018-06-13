package businesslogic_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount_GetAllPartnerships(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock_businesslogic.NewMockIPartnershipRepository(mockCtrl)
	mockRepo.EXPECT().SearchPartnership(businesslogic.SearchPartnershipCriteria{
		LeadID: 9,
	}).Return([]businesslogic.Partnership{
		{PartnershipID: 1, LeadID: 9, FollowID: 8},
		{PartnershipID: 2, LeadID: 9, FollowID: 3},
	}, nil)
	mockRepo.EXPECT().SearchPartnership(businesslogic.SearchPartnershipCriteria{
		FollowID: 9,
	}).Return([]businesslogic.Partnership{
		{PartnershipID: 7, LeadID: 33, FollowID: 9},
	}, nil)

	athlete := businesslogic.Account{ID: 9}
	partnerships, err := athlete.GetAllPartnerships(mockRepo)

	assert.Nil(t, err, "should get all partnerships without error")
	assert.EqualValues(t, 3, len(partnerships), "should get all partnerships as lead and follow")
}
