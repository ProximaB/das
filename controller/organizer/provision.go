// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package organizer

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/controller/util/authentication"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
	"time"
)

type OrganizerProvisionServer struct {
	authentication.IAuthenticationStrategy
	businesslogic.IAccountRepository
	businesslogic.IOrganizerProvisionRepository
}

// GET /api/organizer/organizer/summary
func (server OrganizerProvisionServer) GetOrganizerProvisionSummaryHandler(w http.ResponseWriter, r *http.Request) {

	account, _ := server.GetCurrentUser(r, server.IAccountRepository)
	if account.AccountTypeID != businesslogic.AccountTypeOrganizer || account.ID == 0 {
		util.RespondJsonResult(w, http.StatusUnauthorized, "Access denied", nil)
		return
	}

	summaries, _ := server.SearchOrganizerProvision(businesslogic.SearchOrganizerProvisionCriteria{OrganizerID: account.ID})
	view := viewmodel.OrganizerProvisionSummary{
		OrganizerID: summaries[0].OrganizerID,
		Available:   summaries[0].Available,
		Hosted:      summaries[0].Hosted,
	}

	output, _ := json.Marshal(view)
	w.Write(output)
}

type OrganizerProvisionHistoryEntryViewModel struct {
	OrganizerID       int       `json:"organizer"`
	Allocated         int       `json:"allocated"`
	DateTimeAllocated time.Time `json:"date"`
}

type OrganizerProvisionHistoryServer struct {
	authentication.IAuthenticationStrategy
	businesslogic.IAccountRepository
	businesslogic.IOrganizerProvisionHistoryRepository
}

// GET /api/organizer/organizer/history
func (server OrganizerProvisionHistoryServer) GetOrganizerProvisionHistoryHandler(w http.ResponseWriter, r *http.Request) {

	account, _ := server.GetCurrentUser(r, server.IAccountRepository)
	if account.AccountTypeID != businesslogic.AccountTypeOrganizer && account.AccountTypeID != businesslogic.AccountTypeAdministrator {
		util.RespondJsonResult(w, http.StatusUnauthorized, "Access denied", nil)
		return
	}

	history, err := server.SearchOrganizerProvisionHistory(businesslogic.SearchOrganizerProvisionHistoryCriteria{OrganizerID: account.ID})
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	data := make([]OrganizerProvisionHistoryEntryViewModel, 0)
	for _, each := range history {
		entry := OrganizerProvisionHistoryEntryViewModel{
			OrganizerID:       each.OrganizerID,
			Allocated:         each.Amount,
			DateTimeAllocated: each.DateTimeCreated,
		}
		data = append(data, entry)
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}
