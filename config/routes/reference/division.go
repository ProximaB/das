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

const apiReferenceDivisionEndpoint = "/api/v1.0/reference/division"

var divisionServer = reference.DivisionServer{
	database.DivisionRepository,
}

var searchDivisionController = util.DasController{
	Name:         "SearchDivisionController",
	Description:  "Search divisions in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceDivisionEndpoint,
	Handler:      divisionServer.SearchDivisionHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createDivisionController = util.DasController{
	Name:         "CreateDivisionController",
	Description:  "Create a division in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceDivisionEndpoint,
	Handler:      divisionServer.CreateDivisionHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var deleteDivisionController = util.DasController{
	Name:         "DeleteDivisionController",
	Description:  "Delete a division from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceDivisionEndpoint,
	Handler:      divisionServer.DeleteDivisionHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateDivisionController = util.DasController{
	Name:         "UpdateDivisionController",
	Description:  "Update a division in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceDivisionEndpoint,
	Handler:      divisionServer.UpdateDivisionHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// DivisionControllerGroup is a collection of handler functions for managing divisions in DAS
var DivisionControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchDivisionController,
		createDivisionController,
		deleteDivisionController,
		updateDivisionController,
	},
}
