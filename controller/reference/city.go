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
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type CityServer struct {
	businesslogic.ICityRepository
}

// POST /api/reference/city
func (server CityServer) CreateCityHandler(w http.ResponseWriter, r *http.Request) {
	dto := new(viewmodel.CreateCity)
	if err := util.ParseRequestBodyData(r, dto); err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	city := dto.ToCityDataModel()

	if err := server.ICityRepository.CreateCity(&city); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	util.RespondJsonResult(w, http.StatusOK, "success", nil)
}

// DELETE /api/reference/city
func (server CityServer) DeleteCityHandler(w http.ResponseWriter, r *http.Request) {
	deleteDTO := new(viewmodel.DeleteCity)
	err := util.ParseRequestBodyData(r, deleteDTO)
	if err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if cities, searchErr := server.ICityRepository.SearchCity(businesslogic.SearchCityCriteria{CityID: deleteDTO.ID}); searchErr != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, searchErr.Error(), nil)
		return
	} else if len(cities) != 1 {
		util.RespondJsonResult(w, http.StatusNotFound, "cannot find specified city", nil)
		return
	} else {
		if deleteErr := server.ICityRepository.DeleteCity(cities[0]); deleteErr != nil {
			util.RespondJsonResult(w, http.StatusInternalServerError, "cannot delete specified city", nil)
			return
		}
		util.RespondJsonResult(w, http.StatusOK, "success", nil)
		return
	}
}

// PUT /api/reference/city
func (server CityServer) UpdateCityHandler(w http.ResponseWriter, r *http.Request) {
	updateDTO := new(viewmodel.UpdateCity)
	err := util.ParseRequestBodyData(r, updateDTO)
	if err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	util.RespondJsonResult(w, http.StatusNotImplemented, "", nil)
}

// GET /api/reference/city
func (server CityServer) SearchCityHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(businesslogic.SearchCityCriteria)
	err := util.ParseRequestData(r, criteria)
	if err != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, err.Error())
		return
	}
	cities, err := server.ICityRepository.SearchCity(*criteria)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	dtos := make([]viewmodel.City, 0)
	for _, each := range cities {
		dtos = append(dtos, viewmodel.City{
			CityID: each.ID,
			Name:   each.Name,
			State:  each.StateID,
		})
	}
	output, _ := json.Marshal(dtos)
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
