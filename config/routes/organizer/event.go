// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package organizer

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/DancesportSoftware/das/controller/organizer"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiOrganizerEventEndpointV1_0 = "/api/v1.0/organizer/event"

var organizerEventService = businesslogic.NewOrganizerEventService(
	database.AccountRepository,
	database.AccountRoleRepository,
	database.CompetitionRepository,
	database.EventRepository,
	database.EventDanceRepository)
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
