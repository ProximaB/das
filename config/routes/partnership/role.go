package partnership

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/controller/partnership"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiPartnershipRoleEndpoint = "/api/v1.0/partnership/role"

var partnershipRoleServer = partnership.PartnershipRoleServer{
	database.PartnershipRoleRepository,
}

var GetPartnershipRoleController = util.DasController{
	Name:         "GetPartnershipRoleController",
	Description:  "Get all roles of partnership",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipRoleEndpoint,
	Handler:      partnershipRoleServer.GetPartnershipRolesHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
