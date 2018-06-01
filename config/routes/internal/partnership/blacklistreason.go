package partnership

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/partnership/blacklist"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiPartnershipBlacklistReasonEndpoint = "/api/partnership/blacklist/reason"

var partnershipRequestBlacklistReasonServer = blacklist.PartnershipRequestBlacklistReasonServer{
	database.PartnershipRequestBlacklistReasonRepository,
}

var GetPartnershipBlacklistReasonController = util.DasController{
	Description:  "Get all the partnership blacklist reasons from DAS",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipBlacklistReasonEndpoint,
	Handler:      partnershipRequestBlacklistReasonServer.GetPartnershipBlacklistReasonHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}
