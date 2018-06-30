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
	if account.ID == 0 || account.AccountTypeID != businesslogic.AccountTypeAthlete {
		util.RespondJsonResult(w, http.StatusUnauthorized, "not authorized", nil)
		return
	}

	partnerships, err := server.SearchPartnership(
		businesslogic.SearchPartnershipCriteria{LeadID: account.ID, FollowID: account.ID})
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.Http500ErrorRetrievingData, err.Error())
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
		util.RespondJsonResult(w, http.StatusBadRequest, util.Http400InvalidRequestData, parseErr.Error())
		return
	}
}
