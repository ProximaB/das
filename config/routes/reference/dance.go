package reference

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiReferenceDanceEndpoint = "/api/v1.0/reference/dance"

var danceServer = reference.DanceServer{
	database.DanceRepository,
}

var searchDanceController = util.DasController{
	Name:         "SearchDanceController",
	Description:  "Search dances in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceDanceEndpoint,
	Handler:      danceServer.SearchDanceHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createDanceController = util.DasController{
	Name:         "CreateDanceController",
	Description:  "Create a dance in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceDanceEndpoint,
	Handler:      danceServer.CreateDanceHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var deleteDanceController = util.DasController{
	Name:         "DeleteDanceController",
	Description:  "Delete a dance from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceDanceEndpoint,
	Handler:      danceServer.DeleteDanceHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateDanceController = util.DasController{
	Name:         "UpdateDanceController",
	Description:  "Update a dance in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceDanceEndpoint,
	Handler:      danceServer.UpdateDanceHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// DanceControllerGroup is a collection of handler functions for managing dances in DAS
var DanceControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchDanceController,
		createDanceController,
		deleteDanceController,
		updateDanceController,
	},
}
