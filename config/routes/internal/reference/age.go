package reference

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/reference"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiReferenceAgeEndpoint = "/api/reference/age"

var ageServer = reference.AgeServer{
	database.AgeRepository,
}

var searchAgeController = util.DasController{
	Description:  "Search schools in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceAgeEndpoint,
	Handler:      ageServer.SearchAgeHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var createAgeController = util.DasController{
	Description:  "Create a school in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceAgeEndpoint,
	Handler:      ageServer.CreateAgeHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var deleteAgeController = util.DasController{
	Description:  "Delete a school from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceAgeEndpoint,
	Handler:      ageServer.DeleteAgeHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var updateAgeController = util.DasController{
	Description:  "Update a school in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceAgeEndpoint,
	Handler:      ageServer.UpdateAgeHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var AgeControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchAgeController,
		createAgeController,
		deleteAgeController,
		updateAgeController,
	},
}
