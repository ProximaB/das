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

package viewmodel

import (
	"github.com/DancesportSoftware/das/businesslogic"
	"time"
)

func GetEventDanceIDs(event businesslogic.Event) []int {
	output := make([]int, 0)
	for _, each := range event.GetDances() {
		output = append(output, each)
	}
	return output
}

type CreateEventViewModel struct {
	CompetitionID   int   `json:"competition"`
	EventCategoryID int   `json:"category"`
	FederationID    int   `json:"federation"`
	DivisionID      int   `json:"division"`
	AgeID           int   `json:"age"`
	ProficiencyID   int   `json:"proficiency"`
	StyleID         int   `json:"style"`
	Dances          []int `json:"dances"`
}

func (dto CreateEventViewModel) ToDomainModel(user businesslogic.Account) *businesslogic.Event {
	event := businesslogic.NewEvent()
	event.CompetitionID = dto.CompetitionID
	event.CategoryID = businesslogic.EventCategoryCompetitiveBallroom
	event.StatusID = businesslogic.EVENT_STATUS_DRAFT
	event.FederationID = dto.FederationID
	event.DivisionID = dto.DivisionID
	event.AgeID = dto.AgeID
	event.ProficiencyID = dto.ProficiencyID
	event.StyleID = dto.StyleID

	dances := make([]int, 0)
	for _, each := range dto.Dances {
		dances = append(dances, each)
	}
	event.SetDances(dances)
	event.CreateUserID = user.ID
	event.DateTimeCreated = time.Now()
	event.UpdateUserID = user.ID
	event.DateTimeUpdated = time.Now()

	return event
}
