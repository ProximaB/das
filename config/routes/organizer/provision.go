package organizer

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/config/database"
	"github.com/ProximaB/das/config/routes/middleware"
	"github.com/ProximaB/das/controller/organizer"
	"github.com/ProximaB/das/controller/util"
	"net/http"
)

const apiOrganizerProvisionSummaryEndpoint = "/api/v1.0/organizer/provision/summary"
const apiOrganizerProvisionHistoryEndpoint = "/api/v1.0/organizer/provision/history"

var organizerProvisionServer = organizer.OrganizerProvisionServer{
	middleware.AuthenticationStrategy,
	database.AccountRepository,
	database.OrganizerProvisionRepository,
}

var getOrganizerProvisionSummaryController = util.DasController{
	Name:         "GetOrganizerProvisionSummaryController",
	Description:  "Retrieve organizer provision information for organizer",
	Method:       http.MethodGet,
	Endpoint:     apiOrganizerProvisionSummaryEndpoint,
	Handler:      organizerProvisionServer.GetOrganizerProvisionSummaryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var organizerProvisionHistoryServer = organizer.OrganizerProvisionHistoryServer{
	middleware.AuthenticationStrategy,
	database.AccountRepository,
	database.OrganizerProvisionHistoryRepository,
}
var getOrganizerProvisionHistoryController = util.DasController{
	Name:         "GetOrganizerProvisionHistoryController",
	Description:  "Retrieve organizer provision history for organizer",
	Method:       http.MethodGet,
	Endpoint:     apiOrganizerProvisionHistoryEndpoint,
	Handler:      organizerProvisionHistoryServer.GetOrganizerProvisionHistoryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}
var OrganizerProvisionControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		getOrganizerProvisionHistoryController,
		getOrganizerProvisionSummaryController,
	},
}
