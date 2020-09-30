package reference

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/controller/reference"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiReferenceDivisionEndpoint = "/api/v1.0/reference/division"

var divisionServer = reference.DivisionServer{
	database.DivisionRepository,
}

var searchDivisionController = util.DasController{
	Name:         "SearchDivisionController",
	Description:  "Search divisions in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceDivisionEndpoint,
	Handler:      divisionServer.SearchDivisionHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createDivisionController = util.DasController{
	Name:         "CreateDivisionController",
	Description:  "Create a division in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceDivisionEndpoint,
	Handler:      divisionServer.CreateDivisionHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var deleteDivisionController = util.DasController{
	Name:         "DeleteDivisionController",
	Description:  "Delete a division from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceDivisionEndpoint,
	Handler:      divisionServer.DeleteDivisionHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateDivisionController = util.DasController{
	Name:         "UpdateDivisionController",
	Description:  "Update a division in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceDivisionEndpoint,
	Handler:      divisionServer.UpdateDivisionHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// DivisionControllerGroup is a collection of handler functions for managing divisions in DAS
var DivisionControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchDivisionController,
		createDivisionController,
		deleteDivisionController,
		updateDivisionController,
	},
}
