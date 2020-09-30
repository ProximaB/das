package partnership

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/controller/partnership/request"
	"github.com/ProximaB/das/controller/util"
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
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
