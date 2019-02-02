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

type Country struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

func CountriesToViewModel(countries []businesslogic.Country) []Country {
	results := make([]Country, 0)
	for _, each := range countries {
		results = append(results, countryToViewModel(each))
	}
	return results
}

func countryToViewModel(country businesslogic.Country) Country {
	return Country{
		ID:           country.ID,
		Name:         country.Name,
		Abbreviation: country.Abbreviation,
	}

}

type CreateCountry struct {
	Name         string `schema:"name"`
	Abbreviation string `schema:"abbreviation"`
}

func (cc *CreateCountry) ToDataModel() businesslogic.Country {
	return businesslogic.Country{
		Name:            cc.Name,
		Abbreviation:    cc.Abbreviation,
		DateTimeCreated: time.Now(),
		DateTimeUpdated: time.Now(),
	}
}

type DeleteCountry struct {
	CountryID int    `schema:"id"`
	Name      string `schema:"name"`
}

type UpdateCountry struct {
	CountryID    int    `schema:"CountryID"`
	Name         string `schema:"Name"`
	Abbreviation string `schema:"Abbreviation"`
	UpdateUserID int    `schema:"UpdateUserID"`
}
