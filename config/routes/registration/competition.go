package registration

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/config/routes/middleware"
	"github.com/ProximaB/das/controller"
	"github.com/ProximaB/das/controller/athlete"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiAthleteCompetitionRegistrationEndpoint = "/api/v1.0/athlete/competition/registration"

var athleteCompetitionRegistrationServer = athlete.CompetitionRegistrationServer{
	IAccountRepository:                     database.AccountRepository,
	ICompetitionRepository:                 database.CompetitionRepository,
	IAthleteCompetitionEntryRepository:     database.AthleteCompetitionEntryRepository,
	IPartnershipCompetitionEntryRepository: database.PartnershipCompetitionEntryRepository,
	IPartnershipRepository:                 database.PartnershipRepository,
	IPartnershipEventEntryRepository:       database.PartnershipEventEntryRepository,
	IEventRepository:                       database.EventRepository,
	IAuthenticationStrategy:                middleware.AuthenticationStrategy,
	Service: businesslogic.NewCompetitionRegistrationService(
		database.AccountRepository,
		database.PartnershipRepository,
		database.CompetitionRepository,
		database.EventRepository,
		database.AthleteCompetitionEntryRepository,
		database.AthleteEventEntryRepository,
		database.PartnershipCompetitionEntryRepository,
		database.PartnershipEventEntryRepository,
	),
}

var createCompetitionRegistrationController = util.DasController{
	Name:         "CreateCompetitionRegistrationController",
	Description:  "Athlete creates competition and event registration",
	Method:       http.MethodPost,
	Endpoint:     apiAthleteCompetitionRegistrationEndpoint,
	Handler:      athleteCompetitionRegistrationServer.CreateAthleteRegistrationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

var getPartnershipRegistrationController = util.DasController{
	Name:         "GetPartnershipRegistrationController",
	Description:  "Athlete get the registration for the partnership and competition selected",
	Method:       http.MethodGet,
	Endpoint:     apiAthleteCompetitionRegistrationEndpoint,
	Handler:      athleteCompetitionRegistrationServer.GetAthleteRegistrationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

const apiCompetitionEntryEndpoint = "/api/v1.0/competition/entries"
const apiEventEntryEndpoint = "/api/v1.0/event/entries"
const apiAthleteEntryEndpoint = "/api/v1.0/athlete/entries"

var entryServer = controller.EntryServer{
	Service: businesslogic.NewCompetitionRegistrationService(
		database.AccountRepository,
		database.PartnershipRepository,
		database.CompetitionRepository,
		database.EventRepository,
		database.AthleteCompetitionEntryRepository,
		database.AthleteEventEntryRepository,
		database.PartnershipCompetitionEntryRepository,
		database.PartnershipEventEntryRepository,
	),
}
var searchCompetitionEntryController = util.DasController{
	Name:         "SearchEntryController",
	Description:  "Search Athlete/Partnership Competition/Event entries",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionEntryEndpoint,
	Handler:      entryServer.SearchCompetitionEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var searchEventEntryController = util.DasController{
	Name:         "SearchEntryController",
	Description:  "Search Athlete Competition Entries",
	Method:       http.MethodGet,
	Endpoint:     "/api/v1.0/entries/athlete/competition",
	Handler:      entryServer.SearchEventEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var searchCompetitionEntryByAthleteController = util.DasController{
	Name: "SearchEntryController",
	Description: `Search competition entries of a specific Athlete.
					This returns all the competitions that this Partnership have competed at`,
	Method:       http.MethodGet,
	Endpoint:     "/api/v1.0/entries/athlete/competition",
	Handler:      entryServer.SearchEventEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var searchCompetitionEntryByPartnershipController = util.DasController{
	Name: "SearchEntryController",
	Description: `Search competition entries of a specific Partnership.
					This returns all the competitions that this Partnership have competed at`,
	Method:       http.MethodGet,
	Endpoint:     "/api/v1.0/entries/partnership/competition",
	Handler:      entryServer.SearchEventEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var searchAthleteCompetitionEntryController = util.DasController{
	Name: "SearchAthleteCompetitionEntryController",
	Description: `Search entries of Athletes at a Competition.
					This returns all the Athletes (AthleteCompetitionEntry) who are competing at the specified competition`,
	Method:       http.MethodGet,
	Endpoint:     "/api/v1.0/entries/competition/athlete",
	Handler:      entryServer.SearchAthleteEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var searchPartnershipCompetitionEntryController = util.DasController{
	Name: "SearchPartnershipCompetitionEntryController",
	Description: `Search entries of Partnerships at a Competition.
					This returns all the couples (PartnershipCompetitionEntry) who are competing at the specified competition.`,
	Method:       http.MethodGet,
	Endpoint:     "/api/v1.0/entries/competition/partnership",
	Handler:      entryServer.SearchAthleteEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var searchAthleteEventEntryController = util.DasController{
	Name: "SearchAthleteEventEntryController",
	Description: `Search entries of Athletes at an Event of a Competition.
					This returns all the Athletes (AthleteEventEntry) who are competing at the specified event.`,
	Method:       http.MethodGet,
	Endpoint:     "/api/v1.0/entries/event/athlete",
	Handler:      entryServer.SearchAthleteEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

var searchPartnershipEventEntryController = util.DasController{
	Name: "SearchPartnershipEventEntryController",
	Description: `Search entries of Partnerships at an Event of a Competition.
					This returns all the couples (PartnershipEventEntry) who are competing at the specified event.`,
	Method:       http.MethodGet,
	Endpoint:     "/api/v1.0/entries/event/partnership",
	Handler:      entryServer.SearchPartnershipEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}

// CompetitionRegistrationControllerGroup is a collection of handler functions for managing
// Competition Registration in DAS
var CompetitionRegistrationControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createCompetitionRegistrationController,
		getPartnershipRegistrationController,
		searchCompetitionEntryController,
		searchEventEntryController,
		searchCompetitionEntryByAthleteController,
		searchCompetitionEntryByPartnershipController,
		searchAthleteCompetitionEntryController,
		searchPartnershipCompetitionEntryController,
		searchAthleteEventEntryController,
		searchPartnershipEventEntryController,
	},
}
