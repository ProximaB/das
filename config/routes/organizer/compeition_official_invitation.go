package organizer

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/config/routes/middleware"
	"github.com/ProximaB/das/controller/organizer"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiOrganizerCompetitionOfficialInvitation = "/api/v1/organizer/competition/official/invitation"

var competitionOfficialInvitationService = businesslogic.NewCompetitionOfficialInvitationService(database.AccountRepository, database.CompetitionRepository, database.CompetitionOfficialRepository, database.CompetitionOfficialInvitationRepository)

var organzierCompetitionOfficialInvitationServer = organizer.CompetitionOfficialInvitationServer{
	middleware.AuthenticationStrategy,
	database.AccountRepository,
	competitionOfficialInvitationService,
}

var createCompetitionOfficialInvitationController = util.DasController{
	Name:         "CreateCompetitionOfficialInvitationController",
	Description:  "Organizer creates an invitation for competition official",
	Method:       http.MethodPost,
	Endpoint:     apiOrganizerCompetitionOfficialInvitation,
	Handler:      organzierCompetitionOfficialInvitationServer.OrganizerCreateCompetitionOfficialInvitationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var searchCompetitionOfficialInvitationController = util.DasController{
	Name:         "SearchCompetitionOfficialInvitationController",
	Description:  "Organizer creates an invitation for competition official",
	Method:       http.MethodGet,
	Endpoint:     apiOrganizerCompetitionOfficialInvitation,
	Handler:      organzierCompetitionOfficialInvitationServer.OrganizerCreateCompetitionOfficialInvitationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var updateCompetitionOfficialInvitationController = util.DasController{
	Name:         "UpdateCompetitionOfficialInvitationController",
	Description:  "Organizer creates an invitation for competition official",
	Method:       http.MethodPut,
	Endpoint:     apiOrganizerCompetitionOfficialInvitation,
	Handler:      organzierCompetitionOfficialInvitationServer.OrganizerCreateCompetitionOfficialInvitationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var OrganizerCompetitionOfficialInvitationControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createCompetitionOfficialInvitationController,
		searchCompetitionOfficialInvitationController,
		updateCompetitionOfficialInvitationController,
	},
}
