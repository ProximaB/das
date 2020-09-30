package middleware

import (
	"github.com/ProximaB/das/businesslogic"
	"github.com/ProximaB/das/controller/util"
	"log"
	"net/http"
)

func getRequestUserRole(r *http.Request) ([]int, error) {
	account, err := AuthenticationStrategy.GetCurrentUser(r)
	if err != nil {
		return nil, err
	}
	return account.GetRoles(), nil
}

func allowUnauthorizedRequest(roles []int) bool {
	for _, each := range roles {
		if each == businesslogic.AccountTypeNoAuth {
			return true
		}
	}
	return false
}

// AuthorizeMultipleRoles checks if the user's token contains the role that the handler requires. If not, the handler
// function will not be executed
func AuthorizeMultipleRoles(h http.HandlerFunc, roles []int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		allowNoAuth := allowUnauthorizedRequest(roles)

		// all no auth will be passed from interception
		if allowNoAuth {
			h.ServeHTTP(w, r)
			return
		}

		userRoles, authErr := getRequestUserRole(r)
		if authErr != nil && !allowNoAuth {
			log.Printf("[error] authentication error occurred when the %s requires a role: %v", r.RequestURI, roles)
			util.RespondJsonResult(w, http.StatusUnauthorized, authErr.Error(), nil)
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

		// authorization token is invalid, and request does not allow unauthorized request
		if authErr != nil && !allowNoAuth {
			util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
			return
		}
		// unauthorized request is allowed
		if allowNoAuth {
			h.ServeHTTP(w, r)
		} else if authorized && !allowNoAuth {
			h.ServeHTTP(w, r)
		} else {
			util.RespondJsonResult(w, http.StatusUnauthorized, "unauthorized", nil)
			return
		}
	}
}
