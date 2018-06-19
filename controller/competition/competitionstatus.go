// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	}

	data := make([]viewmodel.CompetitionStatus, 0)
	for _, each := range status {
		data = append(data, viewmodel.CompetitionStatusDataModelToViewModel(each))
	}
	output, _ := json.Marshal(data)
	w.Write(output)

}
