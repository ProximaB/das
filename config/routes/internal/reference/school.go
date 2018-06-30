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

const apiReferenceSchoolEndpoint = "/api/reference/school"

var schoolServer = reference.SchoolServer{
	database.SchoolRepository,
}

var searchSchoolController = util.DasController{
	Name:         "SearchSchoolController",
	Description:  "Search schools in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceSchoolEndpoint,
	Handler:      schoolServer.SearchSchoolHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createSchoolController = util.DasController{
	Name:         "CreateSchoolController",
	Description:  "Create a school in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceSchoolEndpoint,
	Handler:      schoolServer.CreateSchoolHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator, businesslogic.AccountTypeAthlete},
}

var deleteSchoolController = util.DasController{
	Name:         "DeleteSchoolController",
	Description:  "Delete a school from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceSchoolEndpoint,
	Handler:      schoolServer.DeleteSchoolHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateSchoolController = util.DasController{
	Name:         "UpdateSchoolController",
	Description:  "Update a school in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceSchoolEndpoint,
	Handler:      schoolServer.UpdateSchoolHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// SchoolControllerGroup is a collection of handler functions for managing schools in DAS
var SchoolControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchSchoolController,
		createSchoolController,
		deleteSchoolController,
		updateSchoolController,
	},
}
