// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package blacklist

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type PartnershipRequestBlacklistReasonServer struct {
	businesslogic.IPartnershipRequestBlacklistReasonRepository
}

// GET /api/partnership/blacklist/reason
func (server PartnershipRequestBlacklistReasonServer) GetPartnershipBlacklistReasonHandler(w http.ResponseWriter, r *http.Request) {
	reasons, err := server.GetPartnershipRequestBlacklistReasons()
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot retrieve partnership request blacklist reason", nil)
		return
	}

	data := make([]viewmodel.PartnershipRequestStatus, 0)
	for _, each := range reasons {
		view := viewmodel.PartnershipRequestStatus{
			ID:   each.ID,
			Name: each.Name,
		}
		data = append(data, view)
	}
	output, _ := json.Marshal(data)
	w.Write(output)
}
