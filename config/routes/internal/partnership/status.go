package partnership

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/partnership/request"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

var partnershipRequestStatusServer = request.PartnershipRequestStatusServer{
	database.PartnershipRequestStatusRepository,
}

var PartnershipRequestStatusController = util.DasController{
	Name:         "PartnershipRequestStatusController",
	Description:  "Search partnership request status in DAS",
	Method:       http.MethodGet,
	Endpoint:     "/api/partnership/request/status",
	Handler:      partnershipRequestStatusServer.GetPartnershipRequestStatusHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}
