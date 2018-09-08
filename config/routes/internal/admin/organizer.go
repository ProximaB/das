package admin

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/admin"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiAdminManageOrganizerProvision = "/api/v1.0/admin/organizer/provision"

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
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var ManageOrganizerProvisionControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		updateOrganizerProvisionController,
	},
}
