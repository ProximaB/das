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
	Description:  "Search unique event federations of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionFederationEndpoint,
	Handler:      publicCompetitionServer.GetUniqueEventFederationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var searchCompetitionUniqueEventDivisionController = util.DasController{
	Name:         "SearchCompetitionUniqueEventDivisionController",
	Description:  "Search unique event divisions of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionDivisionEndpoint,
	Handler:      publicCompetitionServer.GetEventUniqueDivisionsHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
var searchCompetitionUniqueEventAgeController = util.DasController{
	Name:         "SearchCompetitionUniqueEventAgeController",
	Description:  "Search unique event ages of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionAgeEndpoint,
	Handler:      publicCompetitionServer.GetEventUniqueAgesHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
var searchCompetitionUniqueEventProficiencyController = util.DasController{
	Name:         "SearchCompetitionUniqueEventProficiencyController",
	Description:  "Search unique event proficiencies of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionProficiencyEndpoint,
	Handler:      publicCompetitionServer.GetEventUniqueProficienciesHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
var searchCompetitionUniqueEventStyleController = util.DasController{
	Name:         "SearchCompetitionUniqueEventStyleController",
	Description:  "Search unique event styles of a competition",
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
