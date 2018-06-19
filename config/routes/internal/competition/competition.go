// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package competition

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/competition"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const (
	apiCompetitionEndpoint            = "/api/competitions"
	apiCompetitionFederationEndpoint  = "/api/competition/federation"
	apiCompetitionDivisionEndpoint    = "/api/competition/division"
	apiCompetitionAgeEndpoint         = "/api/competition/age"
	apiCompetitionProficiencyEndpoint = "/api/competition/proficiency"
	apiCompetitionStyleEndpoint       = "/api/competition/style"
)

var publicCompetitionServer = competition.PublicCompetitionServer{
	database.CompetitionRepository,
	database.EventRepository,
	database.EventMetaRepository,
}

var searchCompetitionController = util.DasController{
	Name:         "SearchOpenCompetitionController",
	Description:  "Search competitions that are open",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionEndpoint,
	Handler:      publicCompetitionServer.SearchCompetitionHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var searchCompetitionUniqueEventFederationController = util.DasController{
	Name:         "SearchCompetitionUniqueEventFederationController",
	Description:  "Search unique event federations of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionFederationEndpoint,
	Handler:      publicCompetitionServer.GetUniqueEventFederationHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var searchCompetitionUniqueEventDivisionController = util.DasController{
	Name:         "SearchCompetitionUniqueEventDivisionController",
	Description:  "Search unique event divisions of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionDivisionEndpoint,
	Handler:      publicCompetitionServer.GetEventUniqueDivisionsHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}
var searchCompetitionUniqueEventAgeController = util.DasController{
	Name:         "SearchCompetitionUniqueEventAgeController",
	Description:  "Search unique event ages of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionAgeEndpoint,
	Handler:      publicCompetitionServer.GetEventUniqueAgesHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}
var searchCompetitionUniqueEventProficiencyController = util.DasController{
	Name:         "SearchCompetitionUniqueEventProficiencyController",
	Description:  "Search unique event proficiencies of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionProficiencyEndpoint,
	Handler:      publicCompetitionServer.GetEventUniqueProficienciesHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}
var searchCompetitionUniqueEventStyleController = util.DasController{
	Name:         "SearchCompetitionUniqueEventStyleController",
	Description:  "Search unique event styles of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionStyleEndpoint,
	Handler:      publicCompetitionServer.GetEventUniqueStylesHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_NOAUTH},
}

var PublicCompetitionViewControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchCompetitionController,
		searchCompetitionUniqueEventFederationController,
		searchCompetitionUniqueEventDivisionController,
		searchCompetitionUniqueEventAgeController,
		searchCompetitionUniqueEventProficiencyController,
		searchCompetitionUniqueEventStyleController,
	},
}
