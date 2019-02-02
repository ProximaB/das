// Dancesport Application System (DAS)
// Copyright (C) 2018 Yubing Hou
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
	"github.com/DancesportSoftware/das/config/routes/middleware"
	"github.com/DancesportSoftware/das/controller/account"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiUserPreferenceEndpointV1_0 = "/api/v1.0/account/preference"

var preferenceServer = account.UserPreferenceServer{
	middleware.AuthenticationStrategy,
	database.AccountRepository,
	database.UserPreferenceRepository,
}

var searchUserPreferenceHandler = util.DasController{
	Name:        "SearchUserPreferenceHandler",
	Description: "Search user preference in DAS",
	Method:      http.MethodGet,
	Endpoint:    apiUserPreferenceEndpointV1_0,
	Handler:     preferenceServer.GetUserPreferenceHandler,
	AllowedRoles: []int{
		businesslogic.AccountTypeAthlete,
		businesslogic.AccountTypeAdjudicator,
		businesslogic.AccountTypeScrutineer,
		businesslogic.AccountTypeOrganizer,
		businesslogic.AccountTypeDeckCaptain,
		businesslogic.AccountTypeEmcee,
	},
}

var updateUserPreferenceHandler = util.DasController{
	Name:        "UpdateUserPreferenceHandler",
	Description: "Update user preference in DAS",
	Method:      http.MethodPut,
	Endpoint:    apiUserPreferenceEndpointV1_0,
	Handler:     preferenceServer.UpdateUserPreferenceHandler,
	AllowedRoles: []int{
		businesslogic.AccountTypeAthlete,
		businesslogic.AccountTypeAdjudicator,
		businesslogic.AccountTypeScrutineer,
		businesslogic.AccountTypeOrganizer,
		businesslogic.AccountTypeDeckCaptain,
		businesslogic.AccountTypeEmcee,
	},
}

var UserPreferenceControllerGroup = util.DasControllerGroup{
	Controllers: []util.DasController{
		searchUserPreferenceHandler,
		updateUserPreferenceHandler,
	},
}
