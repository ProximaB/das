// Dancesport Application System (DAS)
// Copyright (C) 2019 Yubing Hou
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

package admin

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/auth"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type AdminUserManagementServer struct {
	auth.IAuthenticationStrategy
	accountRepo businesslogic.IAccountRepository
}

func NewAdminUserManagementServer(auth auth.IAuthenticationStrategy, accountRepo businesslogic.IAccountRepository) AdminUserManagementServer {
	return AdminUserManagementServer{
		auth,
		accountRepo,
	}
}

// GET /api/v1/admin/user
func (server AdminUserManagementServer) SearchUserHandler(w http.ResponseWriter, r *http.Request) {
	currentUser, _ := server.GetCurrentUser(r)
	if !currentUser.HasRole(businesslogic.AccountTypeAdministrator) {
		util.RespondJsonResult(w, http.StatusBadRequest, "Not authorized to search user accounts", nil)
		return
	}
	searchCriteriaDTO := new(viewmodel.SearchAccountDTO)
	parseErr := util.ParseRequestData(r, searchCriteriaDTO)

	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "Invalid search criteria data", nil)
		return
	}

	criteria := businesslogic.SearchAccountCriteria{}
	searchCriteriaDTO.Populate(&criteria)

	results, err := server.accountRepo.SearchAccount(criteria)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "An internal ", nil)
		return
	}

	data := make([]viewmodel.AccountDTO, 0)
	for _, each := range results {
		dto := viewmodel.AccountDTO{}
		dto.Extract(each)
		data = append(data, dto)
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}
