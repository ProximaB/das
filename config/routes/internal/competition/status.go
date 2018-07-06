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

package competition

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/competition"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiCompetitionStatusEndpoint = "/api/competition/status"

var competitionStatusServer = competition.StatusServer{
	database.CompetitionStatusRepository,
}

var GetCompetitionStatusController = util.DasController{
	Name:         "GetCompetitionStatusController",
	Description:  "Get all competition status",
	Method:       http.MethodGet,
	Endpoint:     apiCompetitionStatusEndpoint,
	Handler:      competitionStatusServer.GetStatusHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
