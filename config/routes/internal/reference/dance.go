package reference

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/reference"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiReferenceDanceEndpoint = "/api/reference/dance"

var danceServer = reference.DanceServer{
	database.DanceRepository,
}

var searchDanceController = util.DasController{
	Description:  "Search dances in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceDanceEndpoint,
	Handler:      danceServer.SearchDanceHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var createDanceController = util.DasController{
	Description:  "Create a dance in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceDanceEndpoint,
	Handler:      danceServer.CreateDanceHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var deleteDanceController = util.DasController{
	Description:  "Delete a dance from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceDanceEndpoint,
	Handler:      danceServer.DeleteDanceHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var updateDanceController = util.DasController{
	Description:  "Update a dance in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceDanceEndpoint,
	Handler:      danceServer.UpdateDanceHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var DanceControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchDanceController,
		createDanceController,
		deleteDanceController,
		updateDanceController,
	},
}
