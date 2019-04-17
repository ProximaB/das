package auth

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"net/http"
)

// IAuthenticationStrategy specifies the interface that an authentication strategy should implement. The implementation
// must be able to identify user through the HTTP request (via header, token, cookie, etc.), and must be able to create
// user account in the necessary user repository (database, LDAP, identity provider, etc)
type IAuthenticationStrategy interface {
	GetCurrentUser(r *http.Request) (businesslogic.Account, error)
	CreateUser(account *businesslogic.Account) error
}
