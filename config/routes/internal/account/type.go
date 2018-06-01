package account

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/account"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiAccountTypeEndpoint = "/api/account/type"

var AccountTypeController = util.DasController{
	Description:  "Get all account types in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiAccountTypeEndpoint,
	Handler:      accountTypeServer.GetAccountTypeHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var accountTypeServer = account.AccountTypeServer{
	database.AccountTypeRepository,
}
