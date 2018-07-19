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

package request

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/controller/util/authentication"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
	"time"
)

// PartnershipRequestServer serves requests that are related to Partnership Requests
type PartnershipRequestServer struct {
	authentication.IAuthenticationStrategy
	businesslogic.IAccountRepository
	businesslogic.IPartnershipRepository
	businesslogic.IPartnershipRequestRepository
	businesslogic.IPartnershipRequestBlacklistRepository
}

// POST /api/partnership/request
// Submit a new partnership request
func (server PartnershipRequestServer) CreatePartnershipRequestHandler(w http.ResponseWriter, r *http.Request) {
	dto := new(viewmodel.CreatePartnershipRequest)

	if parseErr := util.ParseRequestBodyData(r, dto); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	sender, _ := server.GetCurrentUser(r)
	recipient := businesslogic.GetAccountByEmail(dto.RecipientEmail, server.IAccountRepository)

	if recipient.ID == 0 {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, "recipient does not exist")
		return
	}

	request := businesslogic.PartnershipRequest{
		SenderID:        sender.ID,
		RecipientID:     recipient.ID,
		RecipientRole:   dto.RecipientRole,
		Message:         dto.Message,
		Status:          businesslogic.PartnershipRequestStatusPending,
		CreateUserID:    sender.ID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    sender.ID,
	}

	if request.RecipientRole == businesslogic.PartnershipRoleLead {
		request.SenderRole = businesslogic.PartnershipRoleFollow
	} else if request.RecipientRole == businesslogic.PartnershipRoleFollow {
		request.SenderRole = businesslogic.PartnershipRoleLead
	} else {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, "invalid role for recipient")
		return
	}

	err := businesslogic.CreatePartnershipRequest(request, server.IPartnershipRepository,
		server.IPartnershipRequestRepository, server.IAccountRepository, server.IPartnershipRequestBlacklistRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "error in submitting partnership request", err.Error())
		return
	}

	util.RespondJsonResult(w, http.StatusOK, "success", nil)
	return
}

// GET /api/partnership/request
// Get a list of received partnership requests
func (server PartnershipRequestServer) SearchPartnershipRequestHandler(w http.ResponseWriter, r *http.Request) {
	account, _ := server.GetCurrentUser(r)
	criteria := new(businesslogic.SearchPartnershipRequestCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	if criteria.Type == businesslogic.PartnershipRequestReceived {
		criteria.Recipient = account.ID
	} else if criteria.Type == businesslogic.PartnershipRequestSent {
		criteria.Sender = account.ID
	} else {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, "invalid partnership request type")
		return
	}

	requests, err := server.SearchPartnershipRequest(*criteria)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]viewmodel.PartnershipRequest, 0)
	for _, each := range requests {
		data = append(data, viewmodel.PartnershipRequest{
			ID:              each.PartnershipRequestID,
			Sender:          businesslogic.GetAccountByID(each.SenderID, server.IAccountRepository).FullName(),
			Recipient:       businesslogic.GetAccountByID(each.RecipientID, server.IAccountRepository).FullName(),
			Message:         each.Message,
			Status:          each.Status,
			DateTimeCreated: each.DateTimeCreated,
			Role:            each.RecipientRole,
		})
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// PUT /api/partnership/request
func (server PartnershipRequestServer) UpdatePartnershipRequestHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, _ := server.GetCurrentUser(r)

	respondDTO := new(viewmodel.PartnershipRequestResponse)
	if parseErr := util.ParseRequestBodyData(r, respondDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	if respondDTO.Response != businesslogic.PartnershipRequestStatusAccepted && respondDTO.Response != businesslogic.PartnershipRequestStatusDeclined {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, "invalid response")
		return
	}

	response := businesslogic.PartnershipRequestResponse{
		RequestID:       respondDTO.RequestID,
		Response:        respondDTO.Response,
		RecipientID:     currentUser.ID,
		DateTimeCreated: time.Now(),
	}

	err := businesslogic.RespondPartnershipRequest(response, server.IPartnershipRequestRepository, server.IAccountRepository, server.IPartnershipRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "error in responding partnership request", err.Error())
		return
	}

	util.RespondJsonResult(w, http.StatusOK, "partnership request responded", nil)
}

func (server PartnershipRequestServer) DeletePartnershipRequestHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}

// GET /api/v1/partnership/role
func (server PartnershipRequestServer) PartnershipRoleHandler(w http.ResponseWriter, r *http.Request) {

}
