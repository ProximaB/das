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

const apiReferenceStudioEndpoint = "/api/reference/studio"

var studioServer = reference.StudioServer{
	database.StudioRepository,
}

var searchStudioController = util.DasController{
	Name:         "SearchStudioController",
	Description:  "Search dance studios in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceStudioEndpoint,
	Handler:      studioServer.SearchStudioHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createStudioController = util.DasController{
	Name:         "CreateStudioController",
	Description:  "Create a dance studio DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceStudioEndpoint,
	Handler:      studioServer.CreateStudioHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator, businesslogic.AccountTypeAthlete},
}

var deleteStudioController = util.DasController{
	Name:         "DeleteStudioController",
	Description:  "Delete a dance studio in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceStudioEndpoint,
	Handler:      studioServer.DeleteStudioHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateStudioController = util.DasController{
	Name:         "UpdateStudioController",
	Description:  "Update a dance studio in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceStudioEndpoint,
	Handler:      studioServer.UpdateStudioHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// StudioControllerGroup is a collection of handler functions for managing dance studios in DAS
var StudioControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchStudioController,
		createStudioController,
		deleteStudioController,
		updateStudioController,
	},
}
