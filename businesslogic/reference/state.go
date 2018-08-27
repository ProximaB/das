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
	"errors"
	"time"
)

type State struct {
	ID              int
	Name            string
	Abbreviation    string
	CountryID       int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

type SearchStateCriteria struct {
	StateID   int    `schema:"id"`
	Name      string `schema:"name"`
	CountryID int    `schema:"country"`
}

type IStateRepository interface {
	CreateState(state *State) error
	SearchState(criteria SearchStateCriteria) ([]State, error)
	UpdateState(state State) error
	DeleteState(state State) error
}

func (state State) GetCities(repo ICityRepository) ([]City, error) {
	if repo == nil {
		return nil, errors.New("null ICityRepository")
	}
	return repo.SearchCity(SearchCityCriteria{StateID: state.ID})
}
