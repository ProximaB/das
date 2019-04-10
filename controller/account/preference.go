package account

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

// UserPreferenceServer provides a virtual server that handles requests that are related to User Preference
type UserPreferenceServer struct {
	auth.IAuthenticationStrategy
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
