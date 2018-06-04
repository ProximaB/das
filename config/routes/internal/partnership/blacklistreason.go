package partnership

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/partnership/blacklist"
	"github.com/DancesportSoftware/das/controller/util"
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
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}
