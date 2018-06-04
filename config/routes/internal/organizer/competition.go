package organizer

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/organizer"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiOrganizerCompetitionEndpoint = "/api/organizer/competition"

var organizerCompetitionServer = organizer.OrganizerCompetitionServer{
	IAccountRepository:                   database.AccountRepository,
	ICompetitionRepository:               database.CompetitionRepository,
	IOrganizerProvisionRepository:        database.OrganizerProvisionRepository,
	IOrganizerProvisionHistoryRepository: database.OrganizerProvisionHistoryRepository,
}

var createCompetitionController = util.DasController{
	Name:         "CreateCompetitionController",
	Description:  "Organizer creates a competition in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiOrganizerCompetitionEndpoint,
	Handler:      organizerCompetitionServer.OrganizerCreateCompetitionHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ORGANIZER},
}

var deleteCompetitionController = util.DasController{
	Name:         "DeleteCompetitionController",
	Description:  "Organizer delete a competition in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiOrganizerCompetitionEndpoint,
	Handler:      organizerCompetitionServer.OrganizerDeleteCompetitionHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ORGANIZER},
}

var searchCompetitionController = util.DasController{
	Name:         "SearchCompetitionController",
	Description:  "Organizer searches a competition in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiOrganizerCompetitionEndpoint,
	Handler:      organizerCompetitionServer.OrganizerSearchCompetitionHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ORGANIZER},
}

var updateCompetitionController = util.DasController{
	Name:         "UpdateCompetitionController",
	Description:  "Organizer updates a competition in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiOrganizerCompetitionEndpoint,
	Handler:      organizerCompetitionServer.OrganizerUpdateCompetitionHandler,
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ORGANIZER},
}

var OrganizerCompetitionManagementControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createCompetitionController,
		updateCompetitionController,
		searchCompetitionController,
		deleteCompetitionController,
	},
}
