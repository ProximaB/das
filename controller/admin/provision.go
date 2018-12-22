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

package admin

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"log"
	"net/http"
)

type OrganizerProvisionServer struct {
	auth.IAuthenticationStrategy
	accountRepo businesslogic.IAccountRepository
	service     businesslogic.OrganizerProvisionService
}

func NewOrganizerProvisionServer(strategy auth.IAuthenticationStrategy, accountRepo businesslogic.IAccountRepository, service businesslogic.OrganizerProvisionService) OrganizerProvisionServer {
	return OrganizerProvisionServer{
		strategy,
		accountRepo,
		service,
	}
}

// UpdateOrganizerProvisionHandler handles the request
//	PUT /api/admin/organizer/organizer
// which allocates organizer competition slot for hosting competitions
func (server OrganizerProvisionServer) UpdateOrganizerProvisionHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, _ := server.GetCurrentUser(r)

	updateDTO := new(viewmodel.UpdateProvision)
	parseErr := util.ParseRequestBodyData(r, updateDTO)
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid request data", nil)
		return
	}

	organizer := businesslogic.GetAccountByUUID(updateDTO.OrganizerID, server.accountRepo)
	update := businesslogic.UpdateOrganizerProvision{
		OrganizerID:   organizer.ID,
		Amount:        updateDTO.AmountAllocated,
		Note:          updateDTO.Note,
		CurrentUserID: currentUser.ID,
	}

	err := server.service.UpdateOrganizerCompetitionProvision(update)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "success", nil)
}

// GetOrganizerProvisionSummaryHandler get the summary of organizer competition provision
func (server OrganizerProvisionServer) GetOrganizerProvisionSummaryHandler(w http.ResponseWriter, r *http.Request) {
	provisions, err := server.service.SearchOrganizerProvision(businesslogic.SearchOrganizerProvisionCriteria{})
	if err != nil {
		log.Println(err)
		util.RespondJsonResult(w, http.StatusInternalServerError, "an error occurred while trying to retrieve organizer provision information", nil)
		return
	}

	records := make([]viewmodel.OrganizerProvisionSummary, 0)
	for _, each := range provisions {
		record := viewmodel.OrganizerProvisionSummary{}
		record.Summarize(each)
		records = append(records, record)
	}

	output, _ := json.Marshal(records)
	w.Write(output)
}
