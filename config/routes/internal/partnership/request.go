// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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
