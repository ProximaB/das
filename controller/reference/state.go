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
	"fmt"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type StateServer struct {
	referencebll.IStateRepository
}

// GET /api/reference/state
func (server StateServer) SearchStateHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(referencebll.SearchStateCriteria)
	err := util.ParseRequestData(r, criteria)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: err.Error()})
		return
	}
	states, err := server.IStateRepository.SearchState(*criteria)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: err.Error()})
		return
	}
	output := make([]viewmodel.State, 0)
	for _, each := range states {
		output = append(output, viewmodel.StateDataModelToViewModel(each))
	}
	data, _ := json.Marshal(output)
	w.Write(data)
}

// POST /api/reference/state
func (server StateServer) CreateStateHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}

// PUT /api/reference/state
func (server StateServer) UpdateStateHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}

// DELETE /api/reference/state
func (server StateServer) DeleteStateHandler(w http.ResponseWriter, r *http.Request) {
	util.RespondJsonResult(w, http.StatusNotImplemented, "not implemented", nil)
}
