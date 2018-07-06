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

package competition

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

// StatusServer serves the referencedal data for competition status.
type StatusServer struct {
	businesslogic.ICompetitionStatusRepository
}

// GetStatusHandler allows client to get all possibles status of a competition.
// GET /api/competition/status
func (server StatusServer) GetStatusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := server.GetCompetitionAllStatus()
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]viewmodel.CompetitionStatus, 0)
	for _, each := range status {
		data = append(data, viewmodel.CompetitionStatusDataModelToViewModel(each))
	}
	output, _ := json.Marshal(data)
	w.Write(output)

}
