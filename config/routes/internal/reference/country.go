package reference

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/reference"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiReferenceCountryEndpoint = "/api/reference/country"

var countryServer = reference.CountryServer{
	database.CountryRepository,
}
var searchCountryController = util.DasController{
	Description:  "Search countries in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceCountryEndpoint,
	Handler:      countryServer.SearchCountryHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var createCountryController = util.DasController{
	Description:  "Create a country in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceCountryEndpoint,
	Handler:      countryServer.CreateCountryHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var deleteCountryController = util.DasController{
	Description:  "Delete a country from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceCountryEndpoint,
	Handler:      countryServer.DeleteCountryHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var updateCountryController = util.DasController{
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
