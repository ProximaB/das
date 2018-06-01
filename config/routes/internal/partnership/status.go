package partnership

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/partnership/request"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

var partnershipRequestStatusServer = request.PartnershipRequestStatusServer{
	database.PartnershipRequestStatusRepository,
}

var PartnershipRequestStatusController = util.DasController{
	Description:  "Search partnership request status in DAS",
	Method:       http.MethodGet,
	Endpoint:     "/api/partnership/request/status",
	Handler:      partnershipRequestStatusServer.GetPartnershipRequestStatusHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}
