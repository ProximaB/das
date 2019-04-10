package businesslogic_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccount_GetBlacklistedAccounts(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	user_17 := businesslogic.Account{ID: 17}

	accountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	blacklistRepo := mock_businesslogic.NewMockIPartnershipRequestBlacklistRepository(mockCtrl)

	blacklistRepo.EXPECT().SearchPartnershipRequestBlacklist(businesslogic.SearchPartnershipRequestBlacklistCriteria{ReporterID: 17, Whitelisted: false}).Return([]businesslogic.PartnershipRequestBlacklistEntry{
		{ID: 21, BlockedUser: businesslogic.Account{ID: 33}, Reporter: businesslogic.Account{ID: 17}},
	}, nil)
	accountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{ID: 33}).Return([]businesslogic.Account{
		{ID: 33, FirstName: "Sharp", LastName: "Shark"},
	}, nil)
	blacklist, err := user_17.GetBlacklistedAccounts(accountRepo, blacklistRepo)

	assert.Nil(t, err, "should retrieve blacklisted accounts without error")
	assert.EqualValues(t, 1, len(blacklist), "should get all blacklisted account without error")
}
