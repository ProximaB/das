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

type FederationServer struct {
	referencebll.IFederationRepository
}

// GET /api/reference/federation
func (server FederationServer) SearchFederationHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(referencebll.SearchFederationCriteria)
	if err := util.ParseRequestData(r, criteria); err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, err.Error())
		return
	}

	federations, err := server.IFederationRepository.SearchFederation(*criteria)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, viewmodel.RESTAPIResult{Message: err.Error()})
		return
	}

	dtos := make([]viewmodel.Federation, 0)
	for _, each := range federations {
		dtos = append(dtos, viewmodel.Federation{
			ID:           each.ID,
			Name:         each.Name,
			Abbreviation: each.Abbreviation,
		})
	}
	output, _ := json.Marshal(dtos)
	w.Write(output)
}

// POST /api/reference/federation
func (server FederationServer) CreateFederationHandler(w http.ResponseWriter, r *http.Request) {}

// DELETE /api/reference/federation
func (server FederationServer) DeleteFederationHandler(w http.ResponseWriter, r *http.Request) {}

// PUT /api/reference/federation
func (server FederationServer) UpdateFederationHandler(w http.ResponseWriter, r *http.Request) {}
