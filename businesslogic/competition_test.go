package businesslogic_test

import (
	"errors"
	"testing"
	"time"

	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateCompetition(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	competitionRepo := mock_businesslogic.NewMockICompetitionRepository(mockCtrl)
	provisionRepo := mock_businesslogic.NewMockIOrganizerProvisionRepository(mockCtrl)
	provisionHistoryRepo := mock_businesslogic.NewMockIOrganizerProvisionHistoryRepository(mockCtrl)

	comp := businesslogic.Competition{
		Name:          "Intergalactic Competition",
		Website:       "http://www.example.com",
		FederationID:  1,
		StartDateTime: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+2, 1, 1, 1, 1, time.UTC),
		EndDateTime:   time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+4, 1, 1, 1, 1, time.UTC),
		ContactName:   "James Bond",
		ContactEmail:  "james.bond@email.com",
		ContactPhone:  "2290092292",
		City:          businesslogic.City{ID: 26},
		State:         businesslogic.State{ID: 17},
		Country:       businesslogic.Country{ID: 19},
		CreateUserID:  1,
		UpdateUserID:  1,
	}

	competitionRepo.EXPECT().CreateCompetition(&comp).Return(nil)
	provisionRepo.EXPECT().SearchOrganizerProvision(businesslogic.SearchOrganizerProvisionCriteria{
		OrganizerID: 1,
	}).Return([]businesslogic.OrganizerProvision{
		{ID: 3, OrganizerRoleID: 1, Available: 3, Hosted: 7},
	}, nil)
	provisionRepo.EXPECT().UpdateOrganizerProvision(gomock.Any()).Return(nil)
	provisionHistoryRepo.EXPECT().CreateOrganizerProvisionHistory(gomock.Any()).Return(nil)

	err := businesslogic.CreateCompetition(comp, competitionRepo, provisionRepo, provisionHistoryRepo)
	assert.Nil(t, err, "should create competition if competition data is correct and organizer has sufficient provision")
}

func TestCompetition_UpdateStatus(t *testing.T) {
	comp := businesslogic.Competition{}

	err_1 := comp.UpdateStatus(businesslogic.CompetitionStatusPreRegistration)
	assert.Nil(t, err_1, "change the status of newly instantiated competition should not result in error")

	err_2 := comp.UpdateStatus(businesslogic.CompetitionStatusInProgress)
	assert.Nil(t, err_2, "change the status of competition from pre-registration to in-progress should not result in error")

	err_3 := comp.UpdateStatus(businesslogic.CompetitionStatusOpenRegistration)
	assert.NotNil(t, err_3, "cannot revert competition status from in-progress to open-registration")

}

func TestCompetition_GetStatus(t *testing.T) {
	comp := businesslogic.Competition{}
	comp.UpdateStatus(businesslogic.CompetitionStatusClosedRegistration)

	assert.Equal(t, businesslogic.CompetitionStatusClosedRegistration, comp.GetStatus())
}

// GetCompetitionByID test helpers
type getCompetitionByIDResult struct {
	comp businesslogic.Competition
	err  error
}

func twoValueReturnHandler(c businesslogic.Competition, e error) getCompetitionByIDResult {
	result := getCompetitionByIDResult{comp: c, err: e}

	return result
}

func getCompetitionByIDMockHandler(m *gomock.Controller, id int, r []businesslogic.Competition,
	e error) businesslogic.ICompetitionRepository {
	searchComp := businesslogic.SearchCompetitionCriteria{ID: id}
	competitionRepo := mock_businesslogic.NewMockICompetitionRepository(m)
	competitionRepo.EXPECT().SearchCompetition(searchComp).Return(r, e).MaxTimes(2)

	return competitionRepo
}

func getCompetitionByIDAssertNilHandler(t *testing.T, competitionRepo businesslogic.ICompetitionRepository) {
	assert.Equal(
		t,
		twoValueReturnHandler(businesslogic.Competition{}, errors.New("Return an error")).comp,
		twoValueReturnHandler(businesslogic.GetCompetitionByID(2, competitionRepo)).comp,
	)
	assert.Nil(t, twoValueReturnHandler(businesslogic.GetCompetitionByID(2, competitionRepo)).err)
}

// GetCompetitionByID tests
func TestCompetition_GetCompetitionByID_ErrorNotNil(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	competitionRepo := getCompetitionByIDMockHandler(
		mockCtrl,
		2,
		[]businesslogic.Competition{},
		errors.New("Return empty competitions and a database error."),
	)
	assert.Equal(
		t,
		twoValueReturnHandler(businesslogic.Competition{}, errors.New("Return an error")).comp,
		twoValueReturnHandler(businesslogic.GetCompetitionByID(2, competitionRepo)).comp,
	)
	assert.Error(t, twoValueReturnHandler(businesslogic.GetCompetitionByID(2, competitionRepo)).err)

}

func TestCompetition_GetCompetitionByID_SearchResultNil(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	competitionRepo := getCompetitionByIDMockHandler(mockCtrl, 2, nil, nil)
	getCompetitionByIDAssertNilHandler(t, competitionRepo)
}

func TestCompetition_GetCompetitionByID_SearchResultLengthNotOne(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	competitionRepo := getCompetitionByIDMockHandler(mockCtrl, 2, make([]businesslogic.Competition, 2), nil)
	getCompetitionByIDAssertNilHandler(t, competitionRepo)
}

func TestCompetition_GetCompetitionByID_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	competitionRepo := getCompetitionByIDMockHandler(mockCtrl, 2, []businesslogic.Competition{}, nil)
	getCompetitionByIDAssertNilHandler(t, competitionRepo)
}

// UpdateCompetition test helper functions
func updateCompetitionMockHandler(m *gomock.Controller, userID int, compID int, status int,
	name string, website string, startDate time.Time, endDate time.Time,
	comp []businesslogic.Competition, err error) (businesslogic.Account,
	businesslogic.OrganizerUpdateCompetition, *mock_businesslogic.MockICompetitionRepository) {

	user := businesslogic.Account{ID: userID}
	competition := businesslogic.OrganizerUpdateCompetition{
		CompetitionID: compID,
		Status:        status,
		Name:          name,
		Website:       website,
		StartDate:     startDate,
		EndDate:       endDate,
	}
	searchComp := businesslogic.SearchCompetitionCriteria{ID: competition.CompetitionID}
	competitionRepo := mock_businesslogic.NewMockICompetitionRepository(m)
	competitionRepo.EXPECT().SearchCompetition(searchComp).Return(comp, err)

	return user, competition, competitionRepo
}

// UpdateCompetition tests
func TestCompetition_UpdateCompetition_ErrorNotNil(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// test parameters
	start := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 18, 9, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 19, 22, 0, 0, 0, time.UTC)

	// initialize mocks
	user, competition, competitionRepo := updateCompetitionMockHandler(mockCtrl, 2, 2,
		businesslogic.CompetitionStatusInProgress, "The Great American Ball", "www.tgab.com",
		start, end, []businesslogic.Competition{}, errors.New("Return an error"))

	assert.Error(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}

func TestCompetition_UpdateCompetition_SearchResultNil(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// test parameters
	start := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 18, 9, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 19, 22, 0, 0, 0, time.UTC)

	// initialize mocks
	user, competition, competitionRepo := updateCompetitionMockHandler(mockCtrl, 2, 2,
		businesslogic.CompetitionStatusInProgress, "The Great American Ball", "www.tgab.com",
		start, end, nil, errors.New("cannot find this competition"))

	assert.Error(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}

func TestCompetition_UpdateCompetition_SearchResultNotOne(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// test parameters
	start := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 18, 9, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 19, 22, 0, 0, 0, time.UTC)

	// initialize mocks
	user, competition, competitionRepo := updateCompetitionMockHandler(mockCtrl, 2, 2,
		businesslogic.CompetitionStatusInProgress, "The Great American Ball", "www.tgab.com",
		start, end, make([]businesslogic.Competition, 2), errors.New("cannot find this competition"))

	assert.Error(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}

func TestCompetition_UpdateCompetition_CompIDEqualsZero(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// test parameters
	start := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 18, 9, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 19, 22, 0, 0, 0, time.UTC)
	comp := []businesslogic.Competition{{ID: 0}}

	// initialize mocks
	user, competition, competitionRepo := updateCompetitionMockHandler(mockCtrl, 2, 2,
		businesslogic.CompetitionStatusInProgress, "The Great American Ball", "www.tgab.com",
		start, end, comp, errors.New("cannot find this competition"))

	assert.Error(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}

func TestCompetition_UpdateCompetition_ValidateUpdateCompetition_CreateUserIDNoMatch(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// test parameters
	start := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 18, 9, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 19, 22, 0, 0, 0, time.UTC)
	comp := []businesslogic.Competition{{ID: 2, CreateUserID: 3}}

	// initialize mocks
	user, competition, competitionRepo := updateCompetitionMockHandler(mockCtrl, 2, 2,
		businesslogic.CompetitionStatusInProgress, "The Great American Ball", "www.tgab.com",
		start, end, comp, nil)

	assert.Error(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}

func TestCompetition_UpdateCompetition_ValidateUpdateCompetition_UpdateStatusInvalid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// test parameters
	start := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 18, 9, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 19, 22, 0, 0, 0, time.UTC)
	comp := []businesslogic.Competition{{ID: 2, CreateUserID: 2}}
	comp[0].UpdateStatus(businesslogic.CompetitionStatusProcessing)

	// initialize mocks
	user, competition, competitionRepo := updateCompetitionMockHandler(mockCtrl, 2, 2,
		businesslogic.CompetitionStatusInProgress, "The Great American Ball", "www.tgab.com",
		start, end, comp, nil)

	assert.Error(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}

func TestCompetition_UpdateCompetition_ValidateUpdateCompetition_StatusClosed(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// test parameters
	start := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 18, 9, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 19, 22, 0, 0, 0, time.UTC)
	comp := []businesslogic.Competition{{ID: 2, CreateUserID: 2}}
	comp[0].UpdateStatus(businesslogic.CompetitionStatusClosed)

	// initialize mocks
	user, competition, competitionRepo := updateCompetitionMockHandler(mockCtrl, 2, 2,
		businesslogic.CompetitionStatusCancelled, "The Great American Ball", "www.tgab.com",
		start, end, comp, nil)

	assert.Error(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}

func TestCompetition_UpdateCompetition_ValidateUpdateCompetition_NameTooShort(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// test parameters
	start := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 18, 9, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 19, 22, 0, 0, 0, time.UTC)
	comp := []businesslogic.Competition{{ID: 2, CreateUserID: 2}}
	comp[0].UpdateStatus(businesslogic.CompetitionStatusClosedRegistration)

	// initialize mocks
	user, competition, competitionRepo := updateCompetitionMockHandler(mockCtrl, 2, 2,
		businesslogic.CompetitionStatusInProgress, "", "www.tgab.com",
		start, end, comp, nil)

	assert.Error(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}

func TestCompetition_UpdateCompetition_ValidateUpdateCompetition_WebAddressInvalid(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// test parameters
	start := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 18, 9, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 19, 22, 0, 0, 0, time.UTC)
	comp := []businesslogic.Competition{{ID: 2, CreateUserID: 2}}
	comp[0].UpdateStatus(businesslogic.CompetitionStatusClosedRegistration)

	// initialize mocks
	user, competition, competitionRepo := updateCompetitionMockHandler(mockCtrl, 2, 2,
		businesslogic.CompetitionStatusInProgress, "The Great American Ball", "",
		start, end, comp, nil)

	assert.Error(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}

func TestCompetition_UpdateCompetition_ValidateUpdateCompetition_StartDateAfterEndDate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// test parameters
	start := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 20, 9, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 19, 22, 0, 0, 0, time.UTC)
	comp := []businesslogic.Competition{{ID: 2, CreateUserID: 2}}
	comp[0].UpdateStatus(businesslogic.CompetitionStatusClosedRegistration)

	// initialize mocks
	user, competition, competitionRepo := updateCompetitionMockHandler(mockCtrl, 2, 2,
		businesslogic.CompetitionStatusInProgress, "The Great American Ball", "www.tgab.com",
		start, end, comp, nil)

	assert.Error(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}

func TestCompetition_UpdateCompetition_ValidateUpdateCompetition_StartDatePassed(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// test parameters
	start := time.Date(time.Now().AddDate(-1, 0, 0).Year(), time.Now().AddDate(0, 6, 0).Month(), 18, 9, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 19, 22, 0, 0, 0, time.UTC)
	comp := []businesslogic.Competition{{ID: 2, CreateUserID: 2}}
	comp[0].UpdateStatus(businesslogic.CompetitionStatusClosedRegistration)

	// initialize mocks
	user, competition, competitionRepo := updateCompetitionMockHandler(mockCtrl, 2, 2,
		businesslogic.CompetitionStatusInProgress, "The Great American Ball", "www.tgab.com",
		start, end, comp, nil)

	assert.Error(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}

func TestCompetition_UpdateCompetition_ValidateUpdateCompetition_StartDateYearLater(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// test parameters
	start := time.Date(time.Now().AddDate(1, 0, 0).Year(), time.Now().AddDate(0, 6, 0).Month(), 18, 9, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().AddDate(1, 0, 0).Year(), time.Now().AddDate(0, 6, 0).Month(), 19, 22, 0, 0, 0, time.UTC)
	comp := []businesslogic.Competition{{ID: 2, CreateUserID: 2}}
	comp[0].UpdateStatus(businesslogic.CompetitionStatusClosedRegistration)

	// initialize mocks
	user, competition, competitionRepo := updateCompetitionMockHandler(mockCtrl, 2, 2,
		businesslogic.CompetitionStatusInProgress, "The Great American Ball", "www.tgab.com",
		start, end, comp, nil)

	assert.Error(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}

func TestCompetition_UpdateCompetition_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// test parameters
	start := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 18, 9, 0, 0, 0, time.UTC)
	end := time.Date(time.Now().Year(), time.Now().AddDate(0, 6, 0).Month(), 19, 22, 0, 0, 0, time.UTC)
	comp := []businesslogic.Competition{{ID: 2, CreateUserID: 2}}
	comp[0].UpdateStatus(businesslogic.CompetitionStatusClosedRegistration)

	// initialize mocks
	user, competition, competitionRepo := updateCompetitionMockHandler(mockCtrl, 2, 2,
		businesslogic.CompetitionStatusInProgress, "The Great American Ball", "www.tgab.com",
		start, end, comp, nil)

	// add expected return for success
	competitionRepo.EXPECT().UpdateCompetition(comp[0]).Return(nil)

	assert.Nil(t, businesslogic.UpdateCompetition(&user, competition, competitionRepo))
}
