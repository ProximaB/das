// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

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
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var searchCompetitionUniqueEventFederationController = util.DasController{
	Name:         "SearchCompetitionUniqueEventFederationController",
	Description:  "Search unique eventdal federations of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionFederationEndpoint,
	Handler:      publicCompetitionServer.GetUniqueEventFederationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var searchCompetitionUniqueEventDivisionController = util.DasController{
	Name:         "SearchCompetitionUniqueEventDivisionController",
	Description:  "Search unique eventdal divisions of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionDivisionEndpoint,
	Handler:      publicCompetitionServer.GetEventUniqueDivisionsHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
var searchCompetitionUniqueEventAgeController = util.DasController{
	Name:         "SearchCompetitionUniqueEventAgeController",
	Description:  "Search unique eventdal ages of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionAgeEndpoint,
	Handler:      publicCompetitionServer.GetEventUniqueAgesHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
var searchCompetitionUniqueEventProficiencyController = util.DasController{
	Name:         "SearchCompetitionUniqueEventProficiencyController",
	Description:  "Search unique eventdal proficiencies of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionProficiencyEndpoint,
	Handler:      publicCompetitionServer.GetEventUniqueProficienciesHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
var searchCompetitionUniqueEventStyleController = util.DasController{
	Name:         "SearchCompetitionUniqueEventStyleController",
	Description:  "Search unique eventdal styles of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionStyleEndpoint,
	Handler:      publicCompetitionServer.GetEventUniqueStylesHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
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
