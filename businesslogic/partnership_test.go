package businesslogic_test

import (
	"errors"
	"testing"

	"github.com/DancesportSoftware/das/businesslogic"
	mock_businesslogic "github.com/DancesportSoftware/das/mock/businesslogic"
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

/*  These are helper functions used by the tests in partnership_test.go.
Their general purpose is to allow for the actual tests to follow the DRY principle  */

type getPartnershipByIDResult struct {
	partnership businesslogic.Partnership
	err         error
}

// Most functions from partnership.go return two values, this function returns a struct that allows
// access to either values.
func partnershipTwoValueReturnHandler(p businesslogic.Partnership, e error) getPartnershipByIDResult {
	result := getPartnershipByIDResult{partnership: p, err: e}

	return result
}

// Handles the generation of general mocks to prevent reapating of code.  Since the mock is tested
// twice, MaxTimes(n) is used to allow the expected results to be called twice
func getPartnershipByIDMockHandler(m *gomock.Controller, partnerID int,
	partnership []businesslogic.Partnership, err error) businesslogic.IPartnershipRepository {

	partnershipRepo := mock_businesslogic.NewMockIPartnershipRepository(m)
	partnershipRepo.EXPECT().SearchPartnership(businesslogic.SearchPartnershipCriteria{
		PartnershipID: partnerID,
	}).Return(partnership, err).MaxTimes(2)

	return partnershipRepo
}

// Handles the generation of general Equal and Nil assertions
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

/*  GetPartnershipByID tests  */

// Tests the search result returning an error
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

// Tests the search result returning Nil
func TestPartnership_GetPartnershipByID_SearchResultNil(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	partnershipRepo := getPartnershipByIDMockHandler(mockCtrl, 5, nil, nil)
	getPartnershipByIDAssertEqualNilHandler(t, 5, partnershipRepo)
}

// Tests the search result returning a length not equal to one
func TestPartnership_GetPartnershipByID_SearchResultLengthNotOne(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	partnershipRepo := getPartnershipByIDMockHandler(mockCtrl, 5, make([]businesslogic.Partnership, 2), nil)
	getPartnershipByIDAssertEqualNilHandler(t, 5, partnershipRepo)
}

// Tests the search result returning a success
func TestPartnership_GetPartnershipByID_SearchResultSuccess(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	partnershipRepo := getPartnershipByIDMockHandler(mockCtrl, 5, []businesslogic.Partnership{}, nil)
	getPartnershipByIDAssertEqualNilHandler(t, 5, partnershipRepo)
}

/*  MustGetPartnershipByID tests, only three paths, no helper functions  */

// Tests the search result returning an error
func TestPartnership_MustGetPartnershipByID_SearchError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	partnershipRepo := mock_businesslogic.NewMockIPartnershipRepository(mockCtrl)
	partnershipRepo.EXPECT().SearchPartnership(businesslogic.SearchPartnershipCriteria{
		PartnershipID: 6,
	}).Return(nil, errors.New("Return an error"))
	assert.Panics(t, func() { businesslogic.MustGetPartnershipByID(6, partnershipRepo) })
}

// Tests the search result returning a length not equal to one
func TestPartnership_MustGetPartnershpiByID_SearchResultLengthNotOne(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	partnershipRepo := mock_businesslogic.NewMockIPartnershipRepository(mockCtrl)
	partnershipRepo.EXPECT().SearchPartnership(businesslogic.SearchPartnershipCriteria{
		PartnershipID: 6,
	}).Return(make([]businesslogic.Partnership, 2), nil)
	assert.Panics(t, func() { businesslogic.MustGetPartnershipByID(6, partnershipRepo) })
}

// Tests the search result returning a success
func TestPartnership_MustGetPartnershipByID_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	partnershipRepo := mock_businesslogic.NewMockIPartnershipRepository(mockCtrl)
	partnershipRepo.EXPECT().SearchPartnership(businesslogic.SearchPartnershipCriteria{
		PartnershipID: 6,
	}).Return(make([]businesslogic.Partnership, 1), nil)
	assert.Equal(
		t,
		businesslogic.Partnership{},
		businesslogic.MustGetPartnershipByID(6, partnershipRepo),
	)
}
