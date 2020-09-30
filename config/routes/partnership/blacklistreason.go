package partnership

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/controller/partnership/blacklist"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiPartnershipBlacklistReasonEndpoint = "/api/partnership/blacklist/reason"

var partnershipRequestBlacklistReasonServer = blacklist.PartnershipRequestBlacklistReasonServer{
	database.PartnershipRequestBlacklistReasonRepository,
}

var GetPartnershipBlacklistReasonController = util.DasController{
	Name:         "GetPartnershipBlacklistReasonController",
	Description:  "Get all the partnership blacklist reasons from DAS",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipBlacklistReasonEndpoint,
	Handler:      partnershipRequestBlacklistReasonServer.GetPartnershipBlacklistReasonHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
