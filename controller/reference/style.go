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

type StyleServer struct {
	referencebll.IStyleRepository
}

// GET /api/reference/style
func (server StyleServer) SearchStyleHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(referencebll.SearchStyleCriteria)
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
