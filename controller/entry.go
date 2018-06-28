// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	entries, err := server.SearchPartnershipEventEntry(*criteria)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	}

	output, _ := json.Marshal(entries)
	w.Write(output)

}
