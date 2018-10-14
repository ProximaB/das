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

package controller

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"

	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/controller/util/authentication"
)

// CompetitionRegistrationServer handles requests that create or update competition registrations
type CompetitionRegistrationServer struct {
	businesslogic.IAccountRepository
	businesslogic.ICompetitionRepository
	businesslogic.IAthleteCompetitionEntryRepository
	businesslogic.IPartnershipCompetitionEntryRepository
	businesslogic.IPartnershipRepository
	businesslogic.IEventRepository
	businesslogic.IPartnershipEventEntryRepository
	authentication.IAuthenticationStrategy
	Service businesslogic.CompetitionRegistrationService
}

// CreateAthleteRegistrationHandler handles the request
//	POST /api/v1.0/competition/registration
// This DasController is for athlete use only. Organizer will have to use a different DasController
func (server CompetitionRegistrationServer) CreateAthleteRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// validate identity first
	account, _ := server.GetCurrentUser(r)

	registrationDTO := new(viewmodel.SubmitCompetitionRegistrationForm)
	if parseErr := util.ParseRequestBodyData(r, registrationDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	form := registrationDTO.EventRegistration()

	validationErr := server.Service.ValidateEventRegistration(account, form)

	// if registration is not valid, return error
	if validationErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, validationErr.Error(), nil)
		return
	}

	server.Service.CreateAthleteCompetitionEntry(account, form)

	createEntryErr := server.Service.CreatePartnershipEventEntries(account, form)
	dropEventErr := server.Service.DropPartnershipEventEntries(account, form)

	if createEntryErr != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "error in creating event entry", createEntryErr.Error())
		return
	}

	if dropEventErr != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "error in dropping event entry", dropEventErr.Error())
		return
	}

	util.RespondJsonResult(w, http.StatusOK, "event entries have been successfully added and/or dropped", nil)
}

// GET /api/athlete/registration
// This DasController is for athlete use only. Organizer will have to use a different DasController
// THis is not for public view. For public view, see getCompetitiveBallroomEventEntryHandler()
func (server CompetitionRegistrationServer) GetAthleteEventRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	account, _ := server.GetCurrentUser(r)

	if account.ID == 0 || !account.HasRole(businesslogic.AccountTypeAthlete) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	searchDTO := new(struct {
		CompetitionID int `schema:"competition"`
		PartnershipID int `schema:"partnership"`
	})

	if parseErr := util.ParseRequestData(r, searchDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	registration, err := businesslogic.GetEventRegistration(searchDTO.CompetitionID,
		searchDTO.PartnershipID, &account, server.IPartnershipRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	output, _ := json.Marshal(registration)
	w.Write(output)
}
