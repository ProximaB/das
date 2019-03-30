package organizer

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/organizer"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiOrganizerEntryEndpointV1_0 = "/api/v1.0/organizer/competition/entry"

var organizerEntryServer = organizer.OrganizerEntryServer{}

var createEntryController = util.DasController{
	Name:         "CreateEntryController",
	Description:  "Organizer creates an entry",
	Method:       http.MethodPost,
	Endpoint:     apiOrganizerEntryEndpointV1_0,
	Handler:      organizerEntryServer.CreateEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var deleteEntryController = util.DasController{
	Name:         "DeleteEntryController",
	Description:  "Organizer deletes an entry",
	Method:       http.MethodDelete,
	Endpoint:     apiOrganizerEntryEndpointV1_0,
	Handler:      organizerEntryServer.DeleteEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var searchEntryController = util.DasController{
	Name:         "SearchEntryController",
	Description:  "Organizer searches an entry",
	Method:       http.MethodGet,
	Endpoint:     apiOrganizerEntryEndpointV1_0,
	Handler:      organizerEntryServer.SearchEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var updateEntryController = util.DasController{
	Name:         "UpdateEntryController",
	Description:  "Organizer updates an entry",
	Method:       http.MethodPut,
	Endpoint:     apiOrganizerEntryEndpointV1_0,
	Handler:      organizerEntryServer.UpdateEntryHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}

var OrganizerEntryManagementControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createEntryController,
		deleteEntryController,
		searchEntryController,
		updateEntryController,
	},
}
