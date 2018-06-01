package organizer

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/admin"
	"github.com/yubing24/das/controller/organizer"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiAdminManageOrganizerProvision = "/api/admin/organizer/provision"

var manageOrganizerProvisionServer = admin.OrganizerProvisionServer{
	database.AccountRepository,
	database.OrganizerProvisionRepository,
}

var updateOrganizerProvisionController = util.DasController{
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

const apiOrganizerProvisionSummaryEndpoint = "/api/organizer/organizer/summary"
const apiOrganizerProvisionHistoryEndpoint = "/api/organizer/organizer/history"

var ProvisionControllerGroup = util.DasControllerGroup{}

var organizerProvisionServer = organizer.OrganizerProvisionServer{
	database.AccountRepository,
	database.OrganizerProvisionRepository,
}

var getOrganizerProvisionSummaryController = util.DasController{
	Endpoint: apiOrganizerProvisionSummaryEndpoint,
	Handler:  organizerProvisionServer.GetOrganizerProvisionSummaryHandler,
}

var organizerProvisionHistoryServer = organizer.OrganizerProvisionHistoryServer{
	database.AccountRepository,
	database.OrganizerProvisionHistoryRepository,
}
var getOrganizerProvisionHistoryController = util.DasController{
	Endpoint: apiOrganizerProvisionHistoryEndpoint,
	Handler:  organizerProvisionHistoryServer.GetOrganizerProvisionHistoryHandler,
}
var OrganizerProvisionControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		getOrganizerProvisionHistoryController,
		getOrganizerProvisionSummaryController,
	},
}
