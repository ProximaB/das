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

package reference

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type ProficiencyServer struct {
	reference.IProficiencyRepository
}

// GET /api/reference/proficiency
func (server ProficiencyServer) SearchProficiencyHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(reference.SearchProficiencyCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid request data", parseErr.Error())
		return
	}

	if proficiencies, err := server.IProficiencyRepository.SearchProficiency(*criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	} else {
		dtos := make([]viewmodel.Proficiency, 0)
		for _, each := range proficiencies {
			dtos = append(dtos, viewmodel.ProficiencyDataModelToViewModel(each))
		}
		output, _ := json.Marshal(dtos)
		w.Write(output)
	}
}

// POST /api/reference/proficiency
func (server ProficiencyServer) CreateProficiencyHandler(w http.ResponseWriter, r *http.Request) {}

// DELETE /api/reference/proficiency
func (server ProficiencyServer) DeleteProficiencyHandler(w http.ResponseWriter, r *http.Request) {}

// PUT /api/reference/proficiency
func (server ProficiencyServer) UpdateProficiencyHandler(w http.ResponseWriter, r *http.Request) {}
