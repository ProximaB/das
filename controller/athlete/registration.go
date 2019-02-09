package athlete

import (
	"encoding/json"
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
