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

const apiReferenceStateEndpoint = "/api/v1.0/reference/state"

var stateServer = reference.StateServer{
	database.StateRepository,
}

var createStateController = util.DasController{
	Name:         "CreateStateController",
	Description:  "Create a state in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceStateEndpoint,
	Handler:      stateServer.CreateStateHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var searchStateController = util.DasController{
	Name:         "SearchStateController",
	Description:  "Search states in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceStateEndpoint,
	Handler:      stateServer.SearchStateHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var deleteStateController = util.DasController{
	Name:         "DeleteStateController",
	Description:  "Delete a state in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceStateEndpoint,
	Handler:      stateServer.DeleteStateHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateStateController = util.DasController{
	Name:         "UpdateStateController",
	Description:  "Update a state in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceStateEndpoint,
	Handler:      stateServer.UpdateStateHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// StateControllerGroup is a collection of handler functions for managing states in DAS
var StateControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createStateController,
		deleteStateController,
		updateStateController,
		searchStateController,
	},
}
