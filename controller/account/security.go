package account

import (
	"net/http"
)

type AccountSecurityServer struct {
}

// UpdatePasswordHandler handles the request
//	PUT /api/v1.0/account/password
func (server AccountSecurityServer) UpdatePasswordHandler(w http.ResponseWriter, r *http.Request) {

}
