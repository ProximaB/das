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

const apiReferenceStudioEndpoint = "/api/reference/studio"

var studioServer = reference.StudioServer{
	database.StudioRepository,
}

var searchStudioController = util.DasController{
	Name:         "SearchStudioController",
	Description:  "Search dance studios in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiReferenceStudioEndpoint,
	Handler:      studioServer.SearchStudioHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var createStudioController = util.DasController{
	Name:         "CreateStudioController",
	Description:  "Create a dance studio DAS",
	Method:       http.MethodPost,
	Endpoint:     apiReferenceStudioEndpoint,
	Handler:      studioServer.CreateStudioHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator, businesslogic.AccountTypeAthlete},
}

var deleteStudioController = util.DasController{
	Name:         "DeleteStudioController",
	Description:  "Delete a dance studio in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiReferenceStudioEndpoint,
	Handler:      studioServer.DeleteStudioHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

var updateStudioController = util.DasController{
	Name:         "UpdateStudioController",
	Description:  "Update a dance studio in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiReferenceStudioEndpoint,
	Handler:      studioServer.UpdateStudioHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAdministrator},
}

// StudioControllerGroup is a collection of handler functions for managing dance studios in DAS
var StudioControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchStudioController,
		createStudioController,
		deleteStudioController,
		updateStudioController,
	},
}
