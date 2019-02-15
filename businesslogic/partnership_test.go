// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package businesslogic_test

import (
	"errors"
	"testing"

	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
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

// HasAthlete tests
func TestPartnership_HasAthlete_False(t *testing.T) {
	leadAccount := businesslogic.Account{ID: 7}
	followAccount := businesslogic.Account{ID: 6}
	partnership := businesslogic.Partnership{
		Lead:   leadAccount,
		Follow: followAccount,
	}

	assert.False(t, partnership.HasAthlete(5))
}

func TestPartnership_HasAthlete_True(t *testing.T) {
	leadAccount := businesslogic.Account{ID: 7}
	followAccount := businesslogic.Account{ID: 6}
	partnership := businesslogic.Partnership{
		Lead:   leadAccount,
		Follow: followAccount,
	}

	assert.True(t, partnership.HasAthlete(6))
}

// GetPartnershipByID helper functions
type getPartnershipByIDResult struct {
	partnership businesslogic.Partnership
	err         error
}

func partnershipTwoValueReturnHandler(p businesslogic.Partnership, e error) getPartnershipByIDResult {
	result := getPartnershipByIDResult{partnership: p, err: e}

	return result
}

func getPartnershipByIDMockHandler(m *gomock.Controller, partnerID int,
	partnership []businesslogic.Partnership, err error) businesslogic.IPartnershipRepository {

	partnershipRepo := mock_businesslogic.NewMockIPartnershipRepository(m)
	partnershipRepo.EXPECT().SearchPartnership(businesslogic.SearchPartnershipCriteria{
		PartnershipID: partnerID,
	}).Return(partnership, err).MaxTimes(2)

	return partnershipRepo
}

func getPartnershipByIDAssertEqualNilHandler(t *testing.T, partnerID int,
	partnershipRepo businesslogic.IPartnershipRepository) {

	assert.Equal(
		t,
		partnershipTwoValueReturnHandler(businesslogic.Partnership{}, errors.New("Return an error")).partnership,
		partnershipTwoValueReturnHandler(businesslogic.GetPartnershipByID(partnerID, partnershipRepo)).partnership,
	)
	assert.Nil(
		t,
		partnershipTwoValueReturnHandler(businesslogic.GetPartnershipByID(5, partnershipRepo)).err,
	)
}

// GetPartnershipByID tests
func TestPartnership_GetPartnershipByID_SearchError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	partnershipRepo := getPartnershipByIDMockHandler(mockCtrl, 5, nil, errors.New("Return an error"))
	assert.Equal(
		t,
		partnershipTwoValueReturnHandler(businesslogic.Partnership{}, errors.New("Return an error")).partnership,
		partnershipTwoValueReturnHandler(businesslogic.GetPartnershipByID(5, partnershipRepo)).partnership,
	)
	assert.Error(
		t,
		partnershipTwoValueReturnHandler(businesslogic.GetPartnershipByID(5, partnershipRepo)).err,
	)
}

func TestPartnership_GetPartnershipByID_SearchResultNil(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	partnershipRepo := getPartnershipByIDMockHandler(mockCtrl, 5, nil, nil)
	getPartnershipByIDAssertEqualNilHandler(t, 5, partnershipRepo)
}

func TestPartnership_GetPartnershipByID_SearchResultLengthNotOne(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	partnershipRepo := getPartnershipByIDMockHandler(mockCtrl, 5, make([]businesslogic.Partnership, 2), nil)
	getPartnershipByIDAssertEqualNilHandler(t, 5, partnershipRepo)
}

func TestPartnership_GetPartnershipByID_SearchResultSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	partnershipRepo := getPartnershipByIDMockHandler(mockCtrl, 5, []businesslogic.Partnership{}, nil)
	getPartnershipByIDAssertEqualNilHandler(t, 5, partnershipRepo)
}

func TestPartnership_MustGetPartnershipByID_SearchError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	partnershipRepo := mock_businesslogic.NewMockIPartnershipRepository(mockCtrl)
	partnershipRepo.EXPECT().SearchPartnership(businesslogic.SearchPartnershipCriteria{
		PartnershipID: 6,
	}).Return(nil, errors.New("Return an error"))
	assert.Panics(t, func() { businesslogic.MustGetPartnershipByID(6, partnershipRepo) })
}

func TestPartnership_MuGetPartnershpiByID_SearchResultLengthNotOne(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	partnershipRepo := mock_businesslogic.NewMockIPartnershipRepository(mockCtrl)
	partnershipRepo.EXPECT().SearchPartnership(businesslogic.SearchPartnershipCriteria{
		PartnershipID: 6,
	}).Return(make([]businesslogic.Partnership, 2), nil)
	assert.Panics(t, func() { businesslogic.MustGetPartnershipByID(6, partnershipRepo) })
}
