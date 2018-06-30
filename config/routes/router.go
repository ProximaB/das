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
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/authentication"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/config/routes/internal/account"
	"github.com/DancesportSoftware/das/config/routes/internal/competition"
	"github.com/DancesportSoftware/das/config/routes/internal/organizer"
	"github.com/DancesportSoftware/das/config/routes/internal/partnership"
	"github.com/DancesportSoftware/das/config/routes/internal/reference"
	"github.com/DancesportSoftware/das/config/routes/internal/registration"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

func setResponseHeader(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cookie")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	}
}

// TODO: this part needs careful rework
func getRequestUserRole(r *http.Request) (int, error) {
	account, err := authentication.AuthenticationStrategy.GetCurrentUser(r, database.AccountRepository)
	if err != nil {
		return businesslogic.AccountTypeUnauthorized, err
	}
	return account.AccountTypeID, nil
}
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
		log.Fatalf("Endpoint of %s is required\n", handler.Endpoint)
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
		Handler(setResponseHeader(authorizeMultipleRoles(handler.Handler, handler.AllowedRoles)))
}
func authorizeMultipleRoles(h http.HandlerFunc, roles []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allowNoAuth := false
		for _, each := range roles {
			if each == businesslogic.AccountTypeNoAuth {
				allowNoAuth = true
				break
			}
		}

		userRole, authErr := getRequestUserRole(r)
		if authErr != nil && !allowNoAuth {
			util.RespondJsonResult(w, http.StatusUnauthorized, "invalid authorization token", nil)
			return
		}

		authorized := false
		for _, each := range roles {
			if each == userRole {
				authorized = true
			}
		}

		if authErr != nil && !allowNoAuth {
			util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
			return
		} else if allowNoAuth {
			h.ServeHTTP(w, r)
		} else if authorized {
			h.ServeHTTP(w, r)
		} else {
			util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
			return
		}
	}
}

func addDasControllerGroup(router *mux.Router, group util.DasControllerGroup) {
	for _, each := range group.Controllers {
		addDasController(router, each)
	}
}

// NewDasRouter creates a new router that handle requests in DAS
func NewDasRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Schemes("https")

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

	// partnership request blacklist
	addDasController(router, partnership.GetPartnershipBlacklistReasonController)
	addDasControllerGroup(router, partnership.PartnershipRequestBlacklistControllerGroup)

	// partnership request
	addDasController(router, partnership.PartnershipRequestStatusController)
	addDasControllerGroup(router, partnership.PartnershipRequestControllerGroup)

	// partnership
	addDasControllerGroup(router, partnership.PartnershipControllerGroup)

	// organizer (multi-user)
	addDasControllerGroup(router, organizer.ManageOrganizerProvisionControllerGroup)
	addDasControllerGroup(router, organizer.ProvisionControllerGroup)
	addDasControllerGroup(router, organizer.OrganizerProvisionControllerGroup)

	// organizer (only)
	addDasControllerGroup(router, organizer.OrganizerCompetitionManagementControllerGroup)

	// competition
	addDasController(router, competition.GetCompetitionStatusController)

	// athlete
	addDasControllerGroup(router, registration.CompetitionRegistrationControllerGroup)

	// scrutineer

	// emcee

	// deck captain

	// adjudicator

	// administrator

	// public only
	addDasControllerGroup(router, competition.PublicCompetitionViewControllerGroup)

	log.Println("[info] finishing controller initialization")
	return router
}
