package partnership

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/config/routes/middleware"
	"github.com/ProximaB/das/controller/partnership/blacklist"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiPartnershipRequestBlacklistEndpoint = "/api/v1.0/partnership/request/blacklist"

var partnershipRequestBlacklistServer = blacklist.PartnershipRequestBlacklistServer{
	middleware.AuthenticationStrategy,
	database.AccountRepository,
	database.PartnershipRequestBlacklistRepository,
}

var searchBlacklistedAccountController = util.DasController{
	Name:         "SearchBlacklistedAccountController",
	Description:  "Search blacklisted account in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipRequestBlacklistEndpoint,
	Handler:      partnershipRequestBlacklistServer.GetBlacklistedAccountHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete, businesslogic.AccountTypeAdministrator},
}

var createBlacklistedAccountController = util.DasController{
	Name:         "CreateBlacklistedAccountController",
	Description:  "Create a blacklist report in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiPartnershipRequestBlacklistEndpoint,
	Handler:      partnershipRequestBlacklistServer.CreatePartnershipRequestBlacklistReportHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

// PartnershipRequestBlacklistControllerGroup is a collection of handler functions for managing
// Partnership request blacklist in DAS
var PartnershipRequestBlacklistControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchBlacklistedAccountController,
		createBlacklistedAccountController,
	},
}
