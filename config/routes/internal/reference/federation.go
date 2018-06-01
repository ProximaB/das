package reference

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/reference"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiReferenceFederationEndpoint = "/api/reference/federation"

var federationServer = reference.FederationServer{
	database.FederationRepository,
}

var searchFederationController = util.DasController{
	Description:  "Search federations in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceFederationEndpoint,
	Handler:      federationServer.SearchFederationHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var createFederationController = util.DasController{
	Description:  "Create a federation in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceFederationEndpoint,
	Handler:      federationServer.CreateFederationHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var deleteFederationController = util.DasController{
	Description:  "Delete a federation from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceFederationEndpoint,
	Handler:      federationServer.DeleteFederationHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var updateFederationController = util.DasController{
	Description:  "Update a federation in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceFederationEndpoint,
	Handler:      federationServer.UpdateFederationHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var FederationControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchFederationController,
		createFederationController,
		deleteFederationController,
		updateFederationController,
	},
}
