package reference

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/controller/reference"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiReferenceAgeEndpoint = "/api/v1.0/reference/age"

var ageServer = reference.AgeServer{
	database.AgeRepository,
}

var searchAgeController = util.DasController{
	Name:         "SearchAgeController",
	Description:  "Search schools in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceAgeEndpoint,
	Handler:      ageServer.SearchAgeHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createAgeController = util.DasController{
	Name:         "CreateAgeController",
	Description:  "Create a school in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceAgeEndpoint,
	Handler:      ageServer.CreateAgeHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var deleteAgeController = util.DasController{
	Name:         "DeleteAgeController",
	Description:  "Delete a school from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceAgeEndpoint,
	Handler:      ageServer.DeleteAgeHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateAgeController = util.DasController{
	Name:         "UpdateAgeController",
	Description:  "Update a school in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceAgeEndpoint,
	Handler:      ageServer.UpdateAgeHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// AgeControllerGroup is a collection of handler functions for managing ages in DAS
var AgeControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchAgeController,
		createAgeController,
		deleteAgeController,
		updateAgeController,
	},
}
