package organizer

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/config/routes/middleware"
	"github.com/ProximaB/das/controller/organizer"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiOrganizerCompetitionEndpoint = "/api/v1.0/organizer/competition"

var organizerCompetitionServer = organizer.OrganizerCompetitionServer{
	IAuthenticationStrategy:              middleware.AuthenticationStrategy,
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
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var deleteCompetitionController = util.DasController{
	Name:         "DeleteCompetitionController",
	Description:  "Organizer delete a competition in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiOrganizerCompetitionEndpoint,
	Handler:      organizerCompetitionServer.OrganizerDeleteCompetitionHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var searchCompetitionController = util.DasController{
	Name:         "SearchCompetitionController",
	Description:  "Organizer searches a competition in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiOrganizerCompetitionEndpoint,
	Handler:      organizerCompetitionServer.OrganizerSearchCompetitionHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var updateCompetitionController = util.DasController{
	Name:         "UpdateCompetitionController",
	Description:  "Organizer updates a competition in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiOrganizerCompetitionEndpoint,
	Handler:      organizerCompetitionServer.OrganizerUpdateCompetitionHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var OrganizerCompetitionManagementControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createCompetitionController,
		updateCompetitionController,
		searchCompetitionController,
		deleteCompetitionController,
	},
}
