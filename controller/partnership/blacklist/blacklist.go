// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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
	account, _ := server.GetCurrentUser(r, server.IAccountRepository)

	blacklist, err := account.GetBlacklistedAccounts(server.IAccountRepository, server.IPartnershipRequestBlacklistRepository)

	if err != nil {
		log.Errorf(ctx, "error in getting partnership blacklist for user: %v", err)
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
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
