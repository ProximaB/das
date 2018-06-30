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
