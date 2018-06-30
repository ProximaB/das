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
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type OrganizerProvisionServer struct {
	businesslogic.IAccountRepository
	businesslogic.IOrganizerProvisionRepository
}

// PUT /api/admin/organizer/organizer
func (server OrganizerProvisionServer) UpdateOrganizerProvisionHandler(w http.ResponseWriter, r *http.Request) {
	updateDTO := new(viewmodel.UpdateProvision)
	parseErr := util.ParseRequestBodyData(r, updateDTO)
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid request data", nil)
		return
	}

	organizer := businesslogic.GetAccountByUUID(updateDTO.OrganizerID, server.IAccountRepository)
	provisions, _ := server.SearchOrganizerProvision(businesslogic.SearchOrganizerProvisionCriteria{OrganizerID: organizer.ID})

	record := provisions[0]
	// TODO: finish implementing the data update

	err := server.UpdateOrganizerProvision(record)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "success", nil)
}
