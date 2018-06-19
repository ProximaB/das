// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package authentication

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"net/http"
)

type IAuthenticationStrategy interface {
	GetCurrentUser(r *http.Request, repository businesslogic.IAccountRepository) (businesslogic.Account, error)
	SetAuthorizationResponse(w http.ResponseWriter)
}
