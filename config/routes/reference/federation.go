package reference

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/controller/reference"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiReferenceFederationEndpoint = "/api/v1.0/reference/federation"

var federationServer = reference.FederationServer{
	database.FederationRepository,
}

var searchFederationController = util.DasController{
	Name:         "SearchFederationController",
	Description:  "Search federations in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceFederationEndpoint,
	Handler:      federationServer.SearchFederationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createFederationController = util.DasController{
	Name:         "CreateFederationController",
	Description:  "Create a federation in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceFederationEndpoint,
	Handler:      federationServer.CreateFederationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var deleteFederationController = util.DasController{
	Name:         "DeleteFederationController",
	Description:  "Delete a federation from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceFederationEndpoint,
	Handler:      federationServer.DeleteFederationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateFederationController = util.DasController{
	Name:         "UpdateFederationController",
	Description:  "Update a federation in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceFederationEndpoint,
	Handler:      federationServer.UpdateFederationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// FederationControllerGroup is a collection of handler functions for managing federations in DAS
var FederationControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchFederationController,
		createFederationController,
		deleteFederationController,
		updateFederationController,
	},
}
