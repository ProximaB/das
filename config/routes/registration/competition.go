package registration

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/DancesportSoftware/das/controller"
	"github.com/DancesportSoftware/das/controller/athlete"
	"github.com/DancesportSoftware/das/controller/util"
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
const apiPartnershipEntryEndpoint = "/api/v1.0/partnership/entries"

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
	Description:  "Search Athlete/Partnership Competition/Event entries",
	Method:       http.MethodGet,
	Endpoint:     apiEventEntryEndpoint,
	Handler:      entryServer.SearchEventEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
var searchAthleteEntryController = util.DasController{
	Name:         "SearchEntryController",
	Description:  "Search Athlete/Partnership Competition/Event entries",
	Method:       http.MethodGet,
	Endpoint:     apiAthleteEntryEndpoint,
	Handler:      entryServer.SearchAthleteEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
var searchPartnershipEntryController = util.DasController{
	Name:         "SearchEntryController",
	Description:  "Search Athlete/Partnership Competition/Event entries",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipEntryEndpoint,
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
		searchAthleteEntryController,
		searchPartnershipEntryController,
	},
}
