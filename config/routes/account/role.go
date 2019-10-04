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
const apiAccountRoleApplicationStatus = "/api/account/role/application/status"

var roleProvisionService = businesslogic.NewRoleProvisionService(
	database.AccountRepository,
	database.RoleApplicationRepository,
	database.RoleApplicationStatusRepository,
	database.AccountRoleRepository,
	database.OrganizerProvisionRepository,
	database.OrganizerProvisionHistoryRepository)
var roleApplicationServer = account.NewRoleApplicationServer(middleware.AuthenticationStrategy, *roleProvisionService)

var createRoleApplicationController = util.DasController{
	Name:        "CreateRoleApplicationController",
	Description: "Create a role application in DAS",
	Method:      http.MethodPost,
	Endpoint:    apiAccountRoleCreateApplication,
	Handler:     roleApplicationServer.CreateRoleApplicationHandler,
	AllowedRoles: []int{
		businesslogic.AccountTypeAthlete, // 2018-12-12: all users have athlete role and are granted to apply for other roles
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

var adminSearchRoleApplicationController = util.DasController{
	Name:        "Admin Search Role Application Controller",
	Description: "Search role application without moderation on criteria",
	Method:      http.MethodGet,
	Endpoint:    "/api/v1.0/admin/role/applications",
	Handler:     roleApplicationServer.AdminGetRoleApplicationHandler,
	AllowedRoles: []int{
		businesslogic.AccountTypeAdministrator,
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

var getRoleApplicationStatusController = util.DasController{
	Name:         "GetRoleApplicationStatusController",
	Description:  "Get all possible Role Application Status ",
	Method:       http.MethodGet,
	Endpoint:     apiAccountRoleApplicationStatus,
	Handler:      roleApplicationServer.GetAllApplicationStatus,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

const apiAccountRole = "/api/v1.0/account/role"

var roleServer = account.RoleServer{
	middleware.AuthenticationStrategy,
	database.AccountRepository,
}

var RoleApplicationControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createRoleApplicationController,
		searchRoleApplicationController,
		adminSearchRoleApplicationController,
		provisionRoleApplicationController,
		getRoleApplicationStatusController,
	},
}

var RoleController = util.DasController{
	Name:        "RoleController",
	Description: "Provide the roles of a user",
	Method:      http.MethodGet,
	Endpoint:    apiAccountRole,
	Handler:     roleServer.GetAccountRolesHandler,
	AllowedRoles: []int{
		businesslogic.AccountTypeAthlete,
	},
}
