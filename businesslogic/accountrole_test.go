package businesslogic_test

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/dataaccess/accountdal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAccountRoleProvisionService(t *testing.T) {
	var testAccountRepo = accountdal.PostgresAccountRepository{}
	var testAccountRoleRepo = accountdal.PostgresAccountRoleRepository{}
	service := businesslogic.NewAccountRoleProvisionService(testAccountRepo, testAccountRoleRepo)
	assert.NotNil(t, service, "should not create a null service")
}
