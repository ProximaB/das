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

package businesslogic

import (
	"errors"
	"time"
)

// Country specifies the data needed to serve as reference data
type Country struct {
	ID              int
	Name            string
	Abbreviation    string
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

// GetStates retrieves all the states that are associated with the caller Country from the repository
func (country Country) GetStates(repo IStateRepository) ([]State, error) {
	if repo != nil {
		return repo.SearchState(SearchStateCriteria{CountryID: country.ID})
	}
	return nil, errors.New("null IStateRepository")
}

// GetFederations retrieves all the federations that are associated with the caller Country from the repository
func (country Country) GetFederations(repo IFederationRepository) ([]Federation, error) {
	if repo != nil {
		return repo.SearchFederation(SearchFederationCriteria{CountryID: country.ID})
	}
	return nil, errors.New("null IFederationRepository")
}

// SearchCountryCriteria specifies the parameters that can be used to search certain countries in DAS
type SearchCountryCriteria struct {
	CountryID    int    `schema:"id"`
	Name         string `schema:"name"`
	Abbreviation string `schema:"abbreviation"`
}

// ICountryRepository specifies the functions that a repository needs to implement
type ICountryRepository interface {
	CreateCountry(country *Country) error
	SearchCountry(criteria SearchCountryCriteria) ([]Country, error)
	DeleteCountry(country Country) error
	UpdateCountry(country Country) error
}
