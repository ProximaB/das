package auth

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"net/http"
)

type IAuthenticationStrategy interface {
	GetCurrentUser(r *http.Request) (businesslogic.Account, error)
	CreateUser(account businesslogic.Account) error
}
