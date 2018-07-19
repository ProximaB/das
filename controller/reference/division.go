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

type DivisionServer struct {
	reference.IDivisionRepository
}

func (server DivisionServer) SearchDivisionHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(reference.SearchDivisionCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid request data", parseErr.Error())
		return
	}

	if divisions, err := server.IDivisionRepository.SearchDivision(*criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	} else {
		data := make([]viewmodel.DivisionViewModel, 0)
		for _, each := range divisions {
			view := viewmodel.DivisionViewModel{
				ID:         each.ID,
				Name:       each.Name,
				Federation: each.FederationID,
			}
			data = append(data, view)
		}
		output, _ := json.Marshal(data)
		w.Write(output)
	}

}

func (server DivisionServer) CreateDivisionHandler(w http.ResponseWriter, r *http.Request) {}
func (server DivisionServer) UpdateDivisionHandler(w http.ResponseWriter, r *http.Request) {}
func (server DivisionServer) DeleteDivisionHandler(w http.ResponseWriter, r *http.Request) {}
