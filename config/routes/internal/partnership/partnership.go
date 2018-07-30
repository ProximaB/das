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
	"github.com/DancesportSoftware/das/controller/partnership"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiPartnershipEndpoint = "/api/v1.0/partnership"

var partnershipServer = partnership.PartnershipServer{
	middleware.AuthenticationStrategy,
	database.AccountRepository,
	database.PartnershipRepository,
}

var searchPartnershipController = util.DasController{
	Name:         "SearchPartnershipController",
	Description:  "Search partnerships in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipEndpoint,
	Handler:      partnershipServer.SearchPartnershipHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

var updatePartnershipController = util.DasController{
	Name:         "UpdatePartnershipController",
	Description:  "Update a partnership in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiPartnershipEndpoint,
	Handler:      partnershipServer.UpdatePartnershipHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

// PartnershipControllerGroup contains a collection of HTTP request handler functions for
// Partnership related request
var PartnershipControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchPartnershipController,
		updatePartnershipController,
	},
}
