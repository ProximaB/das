// Dancesport Application System (DAS)
// Copyright (C) 2019 Yubing Hou
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

const apiAdminUserManagementProvision = "/api/v1/admin/user"

var adminUserManagementServer = admin.NewAdminUserManagementServer(middleware.AuthenticationStrategy, database.AccountRepository)

var adminSearchUserController = util.DasController{
	Name:         "AdminSearchUserController",
	Description:  "Search users in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiAdminUserManagementProvision,
	Handler:      adminUserManagementServer.SearchUserHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var AdminManageUserControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		adminSearchUserController,
	},
}
