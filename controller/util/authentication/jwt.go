package authentication

import (
	"github.com/yubing24/das/businesslogic"
	"errors"
	"net/http"
)

type jwtAuthenticationStrategy struct {
}

func NewJwtAuthenticationStrategy() jwtAuthenticationStrategy {

}

func (strategy jwtAuthenticationStrategy) GetCurrentUser(r *http.Request) (businesslogic.Account, error) {
	return businesslogic.Account{}, errors.New("not implemented")
}

func (strategy jwtAuthenticationStrategy) SetAuthorizationResponse(w http.ResponseWriter) {
}
