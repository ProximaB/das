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

type SchoolServer struct {
	referencebll.ISchoolRepository
}

// GET /api/reference/school
func (server SchoolServer) SearchSchoolHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(referencebll.SearchSchoolCriteria)

	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.Http400InvalidRequestData, parseErr.Error())
		return
	}

	if schools, err := server.SearchSchool(*criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.Http500ErrorRetrievingData, err.Error())
		return
	} else {
		data := make([]viewmodel.School, 0)
		for _, each := range schools {
			data = append(data, viewmodel.SchoolDataModelToViewModel(each))
		}

		output, _ := json.Marshal(data)
		w.Write(output)
	}
}

// POST /api/reference/school
func (server SchoolServer) CreateSchoolHandler(w http.ResponseWriter, r *http.Request) {}

// PUT /api/reference/school
func (server SchoolServer) UpdateSchoolHandler(w http.ResponseWriter, r *http.Request) {}

// DELETE /api/reference/school
func (server SchoolServer) DeleteSchoolHandler(w http.ResponseWriter, r *http.Request) {}
