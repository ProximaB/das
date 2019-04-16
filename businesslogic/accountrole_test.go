package businesslogic_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGrantRole_HasRole(t *testing.T) {
	// user account
	user := businesslogic.Account{ID: 1}
	roles := []businesslogic.AccountRole{
		businesslogic.NewAccountRole(user, businesslogic.AccountTypeAthlete),
		businesslogic.NewAccountRole(user, businesslogic.AccountTypeDeckCaptain),
	}
	user.SetRoles(roles)

	assert.True(t, user.HasRole(businesslogic.AccountTypeAthlete), "user should have Athlete role assigned")
	assert.True(t, user.HasRole(businesslogic.AccountTypeDeckCaptain), "user should have Athlete role assigned")
	assert.False(t, user.HasRole(businesslogic.AccountTypeAdjudicator), "user should NOT have Adjudicator role assigned")
}

func TestRoleProvisionService_CreateRoleApplication(t *testing.T) {

}
