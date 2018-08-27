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

// City contains data to represent a City object in DAS. City is associated with State, which is associated with Country
type City struct {
	ID              int
	Name            string
	StateID         int
	CreateUserID    *int // 2017-10-17 yubing24: use pointer so that if the value is nil, it will be ignored
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

// GetSchools retrieves all schools that are in the caller city and from the repository
func (city City) GetSchools(repo ISchoolRepository) ([]School, error) {
	if repo != nil {
		return repo.SearchSchool(SearchSchoolCriteria{CityID: city.ID})
	}
	return nil, errors.New("null ISchoolRepository")
}

// GetStudios retrieves all the studios that are in the caller city and from the repository
func (city City) GetStudios(repo IStudioRepository) ([]Studio, error) {
	if repo != nil {
		return repo.SearchStudio(SearchStudioCriteria{CityID: city.ID})
	}
	return nil, errors.New("null IStudioRepository")
}

// SearchCityCriteria provides the parameter for search City in ICityRepository. This criteria can be used as
// parameters in REST API or internally
type SearchCityCriteria struct {
	CityID  int    `schema:"id"`
	Name    string `schema:"name"`
	StateID int    `schema:"state"`
}

// ICityRepository specifies the interface that data access layer code should implement
type ICityRepository interface {
	CreateCity(city *City) error
	SearchCity(criteria SearchCityCriteria) ([]City, error)
	UpdateCity(city City) error
	DeleteCity(city City) error
}
