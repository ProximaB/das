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
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
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
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
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
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}
var OrganizerProvisionControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		getOrganizerProvisionHistoryController,
		getOrganizerProvisionSummaryController,
	},
}
