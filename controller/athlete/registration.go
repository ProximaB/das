package athlete

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DancesportSoftware/das/auth"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"

	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
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
	auth.IAuthenticationStrategy
	Service businesslogic.CompetitionRegistrationService
}

// prepareRegistrationForm is a helper function that sanitizes the submitted form for validation
func (server CompetitionRegistrationServer) prepareRegistrationForm(dto viewmodel.AthleteCompetitionRegistrationForm) (businesslogic.EventRegistrationForm, error) {
	form := businesslogic.EventRegistrationForm{}

	competitions, findCompErr := server.ICompetitionRepository.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: dto.CompetitionID})
	if findCompErr != nil || len(competitions) != 1 {
		return form, errors.New(fmt.Sprintf("cannot find competition with ID = %v", dto.CompetitionID))
	}
	form.Competition = competitions[0]

	partnerships, findCoupleErr := server.IPartnershipRepository.SearchPartnership(businesslogic.SearchPartnershipCriteria{PartnershipID: dto.PartnershipID})
	if findCoupleErr != nil || len(partnerships) != 1 {
		return form, errors.New(fmt.Sprintf("cannot find parntership with ID = %v", dto.PartnershipID))
	}
	form.Couple = partnerships[0]

	addedEvents := make([]businesslogic.Event, 0)
	droppedEvents := make([]businesslogic.Event, 0)

	for _, each := range dto.AddedEvents {
		evts, searchErr := server.IEventRepository.SearchEvent(businesslogic.SearchEventCriteria{EventID: each})
		if searchErr != nil {
			return form, nil
		}
		if len(evts) != 1 {
			return form, errors.New(fmt.Sprintf("event with ID = %v cannot be found", each))
		}
		addedEvents = append(addedEvents, evts[0])
	}
	form.EventsAdded = addedEvents

	for _, each := range dto.DroppedEvents {
		evts, searchErr := server.IEventRepository.SearchEvent(businesslogic.SearchEventCriteria{EventID: each})
		if searchErr != nil {
			return form, nil
		}
		if len(evts) != 1 {
			return form, errors.New(fmt.Sprintf("event with ID = %v cannot be found", each))
		}
		droppedEvents = append(droppedEvents, evts[0])
	}
	form.EventsDropped = droppedEvents

	form.CountryRepresented.ID = dto.Representation.CountryId
	form.StateRepresented.ID = dto.Representation.StateId
	form.SchoolRepresented.ID = dto.Representation.SchoolId
	form.StudioRepresented.ID = dto.Representation.StudioId

	return form, nil
}

// CreateAthleteRegistrationHandler handles the request
//	POST /api/v1.0/competition/registration
// This DasController is for athlete use only. Organizer will have to use a different DasController
func (server CompetitionRegistrationServer) CreateAthleteRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// validate identity first
	account, _ := server.GetCurrentUser(r)

	registrationDTO := new(viewmodel.AthleteCompetitionRegistrationForm)
	if parseErr := util.ParseRequestBodyData(r, registrationDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	form, formErr := server.prepareRegistrationForm(*registrationDTO)
	if formErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, formErr.Error(), nil)
		return
	}

	validationErr := server.Service.CreateAndUpdateRegistration(account, form)

	// if registration is not valid, return error
	if validationErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, validationErr.Error(), nil)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "event entries have been successfully added and/or dropped", nil)
}

// GET /api/athlete/registration
// This DasController is for athlete use only. Organizer will have to use a different DasController
// THis is not for public view. For public view, see getCompetitiveBallroomEventEntryHandler()
func (server CompetitionRegistrationServer) GetAthleteRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	account, authErr := server.GetCurrentUser(r)

	if account.ID == 0 || !account.HasRole(businesslogic.AccountTypeAthlete) || authErr != nil {
		util.RespondJsonResult(w, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	searchDTO := new(struct {
		CompetitionID int `schema:"competitionId"`
		PartnershipID int `schema:"partnershipId"`
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
