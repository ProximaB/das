package account

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/controller/account"
	"github.com/ProximaB/das/controller/util"
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
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
