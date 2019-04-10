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
