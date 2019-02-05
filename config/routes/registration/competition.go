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
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/DancesportSoftware/das/controller/athlete"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiAthleteCompetitionRegistrationEndpoint = "/api/v1.0/athlete/competition/registration"

var athleteCompetitionRegistrationServer = athlete.CompetitionRegistrationServer{
	IAuthenticationStrategy: middleware.AuthenticationStrategy,
	Service: businesslogic.CompetitionRegistrationService{
		AccountRepository:               database.AccountRepository,
		PartnershipRepository:           database.PartnershipRepository,
		CompetitionRepository:           database.CompetitionRepository,
		EventRepository:                 database.EventRepository,
		AthleteCompetitionEntryRepo:     database.AthleteCompetitionEntryRepository,
		PartnershipCompetitionEntryRepo: database.PartnershipCompetitionEntryRepository,
		PartnershipEventEntryRepo:       database.PartnershipEventEntryRepository,
	},
}

var createCompetitionRegistrationController = util.DasController{
	Name:         "CreateCompetitionRegistrationController",
	Description:  "Athlete creates competition and event registration",
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
