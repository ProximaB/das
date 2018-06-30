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
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

// AccountTypeServer is a micro-server that serves requests that ask for available
// account types in DAS
type AccountTypeServer struct {
	businesslogic.IAccountTypeRepository
}

// GetAccountTypeHandler handles the request
//	GET /api/account/type
// No parameter is required for this request.
//
// Sample returned result:
//	[
// 		{"id":1,"name":"Athlete"},
// 		{"id":2,"name":"Adjudicator"},
// 		{"id":3,"name":"Scrutineer"},
// 		{"id":4,"name":"Organizer"},
// 		{"id":5,"name":"Deck Captain"},
// 		{"id":6,"name":"Emcee"}
// 	]
func (server AccountTypeServer) GetAccountTypeHandler(w http.ResponseWriter, r *http.Request) {
	data := make([]viewmodel.AccountType, 0)
	types, _ := server.GetAccountTypes()
	for _, each := range types {
		// System administrator should not be available for public registration
		if each.ID != businesslogic.AccountTypeAdministrator {
			data = append(data, viewmodel.AccountTypeDataModelToViewModel(each))
		}
	}
	output, _ := json.Marshal(data)
	w.Write(output)
}
