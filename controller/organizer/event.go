// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou

package organizer

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/controller/util/authentication"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type OrganizerEventServer struct {
	auth    authentication.IAuthenticationStrategy
	service businesslogic.OrganizerEventService
}

func (server OrganizerEventServer) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, _ := server.auth.GetCurrentUser(r)
	createDTO := new(viewmodel.CreateEventViewModel)

	if parseErr := util.ParseRequestBodyData(r, createDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, nil)
		return
	}

	event := createDTO.ToDomainModel(currentUser, danceRepository)
	if err := server.service.CreateEvent(event); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "success", nil)
}

// UpdateEventHandler handles the request
//	PUT /api/v1.0/organizer/event
func (server OrganizerEventServer) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// DeleteEventHandler handles the request:
//	DELETE /api/v1.0/organizer/event
func (server OrganizerEventServer) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// SearchEventHandler handles the request:
//	GET /api/v1.0/organizer/event
func (server OrganizerEventServer) SearchEventHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
