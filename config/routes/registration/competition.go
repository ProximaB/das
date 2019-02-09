package registration

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/DancesportSoftware/das/controller/athlete"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiAthleteCompetitionRegistrationEndpoint = "/api/v1.0/athlete/competition/registration"

var athleteCompetitionRegistrationServer = athlete.CompetitionRegistrationServer{
	IAuthenticationStrategy: middleware.AuthenticationStrategy,
	Service: businesslogic.CompetitionRegistrationService{
		AccountRepository:               database.AccountRepository,
		PartnershipRepository:           database.PartnershipRepository,
		CompetitionRepository:           database.CompetitionRepository,
		EventRepository:                 database.EventRepository,
		AthleteCompetitionEntryRepo:     database.AthleteCompetitionEntryRepository,
		PartnershipCompetitionEntryRepo: database.PartnershipCompetitionEntryRepository,
		PartnershipEventEntryRepo:       database.PartnershipEventEntryRepository,
	},
}

var createCompetitionRegistrationController = util.DasController{
	Name:         "CreateCompetitionRegistrationController",
	Description:  "Athlete creates competition and event registration",
	Method:       http.MethodPost,
	Endpoint:     apiAthleteCompetitionRegistrationEndpoint,
	Handler:      athleteCompetitionRegistrationServer.CreateAthleteRegistrationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

// CompetitionRegistrationControllerGroup is a collection of handler functions for managing
// Competition Registration in DAS
var CompetitionRegistrationControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createCompetitionRegistrationController,
	},
}
