package reference

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiReferenceStyleEndpointV1_0 = "/api/v1.0/reference/style"

var styleServerV1_0 = reference.StyleServer{
	database.StyleRepository,
}

var searchStyleController = util.DasController{
	Name:         "SearchStyleController",
	Description:  "Search schools in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceStyleEndpointV1_0,
	Handler:      styleServerV1_0.SearchStyleHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createStyleController = util.DasController{
	Name:         "CreateStyleController",
	Description:  "Create a school in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceStyleEndpointV1_0,
	Handler:      styleServerV1_0.CreateStyleHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var deleteStyleController = util.DasController{
	Name:         "DeleteStyleController",
	Description:  "Delete a school from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceStyleEndpointV1_0,
	Handler:      styleServerV1_0.DeleteStyleHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateStyleController = util.DasController{
	Name:         "UpdateStyleController",
	Description:  "Update a school in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceStyleEndpointV1_0,
	Handler:      styleServerV1_0.UpdateStyleHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// StyleControllerGroup is a collection of handler functions for managing dance styles in DAS
var StyleControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchStyleController,
		createStyleController,
		deleteStyleController,
		updateStyleController,
	},
}
