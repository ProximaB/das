package businesslogic_test

import (
	"errors"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreatePartnershipEventEntry_NoError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mock_businesslogic.NewMockIPartnershipEventEntryRepository(mockCtrl)

	// defines expected behavior
	repo.EXPECT().SearchPartnershipEventEntry(gomock.Any()).Return([]businesslogic.PartnershipEventEntry{}, nil)
	repo.EXPECT().CreatePartnershipEventEntry(gomock.Any()).Return(nil)

	entry := businesslogic.PartnershipEventEntry{
		Couple: businesslogic.Partnership{ID: 3},
		Event: businesslogic.Event{
			ID: 12,
		},
	}

	result := businesslogic.CreatePartnershipEventEntry(entry, repo)
	assert.Nil(t, result, "should not return error if no entry exists")
}

func TestCreatePartnershipEventEntry_WithError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mock_businesslogic.NewMockIPartnershipEventEntryRepository(mockCtrl)

	// defines expected behavior
	repo.EXPECT().SearchPartnershipEventEntry(gomock.Any()).Return([]businesslogic.PartnershipEventEntry{}, errors.New("a random error"))

	entry := businesslogic.PartnershipEventEntry{
		Couple: businesslogic.Partnership{ID: 3},
		Event: businesslogic.Event{
			ID: 12,
		},
	}

	result := businesslogic.CreatePartnershipEventEntry(entry, repo)
	assert.NotNil(t, result, "should return error if searching process has error")
}

func TestCreatePartnershipEventEntry_ExistingEntry(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	repo := mock_businesslogic.NewMockIPartnershipEventEntryRepository(mockCtrl)

	// defines expected behavior
	repo.EXPECT().SearchPartnershipEventEntry(gomock.Any()).Return([]businesslogic.PartnershipEventEntry{
		{
			Couple: businesslogic.Partnership{ID: 3},
			Event:  businesslogic.Event{ID: 12},
		},
	}, nil)

	entry := businesslogic.PartnershipEventEntry{
		Couple: businesslogic.Partnership{ID: 3},
		Event:  businesslogic.Event{ID: 12},
	}

	result := businesslogic.CreatePartnershipEventEntry(entry, repo)
	assert.NotNil(t, result, "should return error if a matching entry is found")
}
