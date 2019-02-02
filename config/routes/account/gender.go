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

const apiAccountGenderEndpoint = "/api/account/gender"

var genderServer = account.GenderServer{
	database.GenderRepository,
}

var GenderController = util.DasController{
	Name:         "GenderController",
	Description:  "Get all genders in DAS",
	Method:       http.MethodGet,
	Endpoint:     apiAccountGenderEndpoint,
	Handler:      genderServer.GetAccountGenderHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
