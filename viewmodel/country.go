// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package viewmodel

import (
	"github.com/DancesportSoftware/das/businesslogic/reference"
	"time"
)

type Country struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

func CountriesToViewModel(countries []referencebll.Country) []Country {
	results := make([]Country, 0)
	for _, each := range countries {
		results = append(results, countryToViewModel(each))
	}
	return results
}

func countryToViewModel(country referencebll.Country) Country {
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

func (cc *CreateCountry) ToDataModel() referencebll.Country {
	return referencebll.Country{
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
