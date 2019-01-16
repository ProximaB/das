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

const apiReferenceAgeEndpoint = "/api/v1.0/reference/age"

var ageServer = reference.AgeServer{
	database.AgeRepository,
}

var searchAgeController = util.DasController{
	Name:         "SearchAgeController",
	Description:  "Search schools in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceAgeEndpoint,
	Handler:      ageServer.SearchAgeHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createAgeController = util.DasController{
	Name:         "CreateAgeController",
	Description:  "Create a school in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceAgeEndpoint,
	Handler:      ageServer.CreateAgeHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var deleteAgeController = util.DasController{
	Name:         "DeleteAgeController",
	Description:  "Delete a school from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceAgeEndpoint,
	Handler:      ageServer.DeleteAgeHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateAgeController = util.DasController{
	Name:         "UpdateAgeController",
	Description:  "Update a school in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceAgeEndpoint,
	Handler:      ageServer.UpdateAgeHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// AgeControllerGroup is a collection of handler functions for managing ages in DAS
var AgeControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchAgeController,
		createAgeController,
		deleteAgeController,
		updateAgeController,
	},
}
