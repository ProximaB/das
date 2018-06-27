// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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
	AllowedRoles: []int{businesslogic.ACCOUNT_TYPE_ATHLETE},
}

// CompetitionRegistrationControllerGroup is a collection of handler functions for managing
// Competition Registration in DAS
var CompetitionRegistrationControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createCompetitionRegistrationController,
	},
}
