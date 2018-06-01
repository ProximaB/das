package reference

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/reference"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiReferenceStateEndpoint = "/api/reference/state"

var stateServer = reference.StateServer{
	database.StateRepository,
}

var createStateController = util.DasController{
	Description:  "Create a state in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceStateEndpoint,
	Handler:      stateServer.CreateStateHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var searchStateController = util.DasController{
	Description:  "Search states in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceStateEndpoint,
	Handler:      stateServer.SearchStateHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var deleteStateController = util.DasController{
	Description:  "Delete a state in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceStateEndpoint,
	Handler:      stateServer.DeleteStateHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var updateStateController = util.DasController{
	Description:  "Update a state in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceStateEndpoint,
	Handler:      stateServer.UpdateStateHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var StateControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createStateController,
		deleteStateController,
		updateStateController,
		searchStateController,
	},
}
