package businesslogic_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePartnershipRoundEntry_NoError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mock_businesslogic.NewMockIPartnershipRoundEntryRepository(mockCtrl)

	repo.EXPECT().SearchPartnershipRoundEntry(gomock.Any()).Return([]businesslogic.PartnershipRoundEntry{}, nil)
	repo.EXPECT().CreatePartnershipRoundEntry(gomock.Any()).Return(nil)

	result := businesslogic.CreatePartnershipRoundEntry(&businesslogic.PartnershipRoundEntry{
		PartnershipID: 13,
		RoundEntry:    businesslogic.RoundEntry{RoundID: 33},
	}, repo)

	assert.Nil(t, result, "should not return error when search and create entry does not return error")
}
