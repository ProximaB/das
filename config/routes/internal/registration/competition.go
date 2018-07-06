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

package registration

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/authentication"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiAthleteCompetitionRegistrationEndpoint = "/api/athlete/competition/registration"

var athleteCompetitionRegistrationServer = controller.CompetitionRegistrationServer{
	database.AccountRepository,
	database.CompetitionRepository,
	database.AthleteCompetitionEntryRepository,
	database.PartnershipCompetitionEntryRepository,
	database.PartnershipRepository,
	database.EventRepository,
	database.PartnershipEventEntryRepository,
	authentication.AuthenticationStrategy,
}

var createCompetitionRegistrationController = util.DasController{
	Name:         "CreateCompetitionRegistrationController",
	Description:  "Create competition and event registration in DAS",
	Method:       http.MethodPost,
	Endpoint:     apiAthleteCompetitionRegistrationEndpoint,
	Handler:      athleteCompetitionRegistrationServer.CreateAthleteRegistrationHandler,
	AllowedRoles: []int{businesslogic.AccountTypeAthlete},
}

// CompetitionRegistrationControllerGroup is a collection of handler functions for managing
// Competition Registration in DAS
var CompetitionRegistrationControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createCompetitionRegistrationController,
	},
}
