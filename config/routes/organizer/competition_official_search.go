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

const apiOrganizerCompeitionOfficialSearch = "/api/v1/organizer/competition/official/eligible"

var organizerCompetitionOfficialSearchServer = organizer.OrganizerCompetitionOfficialSearchServer{
	IAuthenticationStrategy: middleware.AuthenticationStrategy,
	IAccountRepository:      database.AccountRepository,
	IAccountRoleRepository:  database.AccountRoleRepository,
}

var SearchEligibleCompetitionOfficialController = util.DasController{
	Name:         "SearchEligibleCompetitionOfficialController",
	Description:  "Organzier search eligible officials for competition",
	Method:       http.MethodGet,
	Endpoint:     apiOrganizerCompeitionOfficialSearch,
	Handler:      organizerCompetitionOfficialSearchServer.SearchEligibleOfficialHandler,
	AllowedRoles: []int{businesslogic.AccountTypeOrganizer},
}
