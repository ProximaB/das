package reference

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/controller/reference"
	"github.com/ProximaB/das/controller/util"
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
