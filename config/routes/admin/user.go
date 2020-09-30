package admin

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/config/routes/middleware"
	"github.com/ProximaB/das/controller/admin"
	"github.com/ProximaB/das/controller/util"
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
