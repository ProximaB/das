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

const apiReferenceCountryEndpoint = "/api/v1.0/reference/country"

var countryServer = reference.CountryServer{
	database.CountryRepository,
}
var searchCountryController = util.DasController{
	Name:         "SearchCountryController",
	Description:  "Search countries in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceCountryEndpoint,
	Handler:      countryServer.SearchCountryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createCountryController = util.DasController{
	Name:         "CreateCountryController",
	Description:  "Create a country in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceCountryEndpoint,
	Handler:      countryServer.CreateCountryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var deleteCountryController = util.DasController{
	Name:         "DeleteCountryController",
	Description:  "Delete a country from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceCountryEndpoint,
	Handler:      countryServer.DeleteCountryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateCountryController = util.DasController{
	Name:         "UpdateCountryController",
	Description:  "Update a country in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceCountryEndpoint,
	Handler:      countryServer.UpdateCountryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// CountryControllerGroup is a collection of handler functions for managing countries in DAS
var CountryControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchCountryController,
		createCountryController,
		deleteCountryController,
		updateCountryController,
	},
}
