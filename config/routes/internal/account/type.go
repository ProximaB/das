package account

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/account"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiAccountTypeEndpoint = "/api/account/type"

var accountTypeServer = account.AccountTypeServer{
	database.AccountTypeRepository,
}

var AccountTypeController = util.DasController{
	Name:         "AccountTypeController",
	Description:  "Get all account types in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiAccountTypeEndpoint,
	Handler:      accountTypeServer.GetAccountTypeHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}
