package organizer

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/authentication"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/admin"
	"github.com/DancesportSoftware/das/controller/organizer"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiAdminManageOrganizerProvision = "/api/admin/organizer/provision"

var manageOrganizerProvisionServer = admin.OrganizerProvisionServer{
	database.AccountRepository,
	database.OrganizerProvisionRepository,
}

var updateOrganizerProvisionController = util.DasController{
	Name:         "UpdateOrganizerProvisionController",
	Description:  "Update an organizer's provision in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiAdminManageOrganizerProvision,
	Handler:      manageOrganizerProvisionServer.UpdateOrganizerProvisionHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var ManageOrganizerProvisionControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		updateOrganizerProvisionController,
	},
}
var ProvisionControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{},
}

const apiOrganizerProvisionSummaryEndpoint = "/api/organizer/provision/summary"
const apiOrganizerProvisionHistoryEndpoint = "/api/organizer/provision/history"

var organizerProvisionServer = organizer.OrganizerProvisionServer{
	authentication.AuthenticationStrategy,
	database.AccountRepository,
	database.OrganizerProvisionRepository,
}

var getOrganizerProvisionSummaryController = util.DasController{
	Name:         "GetOrganizerProvisionSummaryController",
	Description:  "Retrieve organizer provision information for organizer",
	Method:       http.MethodGet,
	Endpoint:     apiOrganizerProvisionSummaryEndpoint,
	Handler:      organizerProvisionServer.GetOrganizerProvisionSummaryHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ORGANIZER},
}

var organizerProvisionHistoryServer = organizer.OrganizerProvisionHistoryServer{
	authentication.AuthenticationStrategy,
	database.AccountRepository,
	database.OrganizerProvisionHistoryRepository,
}
var getOrganizerProvisionHistoryController = util.DasController{
	Name:         "GetOrganizerProvisionHistoryController",
	Description:  "Retrieve organizer provision history for organizer",
	Method:       http.MethodGet,
	Endpoint:     apiOrganizerProvisionHistoryEndpoint,
	Handler:      organizerProvisionHistoryServer.GetOrganizerProvisionHistoryHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ORGANIZER},
}
var OrganizerProvisionControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		getOrganizerProvisionHistoryController,
		getOrganizerProvisionSummaryController,
	},
}
