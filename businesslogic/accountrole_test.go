package businesslogic_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/accountdal"
	"github.com/DancesportSoftware/das/mock/businesslogic"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccountRoleProvisionService(t *testing.T) {
	var testAccountRepo = accountdal.PostgresAccountRepository{}
	var testAccountRoleRepo = accountdal.PostgresAccountRoleRepository{}
	service := businesslogic.NewAccountRoleProvisionService(testAccountRepo, testAccountRoleRepo)
	assert.NotNil(t, service, "should not create a null service")
}

func TestGrantRole_HasRole(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// user account
	userRequestingRole := businesslogic.Account{ID: 1}
	roles := []businesslogic.AccountRole{
		{ID: 1, AccountID: 1, AccountTypeID: businesslogic.AccountTypeOrganizer},
	}
	userRequestingRole.SetRoles(roles)

	// admin account
	userAdmin := businesslogic.Account{ID: 5}
	roles = []businesslogic.AccountRole{
		{ID: 1, AccountID: 5, AccountTypeID: businesslogic.AccountTypeAdministrator},
	}
	userAdmin.SetRoles(roles)

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedAccountRoleRepo := mock_businesslogic.NewMockIAccountRoleRepository(mockCtrl)
	mockedAccountRoleProvisionService := businesslogic.NewAccountRoleProvisionService(
		mockedAccountRepo,
		mockedAccountRoleRepo,
	)

	assert.Error(t, mockedAccountRoleProvisionService.GrantRole(
		userAdmin,
		userRequestingRole,
		businesslogic.AccountTypeOrganizer,
	))
}

func TestGrantRole_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	// user account
	userRequestingRole := businesslogic.Account{ID: 1}
	roles := []businesslogic.AccountRole{
		{ID: 1, AccountID: 1, AccountTypeID: businesslogic.AccountTypeOrganizer},
	}
	userRequestingRole.SetRoles(roles)

	// admin account
	userAdmin := businesslogic.Account{ID: 5}
	roles = []businesslogic.AccountRole{
		{ID: 1, AccountID: 5, AccountTypeID: businesslogic.AccountTypeAdministrator},
	}
	userAdmin.SetRoles(roles)

	mockedAccountRepo := mock_businesslogic.NewMockIAccountRepository(mockCtrl)
	mockedAccountRoleRepo := mock_businesslogic.NewMockIAccountRoleRepository(mockCtrl)
	/*mockedAccountRoleRepo.EXPECT().CreateAccountRole(&businesslogic.AccountRole{
		ID:              0,
		AccountID:       1,
		AccountTypeID:   businesslogic.AccountTypeScrutineer,
		CreateUserID:    5,
		DateTimeCreated: time.Now(),
		UpdateUserID:    5,
		DateTimeUpdated: time.Now(),
	}).Return(nil)*/
	mockedAccountRoleRepo.EXPECT().CreateAccountRole(gomock.Any()).Return(nil)

	mockedAccountRoleProvisionService := businesslogic.NewAccountRoleProvisionService(
		mockedAccountRepo,
		mockedAccountRoleRepo,
	)

	assert.Nil(t, mockedAccountRoleProvisionService.GrantRole(
		userAdmin,
		userRequestingRole,
		businesslogic.AccountTypeScrutineer,
	))
}
