package reference

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiReferenceStateEndpoint = "/api/v1.0/reference/state"

var stateServer = reference.StateServer{
	database.StateRepository,
}

var createStateController = util.DasController{
	Name:         "CreateStateController",
	Description:  "Create a state in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceStateEndpoint,
	Handler:      stateServer.CreateStateHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var searchStateController = util.DasController{
	Name:         "SearchStateController",
	Description:  "Search states in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceStateEndpoint,
	Handler:      stateServer.SearchStateHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var deleteStateController = util.DasController{
	Name:         "DeleteStateController",
	Description:  "Delete a state in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceStateEndpoint,
	Handler:      stateServer.DeleteStateHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateStateController = util.DasController{
	Name:         "UpdateStateController",
	Description:  "Update a state in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceStateEndpoint,
	Handler:      stateServer.UpdateStateHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// StateControllerGroup is a collection of handler functions for managing states in DAS
var StateControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createStateController,
		deleteStateController,
		updateStateController,
		searchStateController,
	},
}
