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

package partnership

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/config/database"
	"github.com/DancesportSoftware/das/controller/partnership/blacklist"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

const apiPartnershipBlacklistReasonEndpoint = "/api/partnership/blacklist/reason"

var partnershipRequestBlacklistReasonServer = blacklist.PartnershipRequestBlacklistReasonServer{
	database.PartnershipRequestBlacklistReasonRepository,
}

var GetPartnershipBlacklistReasonController = util.DasController{
	Name:         "GetPartnershipBlacklistReasonController",
	Description:  "Get all the partnership blacklist reasons from DAS",
	Method:       http.MethodGet,
	Endpoint:     apiPartnershipBlacklistReasonEndpoint,
	Handler:      partnershipRequestBlacklistReasonServer.GetPartnershipBlacklistReasonHandler,
	AllowedRoles: []int{businesslogic.AccountTypeNoAuth},
}
