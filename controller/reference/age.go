// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package reference

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
)

type AgeServer struct {
	referencebll.IAgeRepository
}

// SearchAgeHandler handles request
//	GET /api/reference/age
//
// Accepted parameters:
//	{
//		"id": 1,
//		"division": 3
//	}
// Sample results returned:
//	[
//		{ "id": 3, "name": "Adult", "division": 4, "enforced": true, "minimum": 19, "maximum": 99 },
//		{ "id": 4, "name": "Senior I", "division": 4, "enforced": true, "minimum": 36, "maximum": 45 },
//	]
func (server AgeServer) SearchAgeHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(referencebll.SearchAgeCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP_400_INVALID_REQUEST_DATA, parseErr.Error())
		return
	}

	if ages, err := server.SearchAge(*criteria); err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP_500_ERROR_RETRIEVING_DATA, err.Error())
		return
	} else {
		data := make([]viewmodel.Age, 0)
		for _, each := range ages {
			data = append(data, viewmodel.AgeDataModelToViewModel(each))
		}

		output, _ := json.Marshal(data)
		w.Write(output)
	}
}

func (server AgeServer) CreateAgeHandler(w http.ResponseWriter, r *http.Request) {

}

func (server AgeServer) UpdateAgeHandler(w http.ResponseWriter, r *http.Request) {

}

func (server AgeServer) DeleteAgeHandler(w http.ResponseWriter, r *http.Request) {

}
