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

package routes

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/routes/account"
	"github.com/DancesportSoftware/das/config/routes/admin"
	"github.com/DancesportSoftware/das/config/routes/competition"
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/DancesportSoftware/das/config/routes/organizer"
	"github.com/DancesportSoftware/das/config/routes/partnership"
	"github.com/DancesportSoftware/das/config/routes/reference"
	"github.com/DancesportSoftware/das/config/routes/registration"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

/*
var restAPIRouter = []Route{
	// Competition

	// Events
	{"Public view of events", http.MethodGet, "/api/event", authorizeSingleRole(getEventHandler, businesslogic.AccountTypeNoAuth)},
	{"Public view of competitive ballroom events", http.MethodGet, "/api/event/competitive/ballroom", authorizeSingleRole(getCompetitiveBallroomEventHandler, businesslogic.AccountTypeNoAuth)},
	{"[Organizer] Create a competitive ballroom event", http.MethodPost, "/api/organizer/event", authorizeSingleRole(createEventHandler, businesslogic.AccountTypeOrganizer)},

	// Entries
	{"add/drop competitive ballroom event entries", http.MethodPost, "/api/athlete/registration", authorizeSingleRole(createAthleteRegistrationHandler, businesslogic.AccountTypeAthlete)},
	{"Get competitive ballroom entries for partnership", http.MethodGet, "/api/athlete/registration", authorizeSingleRole(getAthleteEventRegistrationHandler, businesslogic.AccountTypeAthlete)},
	{"Get competitive ballroom entries for public view", http.MethodGet, "/api/public/entries", authorizeSingleRole(getCompetitiveBallroomEventEntryHandler, businesslogic.AccountTypeNoAuth)},
}*/

// TODO: this part needs careful rework

func addDasController(router *mux.Router, handler util.DasController) {
	if len(handler.Name) < 1 {
		log.Fatalf("Name of %v is missing\n", handler)
	}
	if len(handler.Description) < 1 {
		log.Fatalf("Description of %s is required\n", handler.Name)
	}
	if len(handler.Method) < 1 {
		log.Fatalf("Method of %s is required\n", handler.Name)
	}
	if len(handler.Endpoint) < 1 {
		log.Fatalf("Endpoint of %s is required\n", handler.Name)
	}
	if handler.Handler == nil {
		log.Fatalf("HandlerFunc of %s is required\n", handler.Name)
	}
	if handler.AllowedRoles == nil {
		log.Fatalf("Alloed Roles of %s is required\n", handler.Name)
	}
	router.
		Methods(handler.Method, http.MethodOptions).
		Path(handler.Endpoint).
		Name(handler.Description).
		Handler(middleware.SetResponseHeader(middleware.AuthorizeMultipleRoles(handler.Handler, handler.AllowedRoles)))
}

func addDasControllerGroup(router *mux.Router, group util.DasControllerGroup) {
	for _, each := range group.Controllers {
		addDasController(router, each)
	}
}

const (
	EnvBuildDate = "BUILD_DATE"
)

func rootController(w http.ResponseWriter, r *http.Request) {
	buildDate := os.Getenv(EnvBuildDate)

	data := struct {
		BuildDate string `json:"version"`
	}{buildDate}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// NewDasRouter creates a new router that handle requests in DAS
func NewDasRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Schemes("https")

	addDasController(router, util.DasController{
		Name:         "RootController",
		Description:  "Handle Server Base Information",
		Endpoint:     "/",
		Handler:      rootController,
		Method:       http.MethodGet,
		AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
	})

	// reference data
	addDasControllerGroup(router, reference.CountryControllerGroup)
	addDasControllerGroup(router, reference.StateControllerGroup)
	addDasControllerGroup(router, reference.CityControllerGroup)
	addDasControllerGroup(router, reference.SchoolControllerGroup)
	addDasControllerGroup(router, reference.StudioControllerGroup)
	addDasControllerGroup(router, reference.FederationControllerGroup)
	addDasControllerGroup(router, reference.DivisionControllerGroup)
	addDasControllerGroup(router, reference.AgeControllerGroup)
	addDasControllerGroup(router, reference.ProficiencyControllerGroup)
	addDasControllerGroup(router, reference.StyleControllerGroup)
	addDasControllerGroup(router, reference.DanceControllerGroup)

	// account
	addDasControllerGroup(router, account.AccountControllerGroup)
	addDasController(router, account.AccountTypeController)
	addDasController(router, account.GenderController)
	addDasController(router, account.RoleController)
	addDasControllerGroup(router, account.UserPreferenceControllerGroup)
	addDasControllerGroup(router, account.RoleApplicationControllerGroup)

	// partnership request blacklist
	addDasController(router, partnership.GetPartnershipBlacklistReasonController)
	addDasControllerGroup(router, partnership.PartnershipRequestBlacklistControllerGroup)

	// partnership request
	addDasController(router, partnership.GetPartnershipRoleController)
	addDasController(router, partnership.PartnershipRequestStatusController)
	addDasControllerGroup(router, partnership.PartnershipRequestControllerGroup)

	// partnership
	addDasControllerGroup(router, partnership.PartnershipControllerGroup)

	// organizer (multi-user)
	addDasControllerGroup(router, organizer.OrganizerProvisionControllerGroup)

	// organizer (only)
	addDasControllerGroup(router, organizer.OrganizerCompetitionManagementControllerGroup)
	addDasControllerGroup(router, organizer.OrganizerEventManagementControllerGroup)
	addDasController(router, organizer.SearchEligibleCompetitionOfficialController)
	addDasControllerGroup(router, organizer.OrganizerCompetitionOfficialInvitationControllerGroup)

	// competition
	addDasController(router, competition.GetCompetitionStatusController)

	// athlete
	addDasControllerGroup(router, registration.CompetitionRegistrationControllerGroup)

	// scrutineer

	// emcee

	// deck captain

	// adjudicator

	// administrator
	addDasControllerGroup(router, admin.AdminManageUserControllerGroup)
	addDasControllerGroup(router, admin.ManageOrganizerProvisionControllerGroup)

	// public only
	addDasControllerGroup(router, competition.PublicCompetitionViewControllerGroup)

	log.Println("[info] finishing controller initialization")
	return router
}
