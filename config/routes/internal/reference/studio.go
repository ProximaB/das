package reference

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/reference"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiReferenceStudioEndpoint = "/api/reference/studio"

var studioServer = reference.StudioServer{
	database.StudioRepository,
}

var searchStudioController = util.DasController{
	Description:  "Search dance studios in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceStudioEndpoint,
	Handler:      studioServer.SearchStudioHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var createStudioController = util.DasController{
	Description:  "Create a dance studio DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceStudioEndpoint,
	Handler:      studioServer.CreateStudioHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR, businesslogic.ACCOUNT_TYPE_ATHLETE},
}

var deleteStudioController = util.DasController{
	Description:  "Delete a dance studio in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceStudioEndpoint,
	Handler:      studioServer.DeleteStudioHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var updateStudioController = util.DasController{
	Description:  "Update a dance studio in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceStudioEndpoint,
	Handler:      studioServer.UpdateStudioHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var StudioControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchStudioController,
		createStudioController,
		deleteStudioController,
		updateStudioController,
	},
}
