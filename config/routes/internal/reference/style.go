package reference

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/config/database"
	"github.com/yubing24/das/controller/reference"
	"github.com/yubing24/das/controller/util"
	"net/http"
)

const apiReferenceStyleEndpoint = "/api/reference/style"

var styleServer = reference.StyleServer{
	database.StyleRepository,
}

var searchStyleController = util.DasController{
	Description:  "Search schools in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceStyleEndpoint,
	Handler:      styleServer.SearchStyleHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var createStyleController = util.DasController{
	Description:  "Create a school in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceStyleEndpoint,
	Handler:      styleServer.CreateStyleHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var deleteStyleController = util.DasController{
	Description:  "Delete a school from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceStyleEndpoint,
	Handler:      styleServer.DeleteStyleHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var updateStyleController = util.DasController{
	Description:  "Update a school in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceStyleEndpoint,
	Handler:      styleServer.UpdateStyleHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ADMINISTRATOR},
}

var StyleControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchStyleController,
		createStyleController,
		deleteStyleController,
		updateStyleController,
	},
}
