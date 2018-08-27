// Dancesport Application System (DAS)
// Copyright (C) 2018 Yubing Hou
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

package account

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/controller/util/authentication"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

// UserPreferenceServer provides a virtual server that handles requests that are related to User Preference
type UserPreferenceServer struct {
	authentication.IAuthenticationStrategy
	businesslogic.IAccountRepository
	businesslogic.IUserPreferenceRepository
}

// GetUserPreferenceHandler handles the request
//	GET /api/v1.0/account/preference
func (server UserPreferenceServer) GetUserPreferenceHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, userErr := server.GetCurrentUser(r)
	if userErr != nil {
		util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
		return
	}

	preferences, dalErr := server.SearchPreference(businesslogic.SearchUserPreferenceCriteria{
		AccountID: currentUser.ID,
	})

	if dalErr != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "an internal error has occurred", nil)
		return
	}
	if preferences == nil || len(preferences) != 1 {
		util.RespondJsonResult(w, http.StatusInternalServerError, "cannot find user preference", nil)
		return
	}

	view := viewmodel.UserPreferenceDataModelToViewModel(preferences[0])
	output, _ := json.Marshal(view)
	w.Write(output)
}

// UpdateUserPreferenceHandler handles the request
//	PUT /api/v1.0/account/preference
func (server UserPreferenceServer) UpdateUserPreferenceHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("not implemented"))
}
