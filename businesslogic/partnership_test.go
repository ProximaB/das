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
		{ID: 1, Lead: businesslogic.Account{ID: 9}, Follow: businesslogic.Account{ID: 8}},
		{ID: 2, Lead: businesslogic.Account{ID: 9}, Follow: businesslogic.Account{ID: 3}},
	}, nil)
	mockRepo.EXPECT().SearchPartnership(businesslogic.SearchPartnershipCriteria{
		FollowID: 9,
	}).Return([]businesslogic.Partnership{
		{ID: 7, Lead: businesslogic.Account{ID: 33}, Follow: businesslogic.Account{ID: 9}},
	}, nil)

	athlete := businesslogic.Account{ID: 9}
	partnerships, err := athlete.GetAllPartnerships(mockRepo)

	assert.Nil(t, err, "should get all partnerships without error")
	assert.EqualValues(t, 3, len(partnerships), "should get all partnerships as lead and follow")
}

func TestCreatePartnershipRequest(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	request := businesslogic.PartnershipRequest{
		SenderID:      12,
		RecipientID:   33,
		SenderRole:    businesslogic.PartnershipRoleLead,
		RecipientRole: businesslogic.PartnershipRoleFollow,
		Message:       "Hi, can you add me please?",
		Status:        businesslogic.PartnershipRequestStatusPending,
	}

	accountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	requestRepo := mock_businesslogic.NewMockIPartnershipRequestRepository(mockCtrl)
	partnershipRepo := mock_businesslogic.NewMockIPartnershipRepository(mockCtrl)
	blacklistRepo := mock_businesslogic.NewMockIPartnershipRequestBlacklistRepository(mockCtrl)

	rolesOfLeadAccount := []businesslogic.AccountRole{
		{ID: 1, AccountID: 1, AccountTypeID: businesslogic.AccountTypeOrganizer},
		{ID: 2, AccountID: 1, AccountTypeID: businesslogic.AccountTypeAthlete},
	}
	leadAccount := businesslogic.Account{
		ID: 1,
	}
	leadAccount.SetRoles(rolesOfLeadAccount)

	rolesOfFollowAccount := []businesslogic.AccountRole{
		{ID: 3, AccountID: 2, AccountTypeID: businesslogic.AccountTypeAdjudicator},
		{ID: 4, AccountID: 2, AccountTypeID: businesslogic.AccountTypeDeckCaptain},
		{ID: 5, AccountID: 2, AccountTypeID: businesslogic.AccountTypeAthlete},
	}
	followAccount := businesslogic.Account{
		ID: 2,
	}
	followAccount.SetRoles(rolesOfFollowAccount)

	// specify behaviors
	accountRepo.EXPECT().SearchAccount(gomock.Any()).Return([]businesslogic.Account{
		leadAccount,
	}, nil)
	accountRepo.EXPECT().SearchAccount(gomock.Any()).Return([]businesslogic.Account{
		followAccount,
	}, nil)
	blacklistRepo.EXPECT().SearchPartnershipRequestBlacklist(gomock.Any()).Return([]businesslogic.PartnershipRequestBlacklistEntry{}, nil)
	partnershipRepo.EXPECT().SearchPartnership(gomock.Any()).Return([]businesslogic.Partnership{}, nil)
	requestRepo.EXPECT().SearchPartnershipRequest(gomock.Any()).Return([]businesslogic.PartnershipRequest{}, nil)
	requestRepo.EXPECT().CreatePartnershipRequest(gomock.Any()).Return(nil)

	err := businesslogic.CreatePartnershipRequest(request, partnershipRepo, requestRepo, accountRepo, blacklistRepo)
	assert.Nil(t, err, "should not throw an error if every step is working correctly")
}
