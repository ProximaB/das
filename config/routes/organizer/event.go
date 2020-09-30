package organizer

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/config/routes/middleware"
	"github.com/ProximaB/das/controller/organizer"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiOrganizerEventEndpointV1_0 = "/api/v1.0/organizer/event"

var organizerEventService = businesslogic.NewOrganizerEventService(
	database.AccountRepository,
	database.AccountRoleRepository,
	database.CompetitionRepository,
	database.EventRepository,
	database.EventDanceRepository,
	database.CompetitionEventTemplateRepository,
	database.FederationRepository,
	database.DivisionRepository,
	database.AgeRepository,
	database.ProficiencyRepository,
	database.StyleRepository,
	database.DanceRepository)
var organizerEventServer = organizer.OrganizerEventServer{
	middleware.AuthenticationStrategy,
	organizerEventService,
}

var createEventController = util.DasController{
	Name:         "CreateEventController",
	Description:  "Organizer creates a event in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiOrganizerEventEndpointV1_0,
	Handler:      organizerEventServer.CreateEventHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var deleteEventController = util.DasController{
	Name:         "DeleteEventController",
	Description:  "Organizer deletes a event in DAS",
	Method:       http.MethodDelete,
	Endpoint:     apiOrganizerEventEndpointV1_0,
	Handler:      organizerEventServer.DeleteEventHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var searchEventController = util.DasController{
	Name:         "SearchEventController",
	Description:  "Organizer searches a event in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiOrganizerEventEndpointV1_0,
	Handler:      organizerEventServer.SearchEventHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var updateEventController = util.DasController{
	Name:         "UpdateEventController",
	Description:  "Organizer updates a event in DAS",
	Method:       http.MethodPut,
	Endpoint:     apiOrganizerEventEndpointV1_0,
	Handler:      organizerEventServer.UpdateEventHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var OrganizerEventManagementControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createEventController,
		deleteEventController,
		searchEventController,
		updateEventController,
	},
}

const apiOrganizeEventTemplateEndpoint = "/api/v1/organizer/event/template"

var searchCompetitionEventTemplateController = util.DasController{
	Name:         "SearchCompetitionEventTemplateController",
	Description:  "Template for populating competition events",
	Method:       http.MethodGet,
	Endpoint:     apiOrganizeEventTemplateEndpoint,
	Handler:      organizerEventServer.SearchCompetitionEventTemplateHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var createCompetitionEventTemplateController = util.DasController{
	Name:         "CreateCompetitionEventTemplateHandler",
	Description:  "Creating a template for particular user",
	Method:       http.MethodPost,
	Endpoint:     apiOrganizeEventTemplateEndpoint,
	Handler:      organizerEventServer.CreateCompetitionEventTemplateHanlder,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var OrganizerCompetitionEventTemplateControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createCompetitionEventTemplateController,
		searchCompetitionEventTemplateController,
	},
}
