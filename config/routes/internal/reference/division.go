package reference

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/reference"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiReferenceDivisionEndpoint = "/api/reference/division"

var divisionServer = reference.DivisionServer{
	database.DivisionRepository,
}

var searchDivisionController = util.DasController{
	Description:  "Search divisions in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceDivisionEndpoint,
	Handler:      divisionServer.SearchDivisionHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var createDivisionController = util.DasController{
	Description:  "Create a division in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceDivisionEndpoint,
	Handler:      divisionServer.CreateDivisionHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var deleteDivisionController = util.DasController{
	Description:  "Delete a division from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceDivisionEndpoint,
	Handler:      divisionServer.DeleteDivisionHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var updateDivisionController = util.DasController{
	Description:  "Update a division in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceDivisionEndpoint,
	Handler:      divisionServer.UpdateDivisionHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var DivisionControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchDivisionController,
		createDivisionController,
		deleteDivisionController,
		updateDivisionController,
	},
}
