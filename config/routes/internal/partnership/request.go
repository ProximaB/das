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
	"github.com/DancesportSoftware/das/config/authentication"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/partnership/request"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiPartnershipRequestEndpoint = "/api/partnership/request"

var partnershipRequestServer = request.PartnershipRequestServer{
	authentication.AuthenticationStrategy,
	database.AccountRepository,
	database.PartnershipRepository,
	database.PartnershipRequestRepository,
	database.PartnershipRequestBlacklistRepository,
}

var createPartnershipRequestController = util.DasController{
	Name:         "CreatePartnershipRequestController",
	Description:  "Create a new partnership request in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiPartnershipRequestEndpoint,
	Handler:      partnershipRequestServer.CreatePartnershipRequestHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

var searchPartnershipRequestController = util.DasController{
	Name:         "SearchPartnershipRequestController",
	Description:  "Search a new partnership request in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipRequestEndpoint,
	Handler:      partnershipRequestServer.SearchPartnershipRequestHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

var updatePartnershipRequestController = util.DasController{
	Name:         "UpdatePartnershipRequestController",
	Description:  "Update a new partnership request in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiPartnershipRequestEndpoint,
	Handler:      partnershipRequestServer.UpdatePartnershipRequestHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

var deletePartnershipRequestController = util.DasController{
	Name:         "DeletePartnershipRequestController",
	Description:  "delete a new partnership request in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiPartnershipRequestEndpoint,
	Handler:      partnershipRequestServer.DeletePartnershipRequestHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

var PartnershipRequestControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createPartnershipRequestController,
		searchPartnershipRequestController,
		updatePartnershipRequestController,
		deletePartnershipRequestController,
	},
}
