package organizer

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"log"
	"net/http"
	"time"
)

type SearchOrganizerCompetitionViewModel struct {
	ID     int  `schema:"id"`
	Future bool `schema:"future"`
}

type OrganizerCompetitionServer struct {
	auth.IAuthenticationStrategy
	businesslogic.IAccountRepository
	businesslogic.ICompetitionRepository
	businesslogic.IOrganizerProvisionRepository
	businesslogic.IOrganizerProvisionHistoryRepository
}

// POST /api/organizer/competition
func (server OrganizerCompetitionServer) OrganizerCreateCompetitionHandler(w http.ResponseWriter, r *http.Request) {
	createDTO := new(viewmodel.CreateCompetition)
	if err := util.ParseRequestBodyData(r, createDTO); err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, err.Error())
		return
	}

	account, _ := server.GetCurrentUser(r)
	competition := createDTO.ToCompetitionDataModel(account)

	err := businesslogic.CreateCompetition(competition, server.ICompetitionRepository, server.IOrganizerProvisionRepository, server.IOrganizerProvisionHistoryRepository)
	if err != nil {
		log.Printf("cannot create competition %v", err)
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "success", nil)
}

// GET /api/organizer/competition
func (server OrganizerCompetitionServer) OrganizerSearchCompetitionHandler(w http.ResponseWriter, r *http.Request) {
	searchDTO := new(SearchOrganizerCompetitionViewModel)
	if parseErr := util.ParseRequestData(r, searchDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
	} else {
		account, _ := server.GetCurrentUser(r)
		if account.ID == 0 ||
			(!account.HasRole(businesslogic.AccountTypeOrganizer) &&
				!account.HasRole(businesslogic.AccountTypeAdministrator)) {
			util.RespondJsonResult(w, http.StatusUnauthorized, "you are not authorized to look up this information", nil)
			return
		}
		criteria := businesslogic.SearchCompetitionCriteria{
			ID:          searchDTO.ID,
			OrganizerID: account.ID,
		}
		if searchDTO.Future {
			criteria.StartDateTime = time.Now()
		}

		comps, err := server.SearchCompetition(criteria)
		if err != nil {
			util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
			return
		} else {
			data := make([]viewmodel.Competition, 0)
			for _, each := range comps {
				data = append(data, viewmodel.CompetitionDataModelToViewModel(each, businesslogic.AccountTypeOrganizer))
			}
			output, _ := json.Marshal(data)
			w.Write(output)
		}
	}
}

// PUT /api/organizer/competition
func (server OrganizerCompetitionServer) OrganizerUpdateCompetitionHandler(w http.ResponseWriter, r *http.Request) {
	account, _ := server.GetCurrentUser(r)
	updateDTO := new(businesslogic.OrganizerUpdateCompetition)

	if parseErr := util.ParseRequestBodyData(r, updateDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, nil)
		return
	}

	competitions, _ := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: updateDTO.CompetitionID})
	if competitions[0].CreateUserID == account.ID {
		competitions[0].Street = updateDTO.Address
		competitions[0].UpdateStatus(updateDTO.Status) // TODO; error prone
		competitions[0].DateTimeUpdated = time.Now()
		competitions[0].UpdateUserID = account.ID
	} else {
		util.RespondJsonResult(w, http.StatusUnauthorized, "not authorized to make changes to this competition", nil)
		return
	}

	if updateErr := server.UpdateCompetition(competitions[0]); updateErr != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, updateErr.Error())
		return
	}

	util.RespondJsonResult(w, http.StatusOK, "competition is updated", nil)
	return
}

// DELETE /api/organizer/competition
func (server OrganizerCompetitionServer) OrganizerDeleteCompetitionHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}
