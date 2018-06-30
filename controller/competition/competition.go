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

package competition

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"github.com/DancesportSoftware/das/viewmodel"
	"net/http"
	"strconv"
)

type PublicCompetitionServer struct {
	businesslogic.ICompetitionRepository
	businesslogic.IEventRepository
	businesslogic.IEventMetaRepository
}

// GET /api/competitions
// Search competition(s). This controller is invokable without authentication
func (server PublicCompetitionServer) SearchCompetitionHandler(w http.ResponseWriter, r *http.Request) {
	searchDTO := new(businesslogic.SearchCompetitionCriteria)
	if parseErr := util.ParseRequestData(r, searchDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	} else {
		if competitions, err := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{
			ID:       searchDTO.ID,
			Name:     searchDTO.Name,
			StatusID: searchDTO.StatusID,
		}); err != nil {
			util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		} else {
			data := make([]viewmodel.Competition, 0)
			for _, each := range competitions {
				data = append(data, viewmodel.CompetitionDataModelToViewModel(each, businesslogic.AccountTypeNoAuth))
			}
			output, _ := json.Marshal(data)
			w.Write(output)
		}
	}
}

// GET /api/competition/federation
func (server PublicCompetitionServer) GetUniqueEventFederationHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	searchResults, _ := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: compID})
	competition := searchResults[0]

	federations, err := competition.GetEventUniqueFederations(server.IEventMetaRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]viewmodel.Federation, 0)
	for _, each := range federations {
		data = append(data, viewmodel.Federation{
			ID:           each.ID,
			Name:         each.Name,
			Abbreviation: each.Abbreviation,
		})
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// GET /api/competition/division
func (server PublicCompetitionServer) GetEventUniqueDivisionsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	searchResults, _ := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: compID})
	competition := searchResults[0]

	divisions, err := competition.GetEventUniqueDivisions(server.IEventMetaRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]viewmodel.DivisionViewModel, 0)
	for _, each := range divisions {
		data = append(data, viewmodel.DivisionViewModel{
			ID:         each.ID,
			Name:       each.Name,
			Federation: each.FederationID,
		})
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// GET /api/competition/age
func (server PublicCompetitionServer) GetEventUniqueAgesHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	searchResults, _ := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: compID})
	competition := searchResults[0]
	ages, err := competition.GetEventUniqueAges(server.IEventMetaRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]viewmodel.Age, 0)
	for _, each := range ages {
		data = append(data, viewmodel.Age{
			ID:       each.ID,
			Name:     each.Name,
			Division: each.DivisionID,
			Maximum:  each.AgeMaximum,
			Minimum:  each.AgeMinimum,
		})
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// GET /api/competition/proficiency
func (server PublicCompetitionServer) GetEventUniqueProficienciesHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}
	searchResults, _ := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: compID})
	competition := searchResults[0]
	proficiencies, err := competition.GetEventUniqueProficiencies(server.IEventMetaRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]viewmodel.Proficiency, 0)
	for _, each := range proficiencies {
		data = append(data, viewmodel.ProficiencyDataModelToViewModel(each))
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}

// GET /api/competition/style
func (server PublicCompetitionServer) GetEventUniqueStylesHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	compID, parseErr := strconv.Atoi(r.Form.Get("competition"))
	if parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	}

	searchResults, _ := server.SearchCompetition(businesslogic.SearchCompetitionCriteria{ID: compID})
	competition := searchResults[0]
	styles, err := competition.GetEventUniqueStyles(server.IEventMetaRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, util.HTTP500ErrorRetrievingData, err.Error())
		return
	}

	data := make([]viewmodel.Style, 0)
	for _, each := range styles {
		data = append(data, viewmodel.Style{
			ID:   each.ID,
			Name: each.Name,
		})
	}

	output, _ := json.Marshal(data)
	w.Write(output)
}
