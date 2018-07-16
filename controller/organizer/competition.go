// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package organizer

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/controller/util/authentication"
	"github.com/DancesportSoftware/das/viewmodel"
	"log"
	"net/http"
	"time"
)

type SearchOrganizerCompetitionViewModel struct {
	Future bool `schema:"future"`
}

type OrganizerCompetitionServer struct {
	authentication.IAuthenticationStrategy
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

	account, _ := server.GetCurrentUser(r, server.IAccountRepository)
	competition := createDTO.ToCompetitionDataModel(account)

	err := businesslogic.CreateCompetition(competition, server.ICompetitionRepository, server.IOrganizerProvisionRepository, server.IOrganizerProvisionHistoryRepository)
	if err != nil {
		log.Printf("cannot create competition %v", err)
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot create competition", nil)
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
		account, _ := server.GetCurrentUser(r, server.IAccountRepository)
		if account.ID == 0 ||
			(!account.HasRole(businesslogic.AccountTypeOrganizer) &&
				!account.HasRole(businesslogic.AccountTypeAdministrator)) {
			util.RespondJsonResult(w, http.StatusUnauthorized, "you are not authorized to look up this information", nil)
			return
		}
		criteria := businesslogic.SearchCompetitionCriteria{
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
	account, _ := server.GetCurrentUser(r, server.IAccountRepository)
	updateDTO := new(businesslogic.OrganizerUpdateCompetition)

	if parseErr := util.ParseRequestBodyData(r, updateDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	competitions, _ := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: updateDTO.CompetitionID})
	competitions[0].Street = updateDTO.Address
	competitions[0].UpdateStatus(updateDTO.Status) // TODO; error prone
	competitions[0].DateTimeUpdated = time.Now()
	competitions[0].UpdateUserID = account.ID

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
