package organizer

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/DancesportSoftware/das/controller/organizer"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

var leadTagService = businesslogic.NewPartnershipCompetitionEntryService(database.AthleteCompetitionEntryRepository, database.PartnershipCompetitionEntryRepository)

var organizerLeadTagServer = organizer.NewOrganizerLeadTagServer(middleware.AuthenticationStrategy, database.CompetitionRepository, leadTagService)

const apiOrganizerLeadTagEndpointV1_0 = "/api/v1.0/organizer/competition/leads"

var getAllLeadController = util.DasController{
	Name:         "GetAllLeadsController",
	Description:  "Get all the leads of a competition",
	Method:       http.MethodGet,
	Endpoint:     apiOrganizerLeadTagEndpointV1_0,
	Handler:      organizerLeadTagServer.GetAllLeadEntries,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var OrganizerLeadTagManagementControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		getAllLeadController,
	},
}
