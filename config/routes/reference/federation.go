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

package reference

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiReferenceFederationEndpoint = "/api/v1.0/reference/federation"

var federationServer = reference.FederationServer{
	database.FederationRepository,
}

var searchFederationController = util.DasController{
	Name:         "SearchFederationController",
	Description:  "Search federations in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceFederationEndpoint,
	Handler:      federationServer.SearchFederationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createFederationController = util.DasController{
	Name:         "CreateFederationController",
	Description:  "Create a federation in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceFederationEndpoint,
	Handler:      federationServer.CreateFederationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var deleteFederationController = util.DasController{
	Name:         "DeleteFederationController",
	Description:  "Delete a federation from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceFederationEndpoint,
	Handler:      federationServer.DeleteFederationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateFederationController = util.DasController{
	Name:         "UpdateFederationController",
	Description:  "Update a federation in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceFederationEndpoint,
	Handler:      federationServer.UpdateFederationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// FederationControllerGroup is a collection of handler functions for managing federations in DAS
var FederationControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchFederationController,
		createFederationController,
		deleteFederationController,
		updateFederationController,
	},
}
