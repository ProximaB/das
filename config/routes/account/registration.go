package account

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/config/routes/middleware"
	"github.com/ProximaB/das/controller/account"
	"github.com/ProximaB/das/controller/util"
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
