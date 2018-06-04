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
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	sender, _ := server.GetCurrentUser(r, server.IAccountRepository)
	recipient := businesslogic.GetAccountByEmail(dto.RecipientEmail, server.IAccountRepository)

	if recipient.ID == 0 {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, "recipient does not exist")
		return
	}

	request := businesslogic.PartnershipRequest{
		SenderID:        sender.ID,
		RecipientID:     recipient.ID,
		RecipientRole:   dto.RecipientRole,
		Message:         dto.Message,
		Status:          businesslogic.PARTNERSHIP_REQUEST_STATUS_PENDING,
		CreateUserID:    sender.ID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    sender.ID,
	}

	if request.RecipientRole == businesslogic.PARTNERSHIP_ROLE_LEAD {
		request.SenderRole = businesslogic.PARTNERSHIP_ROLE_FOLLOW
	} else if request.RecipientRole == businesslogic.PARTNERSHIP_ROLE_FOLLOW {
		request.SenderRole = businesslogic.PARTNERSHIP_ROLE_LEAD
	} else {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, "invalid role for recipient")
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
	account, _ := server.GetCurrentUser(r, server.IAccountRepository)
	criteria := new(businesslogic.SearchPartnershipRequestCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	if criteria.Type == businesslogic.PARTNERSHIP_REQUEST_RECEIVED {
		criteria.Recipient = account.ID
	} else if criteria.Type == businesslogic.PARTNERSHIP_REQUEST_SENT {
		criteria.Sender = account.ID
	} else {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, "invalid partnership request type")
		return
	}

	requests, err := server.SearchPartnershipRequest(*criteria)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	}

	data := make([]viewmodel.PartnershipRequest, 0)
	for _, each := range requests {
		data = append(data, viewmodel.PartnershipRequest{
			ID:              each.PartnershipRequestID,
			Sender:          businesslogic.GetAccountByID(each.SenderID, server.IAccountRepository).GetName(),
			Recipient:       businesslogic.GetAccountByID(each.RecipientID, server.IAccountRepository).GetName(),
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
	currentUser, _ := server.GetCurrentUser(r, server.IAccountRepository)

	respondDTO := new(viewmodel.PartnershipRequestResponse)
	if parseErr := util.ParseRequestBodyData(r, respondDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	if respondDTO.Response != businesslogic.PARTNERSHIP_REQUEST_STATUS_ACCEPTED && respondDTO.Response != businesslogic.PARTNERSHIP_REQUEST_STATUS_DECLINED {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, "invalid response")
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
