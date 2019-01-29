// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package account

import (
	"github.com/DancesportSoftware/das/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"log"
	"net/http"
)

// AccountServer provides a virtual server that handles requests that are related to Account
type AccountServer struct {
	auth.IAuthenticationStrategy
	businesslogic.IAccountRepository
	businesslogic.IAccountRoleRepository
	businesslogic.IOrganizerProvisionRepository
	businesslogic.IOrganizerProvisionHistoryRepository
	businesslogic.IUserPreferenceRepository
}

// RegisterAccountHandler handle the request
// 	POST /api/v1.0/account/register
// A JWT from Firebase should be sent to this endpoint. The handler will then send the token to Firebase for validation.
// Once the token is validated, this handler will pull data from Firebase, including user's name, email, phone number,
// and create a local account profile within the database.
// Identity is completely managed Firebase and DAS only checks the token and store additional user-provided data.
func (server AccountServer) RegisterAccountHandler(w http.ResponseWriter, r *http.Request) {

	log.Printf("%v", r.PostForm)
	currentUser, err := server.IAuthenticationStrategy.GetCurrentUser(r)
	createAccountDTO := new(viewmodel.CreateAccountDTO)
	if err := util.ParseRequestBodyData(r, createAccountDTO); err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	if err != nil {
		util.RespondJsonResult(w, http.StatusUnauthorized, err.Error(), nil)
		return
	}
	model := createAccountDTO.ToAccountModel()

	currentUser.FirstName = model.FirstName
	currentUser.LastName = model.LastName
	currentUser.Phone = model.Phone
	currentUser.Email = model.Email

	if err = currentUser.MeetMinimalRequirement(); err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := server.IAuthenticationStrategy.CreateUser(&currentUser); err != nil {
		log.Printf("[error] creating user in DAS: %v", err)
		util.RespondJsonResult(w, http.StatusInternalServerError, "Error in creating user profile", nil)
		return
	}
	// provision user with default role here
	defaultRole := businesslogic.NewAccountRole(currentUser, businesslogic.AccountTypeAthlete)
	createRoleErr := server.IAccountRoleRepository.CreateAccountRole(&defaultRole)
	if createRoleErr != nil {
		log.Printf("[error] creating user default role: %v", createRoleErr)
		util.RespondJsonResult(w, http.StatusInternalServerError, "Error in creating user's role", nil)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "User is registered in DAS successfully", nil)
}

// AccountAuthenticationHandler handles the request:
// 	POST /api/v1.0/account/authenticate
func (server AccountServer) AccountAuthenticationHandler(w http.ResponseWriter, r *http.Request) {
	account, err := server.IAuthenticationStrategy.GetCurrentUser(r)
	if err != nil {
		log.Printf("[error] Firebase cannot authenticate user : %s\n", err.Error())
		util.RespondJsonResult(w, http.StatusUnauthorized, "error in authentication", nil)
		return
	}
	log.Printf("[info] user %v is authenticated", account.FullName())
	util.RespondJsonResult(w, http.StatusOK, "authorized", nil)
	return
}
