// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package partnership

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/authentication"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/partnership"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiPartnershipEndpoint = "/api/partnership"

var partnershipServer = partnership.PartnershipServer{
	authentication.AuthenticationStrategy,
	database.AccountRepository,
	database.PartnershipRepository,
}

var searchPartnershipController = util.DasController{
	Name:         "SearchPartnershipController",
	Description:  "Search partnerships in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipEndpoint,
	Handler:      partnershipServer.SearchPartnershipHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ATHLETE},
}

var updatePartnershipController = util.DasController{
	Name:         "UpdatePartnershipController",
	Description:  "Update a partnership in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiPartnershipEndpoint,
	Handler:      partnershipServer.UpdatePartnershipHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ATHLETE},
}

var PartnershipControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchPartnershipController,
		updatePartnershipController,
	},
}
