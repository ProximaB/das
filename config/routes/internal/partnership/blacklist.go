// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package partnership

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/authentication"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/partnership/blacklist"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiPartnershipRequestBlacklistEndpoint = "/api/partnership/request/blacklist"

var partnershipRequestBlacklistServer = blacklist.PartnershipRequestBlacklistServer{
	authentication.AuthenticationStrategy,
	database.AccountRepository,
	database.PartnershipRequestBlacklistRepository,
}

var searchBlacklistedAccountController = util.DasController{
	Name:         "SearchBlacklistedAccountController",
	Description:  "Search blacklisted account in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipRequestBlacklistEndpoint,
	Handler:      partnershipRequestBlacklistServer.GetBlacklistedAccountHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ATHLETE, businesslogic.AccountTypeAdministrator},
}

var createBlacklistedAccountController = util.DasController{
	Name:         "CreateBlacklistedAccountController",
	Description:  "Create a blacklist report in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiPartnershipRequestBlacklistEndpoint,
	Handler:      partnershipRequestBlacklistServer.CreatePartnershipRequestBlacklistReportHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ATHLETE},
}

// PartnershipRequestBlacklistControllerGroup is a collection of handler functions for managing
// Partnership request blacklist in DAS
var PartnershipRequestBlacklistControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchBlacklistedAccountController,
		createBlacklistedAccountController,
	},
}
