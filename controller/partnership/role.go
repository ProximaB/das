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

package partnership

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"log"
	"net/http"
)

type PartnershipRoleServer struct {
	businesslogic.IPartnershipRoleRepository
}

func (server PartnershipRoleServer) GetPartnershipRolesHandler(w http.ResponseWriter, r *http.Request) {
	roles, err := server.IPartnershipRoleRepository.GetAllPartnershipRoles()
	if err != nil {
		log.Println(err)
		util.RespondJsonResult(w, http.StatusInternalServerError, "an error occurred while reading the data", nil)
	}

	view := make([]viewmodel.PartnershipRole, 0)
	for _, each := range roles {
		view = append(view, viewmodel.PartnershipRoleDataModelToViewModel(each))
	}
	output, _ := json.Marshal(view)
	w.Write(output)
}
