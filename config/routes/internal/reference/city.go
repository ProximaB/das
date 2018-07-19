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

const apiReferenceCityEndpoint = "/api/v1.0/reference/city"

var cityServer = reference.CityServer{
	ICityRepository: database.CityRepository,
}

var createCityController = util.DasController{
	Name:         "CreateCityController",
	Description:  "Create a city in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceCityEndpoint,
	Handler:      cityServer.CreateCityHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var searchCityController = util.DasController{
	Name:         "SearchCityController",
	Description:  "Search cities in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceCityEndpoint,
	Handler:      cityServer.SearchCityHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var deleteCityController = util.DasController{
	Name:         "DeleteCityController",
	Description:  "Delete a city in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceCityEndpoint,
	Handler:      cityServer.DeleteCityHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateCityController = util.DasController{
	Name:         "UpdateCityController",
	Description:  "Update a city in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceCityEndpoint,
	Handler:      cityServer.UpdateCityHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// CityControllerGroup is a collection of handler functions for managing cities in DAS
var CityControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createCityController,
		deleteCityController,
		updateCityController,
		searchCityController,
	},
}
