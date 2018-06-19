// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package reference

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiReferenceStyleEndpoint = "/api/reference/style"

var styleServer = reference.StyleServer{
	database.StyleRepository,
}

var searchStyleController = util.DasController{
	Name:         "SearchStyleController",
	Description:  "Search schools in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceStyleEndpoint,
	Handler:      styleServer.SearchStyleHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var createStyleController = util.DasController{
	Name:         "CreateStyleController",
	Description:  "Create a school in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceStyleEndpoint,
	Handler:      styleServer.CreateStyleHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var deleteStyleController = util.DasController{
	Name:         "DeleteStyleController",
	Description:  "Delete a school from DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceStyleEndpoint,
	Handler:      styleServer.DeleteStyleHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateStyleController = util.DasController{
	Name:         "UpdateStyleController",
	Description:  "Update a school in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceStyleEndpoint,
	Handler:      styleServer.UpdateStyleHandler,
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
