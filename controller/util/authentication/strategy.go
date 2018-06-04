package authentication

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"net/http"
)

type IAuthenticationStrategy interface {
	GetCurrentUser(r *http.Request, repository businesslogic.IAccountRepository) (businesslogic.Account, error)
	SetAuthorizationResponse(w http.ResponseWriter)
}
