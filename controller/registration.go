package controller

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/controller/util/authentication"
	"net/http"
)

type CompetitionRegistrationServer struct {
	businesslogic.IAccountRepository
	businesslogic.ICompetitionRepository
	businesslogic.ICompetitionEntryRepository
	businesslogic.IPartnershipRepository
	businesslogic.IEventRepository
	businesslogic.IEventEntryRepository
	authentication.IAuthenticationStrategy
}

// POST /api/athlete/registration
// This DasController is for athlete use only. Organizer will have to use a different DasController
func (server CompetitionRegistrationServer) CreateAthleteRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// validate identity first
	account, _ := server.GetCurrentUser(r, server.IAccountRepository)

	registrationDTO := new(businesslogic.CompetitiveBallroomEventRegistration)
	if parseErr := util.ParseRequestBodyData(r, registrationDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	// if registration is not valid, return error
	validationErr := businesslogic.ValidateCompetitiveBallroomEventRegistration(&account,
		registrationDTO,
		server.ICompetitionRepository,
		server.IEventRepository,
		server.ICompetitionEntryRepository,
		server.IAccountRepository,
		server.IPartnershipRepository)
	if validationErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, validationErr.Error(), nil)
		return
	}

	compEntryErr := businesslogic.CreateCompetitionEntry(&account, registrationDTO,
		server.ICompetitionEntryRepository,
		server.IAccountRepository,
		server.IPartnershipRepository)
	createEntryErr := businesslogic.CreateEventEntries(&account, registrationDTO, server.IEventEntryRepository)
	dropEventErr := businesslogic.DropEventEntries(&account, registrationDTO, server.IEventEntryRepository)

	if compEntryErr != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot enter competition", compEntryErr.Error())
		return
	}
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
	account, _ := server.GetCurrentUser(r, server.IAccountRepository)

	if account.ID == 0 || account.AccountTypeID != businesslogic.ACCOUNT_TYPE_ATHLETE {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	searchDTO := new(struct {
		CompetitionID int `schema:"competition"`
		PartnershipID int `schema:"partnership"`
	})

	if parseErr := util.ParseRequestData(r, searchDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	registration, err := businesslogic.GetEventRegistration(searchDTO.CompetitionID,
		searchDTO.PartnershipID, &account, server.IPartnershipRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	}

	output, _ := json.Marshal(registration)
	w.Write(output)
}
