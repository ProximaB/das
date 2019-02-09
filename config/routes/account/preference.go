package account

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/DancesportSoftware/das/controller/account"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiUserPreferenceEndpointV1_0 = "/api/v1.0/account/preference"

var preferenceServer = account.UserPreferenceServer{
	middleware.AuthenticationStrategy,
	database.AccountRepository,
	database.UserPreferenceRepository,
}

var searchUserPreferenceHandler = util.DasController{
	Name:        "SearchUserPreferenceHandler",
	Description: "Search user preference in DAS",
	Method:      http.MethodGet,
	Endpoint:    apiUserPreferenceEndpointV1_0,
	Handler:     preferenceServer.GetUserPreferenceHandler,
	AllowedRoles: []int{
		businesslogic.AccountTypeAthlete,
		businesslogic.AccountTypeAdjudicator,
		businesslogic.AccountTypeScrutineer,
		businesslogic.AccountTypeOrganizer,
		businesslogic.AccountTypeDeckCaptain,
		businesslogic.AccountTypeEmcee,
	},
}

var updateUserPreferenceHandler = util.DasController{
	Name:        "UpdateUserPreferenceHandler",
	Description: "Update user preference in DAS",
	Method:      http.MethodPut,
	Endpoint:    apiUserPreferenceEndpointV1_0,
	Handler:     preferenceServer.UpdateUserPreferenceHandler,
	AllowedRoles: []int{
		businesslogic.AccountTypeAthlete,
		businesslogic.AccountTypeAdjudicator,
		businesslogic.AccountTypeScrutineer,
		businesslogic.AccountTypeOrganizer,
		businesslogic.AccountTypeDeckCaptain,
		businesslogic.AccountTypeEmcee,
	},
}

var UserPreferenceControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchUserPreferenceHandler,
		updateUserPreferenceHandler,
	},
}
