// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package admin

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/DancesportSoftware/das/controller/admin"
	"github.com/DancesportSoftware/das/controller/util"
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
