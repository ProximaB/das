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

type StyleServer struct {
	reference.IStyleRepository
}

// GET /api/reference/style
func (server StyleServer) SearchStyleHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(reference.SearchStyleCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid request data", parseErr.Error())
		return
	}

	if styles, err := server.IStyleRepository.SearchStyle(*criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, "error in retrieving styles", err.Error())
		return
	} else {
		data := make([]viewmodel.Style, 0)
		for _, each := range styles {
			viewmodel := viewmodel.Style{
				ID:   each.ID,
				Name: each.Name,
			}
			data = append(data, viewmodel)
		}

		output, _ := json.Marshal(data)
		w.Write(output)
	}

}

func (server StyleServer) CreateStyleHandler(w http.ResponseWriter, r *http.Request) {}
func (server StyleServer) UpdateStyleHandler(w http.ResponseWriter, r *http.Request) {}
func (server StyleServer) DeleteStyleHandler(w http.ResponseWriter, r *http.Request) {}
