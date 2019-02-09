package account

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/DancesportSoftware/das/controller/account"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiAccountRegistrationEndpoint = "/api/v1.0/account/register"
const apiAccountAuthenticationEndpoint = "/api/v1.0/account/authenticate"

var accountServer = account.AccountServer{
	middleware.AuthenticationStrategy,
	database.AccountRepository,
	database.AccountRoleRepository,
	database.OrganizerProvisionRepository,
	database.OrganizerProvisionHistoryRepository,
	database.UserPreferenceRepository,
}

var accountRegistrationController = util.DasController{
	Name:         "AccountRegistrationController",
	Description:  "Create an account in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiAccountRegistrationEndpoint,
	Handler:      accountServer.RegisterAccountHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var accountAuthenticationController = util.DasController{
	Name:         "AccountAuthenticationController",
	Description:  "Authenticate user account",
	Method:       http.MethodPost,
	Endpoint:     apiAccountAuthenticationEndpoint,
	Handler:      accountServer.AccountAuthenticationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var AccountControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		accountRegistrationController,
		accountAuthenticationController,
	},
}
