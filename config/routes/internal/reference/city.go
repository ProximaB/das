package reference

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiReferenceCityEndpoint = "/api/reference/city"

var cityServer = reference.CityServer{
	ICityRepository: database.CityRepository,
}

var createCityController = util.DasController{
	Name:         "CreateCityController",
	Description:  "Create a city in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceCityEndpoint,
	Handler:      cityServer.CreateCityHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var searchCityController = util.DasController{
	Name:         "SearchCityController",
	Description:  "Search cities in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceCityEndpoint,
	Handler:      cityServer.SearchCityHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var deleteCityController = util.DasController{
	Name:         "DeleteCityController",
	Description:  "Delete a city in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceCityEndpoint,
	Handler:      cityServer.DeleteCityHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var updateCityController = util.DasController{
	Name:         "UpdateCityController",
	Description:  "Update a city in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceCityEndpoint,
	Handler:      cityServer.UpdateCityHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var CityControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createCityController,
		deleteCityController,
		updateCityController,
		searchCityController,
	},
}
