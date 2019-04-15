package organizer

import (
	"encoding/json"
	"fmt"
	"github.com/DancesportSoftware/das/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
)

// OrganizerEventServer serves requests regarding to event management
type OrganizerEventServer struct {
	Authentication auth.IAuthenticationStrategy
	Service        businesslogic.OrganizerEventService
}

// CreateEventHandler handles the request:
//	POST /api/v1.0/organizer/event
func (server OrganizerEventServer) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, _ := server.Authentication.GetCurrentUser(r)
	createDTO := new(viewmodel.CreateEventForm)

	if parseErr := util.ParseRequestBodyData(r, createDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, nil)
		return
	}

	event := createDTO.ToDomainModel(currentUser)
	if err := server.Service.CreateEvent(event); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "event is created", nil)
}

// UpdateEventHandler handles the request
//	PUT /api/v1.0/organizer/event
func (server OrganizerEventServer) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// DeleteEventHandler handles the request:
//	DELETE /api/v1.0/organizer/event
func (server OrganizerEventServer) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, _ := server.Authentication.GetCurrentUser(r)
	deleteDTO := new(viewmodel.DeleteEventForm)

	if parseErr := util.ParseRequestBodyData(r, deleteDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, nil)
		return
	} else if validErr := validator.Validate(deleteDTO); validErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, validErr.Error(), nil)
		return
	} else {
		events, searchErr := server.Service.SearchEvents(businesslogic.SearchEventCriteria{EventID: deleteDTO.ID})
		if searchErr != nil {
			util.RespondJsonResult(w, http.StatusInternalServerError, searchErr.Error(), nil)
			return
		}
		if len(events) != 1 {
			util.RespondJsonResult(w, http.StatusNotFound, "event with this ID does not exist ", nil)
			return
		}
		deletionErr := server.Service.DeleteEvent(events[0], currentUser)
		if deletionErr != nil {
			util.RespondJsonResult(w, http.StatusInternalServerError, deletionErr.Error(), nil)
			return
		}
		util.RespondJsonResult(w, http.StatusOK, fmt.Sprintf("Success: event with ID = %v is deleted", events[0].ID), nil)
		return
	}
}

// SearchEventHandler handles the request:
//	GET /api/v1.0/organizer/event
func (server OrganizerEventServer) SearchEventHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, _ := server.Authentication.GetCurrentUser(r)
	searchCriteriaDTO := new(viewmodel.OrganizerSearchEventCriteria)

	if parseErr := util.ParseRequestData(r, searchCriteriaDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, nil)
		return
	}

	searchCriteriaDTO.OrganizerID = currentUser.ID
	criteria := searchCriteriaDTO.ToBusinessModel()

	events, searchErr := server.Service.SearchEvents(criteria)
	if searchErr != nil {
		log.Printf("[error] searching event by organizer %v caught error: %v", currentUser.FullName(), searchErr)
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, nil)
		return
	}

	viewbag := make([]viewmodel.EventViewModel, 0)
	for _, each := range events {
		item := viewmodel.EventViewModel{}
		item.PopulateViewModel(each)
		viewbag = append(viewbag, item)
	}
	output, _ := json.Marshal(viewbag)
	w.Write(output)
}

// SearchCompetitionEventTemplateHandler handles the request:
//	GET /api/v1/organizer/event/template
func (server OrganizerEventServer) SearchCompetitionEventTemplateHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, _ := server.Authentication.GetCurrentUser(r)
	searchCriteriaDTO := new(viewmodel.SearchCompetitionEventTemplateForm)

	if parseErr := util.ParseRequestData(r, searchCriteriaDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, nil)
		return
	}

	searchCriteriaDTO.OwnerID = currentUser.ID
	results, _ := server.Service.SearchCompetitionEventTemplate(businesslogic.SearchCompetitionEventTemplateCriteria{
		ID:           searchCriteriaDTO.ID,
		Name:         searchCriteriaDTO.Name,
		CreateUserID: searchCriteriaDTO.OwnerID,
	})

	output, _ := json.Marshal(results)
	w.Write(output)
}

// SearchCompetitionEventTemplateHandler handles the request:
//	POST /api/v1/organizer/event/template
func (server OrganizerEventServer) CreateCompetitionEventTemplateHanlder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("Not implemented"))
}
