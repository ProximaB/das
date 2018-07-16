package middleware

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/authentication"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

func getRequestUserRole(r *http.Request) ([]int, error) {
	account, err := authentication.AuthenticationStrategy.GetCurrentUser(r, database.AccountRepository)
	if err != nil {
		return nil, err
	}
	return account.GetRoles(), nil
}

func AuthorizeMultipleRoles(h http.HandlerFunc, roles []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allowNoAuth := false
		for _, each := range roles {
			if each == businesslogic.AccountTypeNoAuth {
				allowNoAuth = true
				break
			}
		}

		userRoles, authErr := getRequestUserRole(r)
		if authErr != nil && !allowNoAuth {
			util.RespondJsonResult(w, http.StatusUnauthorized, "invalid authorization token", nil)
			return
		}

		authorized := false
		for _, each := range roles {
			for _, availableRole := range userRoles {
				if each == availableRole {
					authorized = true
					break
				}
			}
		}

		if authErr != nil && !allowNoAuth {
			util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
			return
		} else if allowNoAuth {
			h.ServeHTTP(w, r)
		} else if authorized {
			h.ServeHTTP(w, r)
		} else {
			util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
			return
		}
	}
}
