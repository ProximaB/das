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
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/account"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiAccountTypeEndpoint = "/api/account/type"

var accountTypeServer = account.AccountTypeServer{
	database.AccountTypeRepository,
}

var AccountTypeController = util.DasController{
	Name:         "AccountTypeController",
	Description:  "Get all account types in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiAccountTypeEndpoint,
	Handler:      accountTypeServer.GetAccountTypeHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
