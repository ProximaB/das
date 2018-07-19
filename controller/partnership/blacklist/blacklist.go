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

package blacklist

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/controller/util/authentication"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"net/http"
	"time"
)

type PartnershipBlacklistViewModel struct {
	Username string    `json:"user"`
	Since    time.Time `json:"since"`
}

type PartnershipRequestBlacklistServer struct {
	authentication.IAuthenticationStrategy
	businesslogic.IAccountRepository
	businesslogic.IPartnershipRequestBlacklistRepository
}

// GET /api/partnership/blacklist
func (server PartnershipRequestBlacklistServer) GetBlacklistedAccountHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	account, _ := server.GetCurrentUser(r)

	blacklist, err := account.GetBlacklistedAccounts(server.IAccountRepository, server.IPartnershipRequestBlacklistRepository)

	if err != nil {
		log.Errorf(ctx, "error in getting partnership blacklist for user: %v", err)
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]PartnershipBlacklistViewModel, 0)
	for _, each := range blacklist {
		entry := PartnershipBlacklistViewModel{
			Username: each.FirstName + " " + each.LastName,
		}
		data = append(data, entry)
	}
	output, _ := json.Marshal(data)
	w.Write(output)
}

// POST /api/partnership/blacklist/report
func (server PartnershipRequestBlacklistServer) CreatePartnershipRequestBlacklistReportHandler(w http.ResponseWriter, r *http.Request) {

}
