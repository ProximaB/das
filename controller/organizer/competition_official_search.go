package organizer

import (
	"encoding/json"
	"github.com/ProximaB/das/auth"
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/controller/util"
	"github.com/ProximaB/das/viewmodel"
	"log"
	"net/http"
)

type OrganizerCompetitionOfficialSearchServer struct {
	auth.IAuthenticationStrategy
	businesslogic.IAccountRepository
	businesslogic.IAccountRoleRepository
}

// SearchEligibleOfficialHandler handles the request
//	GET /api/v1/organizer/official/eligible
//
// - Authorization: Organizer only
func (server OrganizerCompetitionOfficialSearchServer) SearchEligibleOfficialHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, err := server.IAuthenticationStrategy.GetCurrentUser(r)
	if err != nil || !currentUser.HasRole(businesslogic.AccountTypeOrganizer) {
		util.RespondJsonResult(w, http.StatusUnauthorized, "Not authorized", nil)
		return
	}

	// search based on the criteira
	criteriaDTO := new(viewmodel.SearchEligibleCompetitionOfficialDTO)
	if parseErr := util.ParseRequestData(r, criteriaDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "Bad request data. `accountTypeId` must be an integer", nil)
		return
	}

	if criteriaDTO.AccountTypeID == businesslogic.AccountTypeAthlete || criteriaDTO.AccountTypeID == businesslogic.AccountTypeAdministrator {
		util.RespondJsonResult(w, http.StatusBadRequest, "Only Adjudicator, Scrutineer, Deck Captain, and Emcee are searchable", nil)
		return
	}

	accounts, searchErr := server.IAccountRepository.SearchAccount(businesslogic.SearchAccountCriteria{AccountType: criteriaDTO.AccountTypeID})
	if searchErr != nil {
		log.Printf("[error] %v", searchErr)
		util.RespondJsonResult(w, http.StatusInternalServerError, "An internal error occurred. Please notify site administrator about this incident.", nil)
		return
	}

	data := make([]viewmodel.CompetitionOfficialProfileDTO, 0)
	for _, each := range accounts {
		dto := viewmodel.CompetitionOfficialProfileDTO{}
		dto.Populate(each)
		data = append(data, dto)
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}
