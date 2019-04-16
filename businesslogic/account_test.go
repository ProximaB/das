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

var testAthleteAccount = businesslogic.Account{
	FirstName:             "First Name",
	LastName:              "Last Name",
	UserGenderID:          businesslogic.GENDER_MALE,
	DateOfBirth:           time.Date(2017, time.January, 1, 1, 1, 1, 1, time.UTC),
	ToSAccepted:           true,
	PrivacyPolicyAccepted: true,
	Email:                 "test@test.com",
	Phone:                 "1232234442",
	Signature:             "I am a parent",
	ByGuardian:            true,
}

var testOrganizerAccount = businesslogic.Account{
	FirstName:             "Mighty",
	LastName:              "Meerkat",
	UserGenderID:          businesslogic.GENDER_FEMALE,
	DateOfBirth:           time.Date(1997, time.May, 22, 1, 1, 1, 1, time.UTC),
	ToSAccepted:           true,
	PrivacyPolicyAccepted: true,
	Email:                 "mighty.meerkat@email.com",
	Phone:                 "3321231232",
}

func TestGetAccountByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedAccountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{
		ID: 1,
	}).Return(nil, errors.New("should not return an account"))
	mockedAccountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{
		ID: 2,
	}).Return([]businesslogic.Account{
		{
			ID: 2, Email: "newuser@email.com",
		},
	}, nil)

	result := businesslogic.GetAccountByID(1, mockedAccountRepo)
	assert.Equal(t, 0, result.ID)
	assert.Equal(t, "", result.Email)

	result = businesslogic.GetAccountByID(2, mockedAccountRepo)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.ID)
	assert.Equal(t, "newuser@email.com", result.Email)
}

func TestGetAccountByUUID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedAccountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{
		UUID: "abc",
	}).Return(nil, errors.New("should not return an account"))
	mockedAccountRepo.EXPECT().SearchAccount(businesslogic.SearchAccountCriteria{
		UUID: "123",
	}).Return([]businesslogic.Account{
		{
			ID: 2, Email: "newuser@email.com",
		},
	}, nil)

	result := businesslogic.GetAccountByUUID("abc", mockedAccountRepo)
	assert.Equal(t, 0, result.ID)
	assert.Equal(t, "", result.Email)

	result = businesslogic.GetAccountByUUID("123", mockedAccountRepo)
	assert.NotNil(t, result)
	assert.Equal(t, 2, result.ID)
	assert.Equal(t, "newuser@email.com", result.Email)
}

func TestAccount_GetRoles(t *testing.T) {
	rolesOfUserAccount := []businesslogic.AccountRole{
		{ID: 1, AccountID: 1, AccountTypeID: businesslogic.AccountTypeOrganizer},
		{ID: 2, AccountID: 1, AccountTypeID: businesslogic.AccountTypeAthlete},
	}
	userAccount := businesslogic.Account{
		ID: 1,
	}
	userAccount.SetRoles(rolesOfUserAccount)

	assert.Equal(t, 2, len(userAccount.GetRoles()))
	assert.True(t, userAccount.HasRole(businesslogic.AccountTypeOrganizer))
	assert.True(t, userAccount.HasRole(businesslogic.AccountTypeAthlete))
}

func TestAccount_SetRoles(t *testing.T) {
	rolesOfUserAccount := []businesslogic.AccountRole{
		{ID: 1, AccountID: 1, AccountTypeID: businesslogic.AccountTypeOrganizer},
		{ID: 2, AccountID: 1, AccountTypeID: businesslogic.AccountTypeAthlete},
	}
	userAccount := businesslogic.Account{
		ID: 1,
	}
	userAccount.SetRoles(rolesOfUserAccount)

	assert.True(t, userAccount.HasRole(businesslogic.AccountTypeAthlete))
	assert.True(t, userAccount.HasRole(businesslogic.AccountTypeOrganizer))
	assert.False(t, userAccount.HasRole(businesslogic.AccountTypeAdjudicator))
	assert.False(t, userAccount.HasRole(businesslogic.AccountTypeScrutineer))
	assert.False(t, userAccount.HasRole(businesslogic.AccountTypeEmcee))
	assert.False(t, userAccount.HasRole(businesslogic.AccountTypeDeckCaptain))
	assert.False(t, userAccount.HasRole(businesslogic.AccountTypeAdministrator))

	accounts := [3]businesslogic.Account{
		{},
		{},
		{},
	}
	for i := 0; i < len(accounts); i++ {
		accounts[i].SetRoles(rolesOfUserAccount)
	}
	for _, each := range accounts {
		assert.True(t, each.HasRole(businesslogic.AccountTypeAthlete))
	}
}

func TestAccount_MeetMinimalRequirement_NameTooShort(t *testing.T) {
	account := businesslogic.Account{
		FirstName: "A", LastName: "J", Email: "t@e.c", Phone: "1234567890",
	}
	assert.Error(t, account.MeetMinimalRequirement(), "account should be invalid if the name is too short")
}

func TestAccount_MeetMinimalRequirement_NameTooLong(t *testing.T) {
	account := businesslogic.Account{
		FirstName: "CatCatCatCatCatCatCatCatCatCatCatCat",
		LastName:  "DogDogDogDogDogDogDogDogDogDogDogDog",
		Email:     "a@b.c",
		Phone:     "1112223333",
	}
	assert.Error(t, account.MeetMinimalRequirement(), "account should be invalid the name is too long")
}

func TestAccount_MeetMinimalRequirement_InvalidEmail(t *testing.T) {
	account := businesslogic.Account{
		FirstName: "Cat", LastName: "Garfield", Email: ".com", Phone: "1112223333",
	}
	assert.Error(t, account.MeetMinimalRequirement(), "account should be invalid if email is invalid")
}

func TestAccount_MeetMinimalRequirement_InvalidPhone(t *testing.T) {
	account := businesslogic.Account{
		FirstName: "Cat", LastName: "Garfield", Email: "a@b.c", Phone: "123456789",
	}
	assert.Error(t, account.MeetMinimalRequirement(), "account should be invalid if phone number is too short")
}

func TestAccount_MeetMinimalRequirement_ValidAccount(t *testing.T) {
	account := businesslogic.Account{
		FirstName: "Cat", LastName: "Garfield", Email: "rimuru@slime.com", Phone: "1234567890",
	}
	assert.Nil(t, account.MeetMinimalRequirement(), "account should be valid")
}

func TestAccount_FullName(t *testing.T) {
	account := businesslogic.Account{
		FirstName: "Michael",
		LastName:  "Kaplan",
	}
	assert.Equal(t, "Michael Kaplan", account.FullName(), "should generate full name correctly")
}
