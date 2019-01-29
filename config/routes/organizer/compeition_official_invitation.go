// Dancesport Application System (DAS)
// Copyright (C) 2019 Yubing Hou
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
