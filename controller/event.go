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

package controller

import (
	"encoding/json"
	"github.com/DancesportSoftware/das/businesslogic"
	"github.com/DancesportSoftware/das/controller/util"
	"net/http"
)

type EventViewModel struct {
	ID            int    `json:"id"`
	CompetitionID int    `json:"competition"`
	Category      int    `json:"category"`
	Description   string `json:"description"`
	StatusID      int    `json:"status"`
}

type EventServer struct {
	businesslogic.IEventRepository
}

// GET /api/event
func (server EventServer) GetEventHandler(w http.ResponseWriter, r *http.Request) {
	criteria := new(businesslogic.SearchEventCriteria)
	if parseErr := util.ParseRequestData(r, criteria); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
		return
	} else {

		events, err := server.SearchEvent(*criteria)
		if err != nil {
			util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		data := make([]EventViewModel, 0)
		for _, each := range events {
			data = append(data, EventViewModel{
				ID:            each.ID,
				CompetitionID: each.CompetitionID,
				Category:      each.CategoryID,
				Description:   each.Description,
				StatusID:      each.StatusID,
			})
		}
		output, _ := json.Marshal(data)
		w.Write(output)
	}
}

type SearchCompetitiveBallroomEventViewModel struct {
	ID               int  `schema:"id"`
	CompetitionID    int  `schema:"competition"`
	FederationID     int  `schema:"federation"`
	DivisionID       int  `schema:"division"`
	AgeID            int  `schema:"age"`
	ProficiencyID    int  `schema:"proficiency"`
	StyleID          int  `schema:"style"`
	OpenRegistration bool `schema:"open"`
}
type CompetitiveBallroomEventViewModel struct {
	ID                         int   `json:"eventID"` // competitive ballroom event id
	CompetitiveBallroomEventID int   `json:"cbeID"`   // event id
	FederationID               int   `json:"federation"`
	DivisionID                 int   `json:"division"`
	AgeID                      int   `json:"age"`
	ProficiencyID              int   `json:"proficiency"`
	StyleID                    int   `json:"style"`
	Dances                     []int `json:"dances"`
}

// GET /api/event
// This does not require identity as site visitors may want to see how it works, too.
/*
func getCompetitiveBallroomEventHandler(w http.ResponseWriter, r *http.Request) {
	searchDTO := new(SearchCompetitiveBallroomEventViewModel)

	if parseErr := util.ParseRequestData(r, searchDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, util.HTTP400InvalidRequestData, parseErr.Error())
	} else {
		criteria := businesslogic.SearchEventCriteria{
			EventID:       searchDTO.ID,
			ID: searchDTO.ID,
			FederationID:  searchDTO.FederationID,
			DivisionID:    searchDTO.DivisionID,
			AgeID:         searchDTO.AgeID,
			ProficiencyID: searchDTO.ProficiencyID,
			ID:       searchDTO.ID,
		}
		if searchDTO.OpenRegistration {
			criteria.StatusID = businesslogic.EVENT_STATUS_OPEN
		}

		events, _ := eventRepo.SearchEvent(&criteria)
		data := make([]CompetitiveBallroomEventViewModel, 0)
		for _, each := range events {
			data = append(data, CompetitiveBallroomEventViewModel{
				ID: each.ID,
				CompetitiveBallroomEventID: each.ID,
				FederationID:               each.FederationID,
				DivisionID:                 each.DivisionID,
				AgeID:                      each.AgeID,
				ProficiencyID:              each.ProficiencyID,
				ID:                    each.ID,
				Dances:                     viewmodel.GetEventDanceIDs(each),
			})
		}
		output, _ := json.Marshal(data)
		w.Write(output)
	}
}

// POST /api/organizer/event
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	account, _ := GetCurrentUser(r, accountRepository)
	createDTO := new(viewmodel.CreateEventViewModel)

	if parseErr := util.ParseRequestBodyData(r, createDTO); parseErr != nil {
		util.RespondJsonResult(w, http.StatusBadRequest, "invalid request data", nil)
		return
	}

	event := createDTO.ToDomainModel(*account, danceRepository)
	err := businesslogic.CreateEvent(event, competitionRepository, eventRepo, eventDanceRepository)
	if err != nil {
		util.RespondJsonResult(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	util.RespondJsonResult(w, http.StatusOK, "success", nil)
}
*/
