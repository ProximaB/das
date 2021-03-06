package admin

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/config/routes/middleware"
	"github.com/ProximaB/das/controller/admin"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiAdminManageOrganizerProvision = "/api/v1.0/admin/organizer/provision"

var organizerProvisionService = businesslogic.NewOrganizerProvisionService(
	database.AccountRepository,
	database.AccountRoleRepository,
	database.OrganizerProvisionRepository,
	database.OrganizerProvisionHistoryRepository)
var manageOrganizerProvisionServer = admin.NewOrganizerProvisionServer(
	middleware.AuthenticationStrategy,
	database.AccountRepository,
	organizerProvisionService)

var updateOrganizerProvisionController = util.DasController{
	Name:         "UpdateOrganizerProvisionController",
	Description:  "Update an organizer's provision in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiAdminManageOrganizerProvision,
	Handler:      manageOrganizerProvisionServer.UpdateOrganizerProvisionHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var getOrganizerProvisionSummaryController = util.DasController{
	Name:         "GetOrganizerProvisionSummaryController",
	Description:  "Admin gets the summarized provision information of organizer",
	Method:       http.MethodGet,
	Endpoint:     apiAdminManageOrganizerProvision,
	Handler:      manageOrganizerProvisionServer.GetOrganizerProvisionSummaryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var ManageOrganizerProvisionControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		updateOrganizerProvisionController,
		getOrganizerProvisionSummaryController,
	},
}
