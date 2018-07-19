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
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

// GenderServer serves requests that ask for all possible gender options in DAS
type GenderServer struct {
	reference.IGenderRepository
}

// GetAccountGenderHandler handles request
//		GET /api/account/gender
// No parameter is required for this request.
//
// Sample returned result:
//	[
// 		{"id":1,"name":"Female"},
// 		{"id":2,"name":"Male"}
// 	]
func (handler GenderServer) GetAccountGenderHandler(w http.ResponseWriter, r *http.Request) {
	data := make([]viewmodel.Gender, 0)
	genders, _ := handler.IGenderRepository.GetAllGenders()
	for _, each := range genders {
		data = append(data, viewmodel.GenderDataModelToViewModel(each))
	}
	output, _ := json.Marshal(data)
	w.Write(output)
}
