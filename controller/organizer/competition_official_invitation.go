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
	"github.com/DancesportSoftware/das/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type CompetitionOfficialInvitationServer struct {
	auth.IAuthenticationStrategy
	businesslogic.IAccountRepository
	businesslogic.CompetitionOfficialInvitationService
}

// GET /api/v1/organizer/competition/official/invitation
func (server CompetitionOfficialInvitationServer) OrganizerGetCompetitionOfficialInvitationHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}

// POST /api/v1/organizer/competition/official/invitation
func (server CompetitionOfficialInvitationServer) OrganizerCreateCompetitionOfficialInvitationHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)

	createDTO := new(viewmodel.CreateCompetitionOfficialInvitationDTO)
	if parseErr := util.ParseRequestBodyData(r, createDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	//currentUser, _ := server.GetCurrentUser(r)
	// server.CompetitionOfficialInvitationService.CreateCompetitionOfficialInvitation(currentUser, )

	w.Write([]byte("not implemented"))
}

// PUT /api/v1/organizer/competition/official/invitation
func (server CompetitionOfficialInvitationServer) OrganizerUpdateCompetitionOfficialInvitationHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}
