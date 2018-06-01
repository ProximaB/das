package organizer

import (
	"github.com/yubing24/das/businesslogic"
	"github.com/yubing24/das/controller/util"
	"github.com/yubing24/das/viewmodel"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const apiOrganizerCompetitionEndpoint = "/api/organizer/competition"

var createCompetitionController = util.DasController{}
var updateCompetitionController = util.DasController{}
var searchCompetitionController = util.DasController{}
var deleteCompetitionController = util.DasController{}

var OrganizerCompetitionManagementControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		createCompetitionController,
		updateCompetitionController,
		searchCompetitionController,
		deleteCompetitionController,
	},
}

type SearchOrganizerCompetitionViewModel struct {
	Future bool `schema:"future"`
}

type OrganizerCompetitionServer struct {
	businesslogic.IAccountRepository
	businesslogic.ICompetitionRepository
	businesslogic.IOrganizerProvisionRepository
	businesslogic.IOrganizerProvisionHistoryRepository
}

// POST /api/organizer/competition
func (server OrganizerCompetitionServer) CreateNewCompetitionHandler(w http.ResponseWriter, r *http.Request) {

	createDTO := new(viewmodel.CreateCompetition)

	if err := util.ParseRequestBodyData(r, createDTO); err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, err.Error())
		return
	}

	account, _ := util.GetCurrentUser(r, server.IAccountRepository)
	competition := createDTO.ToCompetitionDataModel(*account)

	err := businesslogic.CreateCompetition(competition, server.ICompetitionRepository, server.IOrganizerProvisionRepository, server.IOrganizerProvisionHistoryRepository)
	if err != nil {
		log.Printf("cannot create competition %v", err)
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot create competition", nil)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "success", nil)
}

// GET /api/organizer/competition
func (server OrganizerCompetitionServer) GetOrganizerCompetitionsHandler(w http.ResponseWriter, r *http.Request) {
	searchDTO := new(SearchOrganizerCompetitionViewModel)
	if parseErr := util.ParseRequestData(r, searchDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
	} else {
		account, _ := util.GetCurrentUser(r, server.IAccountRepository)
		if account.ID == 0 ||
			(account.AccountTypeID != businesslogic.ACCOUNT_TYPE_ORGANIZER &&
				account.AccountTypeID != businesslogic.ACCOUNT_TYPE_ADMINISTRATOR) {
			util.RespondJsonResult(w, http.StatusUnauthorized, "you are not authorized to look up this information", nil)
			return
		}
		criteria := businesslogic.SearchCompetitionCriteria{
			OrganizerID: account.ID,
		}
		if searchDTO.Future {
			criteria.StartDateTime = time.Now()
		}

		comps, err := server.SearchCompetition(&criteria)
		if err != nil {
			util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
			return
		} else {
			data := make([]viewmodel.Competition, 0)
			for _, each := range comps {
				data = append(data, viewmodel.CompetitionDataModelToViewModel(each, businesslogic.ACCOUNT_TYPE_ORGANIZER))
			}
			output, _ := json.Marshal(data)
			w.Write(output)
		}
	}
}

// PUT /api/organizer/competition
func (server OrganizerCompetitionServer) UpdateOrganizerCompetitionHandler(w http.ResponseWriter, r *http.Request) {
	account, _ := util.GetCurrentUser(r, server.IAccountRepository)
	updateDTO := new(businesslogic.OrganizerUpdateCompetition)

	if parseErr := util.ParseRequestBodyData(r, updateDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	competitions, _ := server.SearchCompetition(&businesslogic.SearchCompetitionCriteria{ID: updateDTO.CompetitionID})
	competitions[0].Street = updateDTO.Address
	competitions[0].UpdateStatus(updateDTO.Status) // TODO; error prone
	competitions[0].DateTimeUpdated = time.Now()
	competitions[0].UpdateUserID = account.ID

	if updateErr := server.UpdateCompetition(competitions[0]); updateErr != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, updateErr.Error())
		return
	}

	util.RespondJsonResult(w, http.StatusOK, "competition is updated", nil)
	return
}

// DELETE /api/organizer/competition
func deleteOrganzierCompetitionHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}
