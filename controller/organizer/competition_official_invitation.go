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
