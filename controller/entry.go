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

package controller

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

type EntryServer struct {
	businesslogic.IEventRepository
	businesslogic.IPartnershipEventEntryRepository
}

// GET /api/entries
// Public view for competitive event entry
func (server EntryServer) getCompetitiveBallroomEventEntryHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(businesslogic.SearchPartnershipEventEntryCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.Http400InvalidRequestData, parseErr.Error())
		return
	}

	entries, err := server.SearchPartnershipEventEntry(*criteria)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.Http500ErrorRetrievingData, err.Error())
		return
	}

	output, _ := json.Marshal(entries)
	w.Write(output)

}
