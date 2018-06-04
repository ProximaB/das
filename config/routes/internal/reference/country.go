package reference

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiReferenceCountryEndpoint = "/api/reference/country"

var countryServer = reference.CountryServer{
	database.CountryRepository,
}
var searchCountryController = util.DasController{
	Name:         "SearchCountryController",
	Description:  "Search countries in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceCountryEndpoint,
	Handler:      countryServer.SearchCountryHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var createCountryController = util.DasController{
	Name:         "CreateCountryController",
	Description:  "Create a country in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceCountryEndpoint,
	Handler:      countryServer.CreateCountryHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var deleteCountryController = util.DasController{
	Name:         "DeleteCountryController",
	Description:  "Delete a country from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceCountryEndpoint,
	Handler:      countryServer.DeleteCountryHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var updateCountryController = util.DasController{
	Name:         "UpdateCountryController",
	Description:  "Update a country in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceCountryEndpoint,
	Handler:      countryServer.UpdateCountryHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var CountryControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchCountryController,
		createCountryController,
		deleteCountryController,
		updateCountryController,
	},
}
