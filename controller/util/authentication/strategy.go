package authentication

import (
	"github.com/yubing24/das/businesslogic"
	"net/http"
)

type AuthenticationStrategy interface {
	GetCurrentUser(r *http.Request) (businesslogic.Account, error)
	SetAuthorizationResponse(w http.ResponseWriter)
}
