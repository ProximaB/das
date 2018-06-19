// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package partnership

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/controller/util/authentication"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type PartnershipServer struct {
	authentication.IAuthenticationStrategy
	businesslogic.IAccountRepository
	businesslogic.IPartnershipRepository
}

// GET /api/partnership
func (server PartnershipServer) SearchPartnershipHandler(w http.ResponseWriter, r *http.Request) {
	account, _ := server.GetCurrentUser(r, server.IAccountRepository)
	if account.ID == 0 || account.AccountTypeID != businesslogic.ACCOUNT_TYPE_ATHLETE {
		util.RespondJsonResult(w, http.StatusUnauthorized, "not authorized", nil)
		return
	}

	partnerships, err := server.SearchPartnership(
		businesslogic.SearchPartnershipCriteria{LeadID: account.ID, FollowID: account.ID})
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	}

	data := make([]viewmodel.Partnership, 0)
	for _, each := range partnerships {
		data = append(data, viewmodel.PartnershipDataModelToViewModel(each))
	}
	output, _ := json.Marshal(data)
	w.Write(output)

}

type updatePartnership struct {
	PartnershipID int  `json:"partnership"`
	Favorite      bool `json:"favorite"`
}

// PUT /api/partnership
func (server PartnershipServer) UpdatePartnershipHandler(w http.ResponseWriter, r *http.Request) {
	account, _ := server.GetCurrentUser(r, server.IAccountRepository)
	if account.ID == 0 {
		util.RespondJsonResult(w, http.StatusUnauthorized, "not authorized", nil)
		return
	}

	updateDTO := new(updatePartnership)
	if parseErr := util.ParseRequestBodyData(r, updateDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}
}
