package request

import (
	"encoding/json"
	"github.com/ProximaB/das/auth"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/controller/util"
	"github.com/ProximaB/das/viewmodel"
	"log"
	"net/http"
	"time"
)

// PartnershipRequestServer serves requests that are related to Partnership Requests
type PartnershipRequestServer struct {
	auth.IAuthenticationStrategy
	businesslogic.IAccountRepository
	businesslogic.IPartnershipRepository
	businesslogic.IPartnershipRequestRepository
	businesslogic.IPartnershipRequestBlacklistRepository
}

// CreatePartnershipRequestHandler handles the request:
//	POST /athlete/partnership/request
// which allows user to submit a new partnership request
func (server PartnershipRequestServer) CreatePartnershipRequestHandler(w http.ResponseWriter, r *http.Request) {
	dto := new(viewmodel.CreatePartnershipRequest)

	if parseErr := util.ParseRequestBodyData(r, dto); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	sender, _ := server.GetCurrentUser(r)
	searchResults, err := server.SearchAccount(businesslogic.SearchAccountCriteria{Email: dto.RecipientEmail})
	if len(searchResults) == 0 {
		util.RespondJsonResult(w, http.StatusNotFound, util.Http404NoDataFound, nil)
		return
	}
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}
	recipient := searchResults[0]

	if recipient.ID == 0 {
		util.RespondJsonResult(w, http.StatusBadRequest, "recipient does not exist", nil)
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
		DateTimeUpdated: time.Now(),
	}

	if request.RecipientRole == businesslogic.PartnershipRoleLead {
		request.SenderRole = businesslogic.PartnershipRoleFollow
	} else if request.RecipientRole == businesslogic.PartnershipRoleFollow {
		request.SenderRole = businesslogic.PartnershipRoleLead
	} else {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, "invalid role for recipient")
		return
	}

	err = businesslogic.CreatePartnershipRequest(request, server.IPartnershipRepository,
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
		eachReq := viewmodel.PartnershipRequest{
			ID:              each.PartnershipRequestID,
			Sender:          businesslogic.GetAccountByID(each.SenderID, server.IAccountRepository).FullName(),
			Recipient:       businesslogic.GetAccountByID(each.RecipientID, server.IAccountRepository).FullName(),
			Message:         each.Message,
			Status:          each.Status,
			DateTimeCreated: each.DateTimeCreated,
		}
		if each.SenderRole == businesslogic.PartnershipRoleLead {
			eachReq.Role = "Lead"
		} else {
			eachReq.Role = "Follow"
		}
		data = append(data, eachReq)
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// UpdatePartnershipRequestHandler handles the request
//	PUT /api/partnership/request
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
		log.Printf("[error] responding to partnership request failed: %v", err)
		util.RespondJsonResult(w, http.StatusInternalServerError, "error in responding partnership request", err.Error())
		return
	}

	util.RespondJsonResult(w, http.StatusOK, "response is sent", nil)
}

// DeletePartnershipRequestHandler handles the request
//	DELETE /api/v1.0/partnership
func (server PartnershipRequestServer) DeletePartnershipRequestHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusMethodNotAllowed, "method is not allowed", nil)
}
