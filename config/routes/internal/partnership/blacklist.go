package partnership

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/partnership/blacklist"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiPartnershipRequestBlacklistEndpoint = "/api/partnership/request/blacklist"

var partnershipRequestBlacklistServer = blacklist.PartnershipRequestBlacklistServer{
	database.AccountRepository,
	database.PartnershipRequestBlacklistRepository,
}

var searchBlacklistedAccountController = util.DasController{
	Description:  "Search blacklisted account in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipRequestBlacklistEndpoint,
	Handler:      partnershipRequestBlacklistServer.GetBlacklistedAccountHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ATHLETE, businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var createBlacklistedAccountController = util.DasController{
	Description:  "Create a blacklist report in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiPartnershipRequestBlacklistEndpoint,
	Handler:      partnershipRequestBlacklistServer.CreatePartnershipRequestBlacklistReportHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ATHLETE},
}

var PartnershipRequestBlacklistControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchBlacklistedAccountController,
		createBlacklistedAccountController,
	},
}
