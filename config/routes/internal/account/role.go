package account

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/DancesportSoftware/das/controller/account"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiAccountRoleCreateApplication = "/api/v1.0/account/role/application"
const apiAccountRoleRespondApplication = "/api/v1.0/account/role/provision" // Admin use only

var roleProvisionService = businesslogic.NewRoleProvisionService(database.AccountRepository, database.RoleApplicationRepository, database.AccountRoleRepository)
var roleApplicationServer = account.NewRoleApplicationServer(middleware.AuthenticationStrategy, *roleProvisionService)

var createRoleApplicationController = util.DasController{
	Name:        "CreateRoleApplicationController",
	Description: "Create a role application in DAS",
	Method:      http.MethodPost,
	Endpoint:    apiAccountRoleCreateApplication,
	Handler:     roleApplicationServer.CreateRoleApplicationHandler,
	AllowedRoles: []int{
		businesslogic.AccountTypeAthlete,
		businesslogic.AccountTypeAdjudicator,
		businesslogic.AccountTypeScrutineer,
		businesslogic.AccountTypeOrganizer,
		businesslogic.AccountTypeDeckCaptain,
		businesslogic.AccountTypeEmcee,
	},
}

var searchRoleApplicationController = util.DasController{
	Name:        "SearchRoleApplicationController",
	Description: "Search role applications in DAS",
	Method:      http.MethodGet,
	Endpoint:    apiAccountRoleCreateApplication,
	Handler:     roleApplicationServer.SearchRoleApplicationHandler,
	AllowedRoles: []int{
		businesslogic.AccountTypeAthlete,
		businesslogic.AccountTypeAdjudicator,
		businesslogic.AccountTypeScrutineer,
		businesslogic.AccountTypeOrganizer,
		businesslogic.AccountTypeDeckCaptain,
		businesslogic.AccountTypeEmcee,
	},
}

var provisionRoleApplicationController = util.DasController{
	Name:        "ProvisionRoleApplicationController",
	Description: "Admin provision applications to restricted roles",
	Method:      http.MethodPut,
	Endpoint:    apiAccountRoleRespondApplication,
	Handler:     roleApplicationServer.ProvisionRoleApplicationHandler,
	AllowedRoles: []int{
		businesslogic.AccountTypeOrganizer,
		businesslogic.AccountTypeAdministrator,
	},
}

var RoleApplicationControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createRoleApplicationController,
		searchRoleApplicationController,
		provisionRoleApplicationController,
	},
}
