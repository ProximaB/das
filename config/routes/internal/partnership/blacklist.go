// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package partnership

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/DancesportSoftware/das/controller/partnership/blacklist"
	"github.com/DancesportSoftware/das/controller/util"
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
