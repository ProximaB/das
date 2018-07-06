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

package request

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type PartnershipRequestStatusServer struct {
	businesslogic.IPartnershipRequestStatusRepository
}

// GET /api/partnership/status
func (server PartnershipRequestStatusServer) GetPartnershipRequestStatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := server.GetPartnershipRequestStatus()
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot retrieve partnership request status list", nil)
		return
	}

	data := make([]viewmodel.PartnershipRequestStatus, 0)
	for _, each := range status {
		view := viewmodel.PartnershipRequestStatus{
			ID:   each.ID,
			Name: each.Description,
		}
		data = append(data, view)
	}
	output, _ := json.Marshal(data)
	w.Write(output)
}
