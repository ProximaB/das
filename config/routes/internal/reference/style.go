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

const apiReferenceStyleEndpoint = "/api/reference/style"

var styleServer = reference.StyleServer{
	database.StyleRepository,
}

var searchStyleController = util.DasController{
	Name:         "SearchStyleController",
	Description:  "Search schools in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceStyleEndpoint,
	Handler:      styleServer.SearchStyleHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createStyleController = util.DasController{
	Name:         "CreateStyleController",
	Description:  "Create a school in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceStyleEndpoint,
	Handler:      styleServer.CreateStyleHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var deleteStyleController = util.DasController{
	Name:         "DeleteStyleController",
	Description:  "Delete a school from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceStyleEndpoint,
	Handler:      styleServer.DeleteStyleHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateStyleController = util.DasController{
	Name:         "UpdateStyleController",
	Description:  "Update a school in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceStyleEndpoint,
	Handler:      styleServer.UpdateStyleHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// StyleControllerGroup is a collection of handler functions for managing dance styles in DAS
var StyleControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchStyleController,
		createStyleController,
		deleteStyleController,
		updateStyleController,
	},
}
