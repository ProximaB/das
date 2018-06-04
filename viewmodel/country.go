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

func CountriesToViewModel(countries []reference.Country) []Country {
	results := make([]Country, 0)
	for _, each := range countries {
		results = append(results, countryToViewModel(each))
	}
	return results
}

func countryToViewModel(country reference.Country) Country {
	return Country{
		ID:           country.ID,
		Name:         country.Name,
		Abbreviation: country.Abbreviation,
	}

}

type CreateCountry struct {
	Name         string `schema:"name"`
	Abbreviation string `schema:"abbreviation"`
	CreateUserID int    `schema:"CreateUserID"`
}

func (cc *CreateCountry) ToDataModel() reference.Country {
	return reference.Country{
		Name:            cc.Name,
		Abbreviation:    cc.Abbreviation,
		CreateUserID:    &cc.CreateUserID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    &cc.CreateUserID,
		DateTimeUpdated: time.Now(),
	}
}

type DeleteCountry struct {
	CountryID    int    `schema:"id"`
	Name         string `schema:"name"`
	UpdateUserID int    `schema:"UpdateUserID"`
}

type UpdateCountry struct {
	CountryID    int    `schema:"CountryID"`
	Name         string `schema:"Name"`
	Abbreviation string `schema:"Abbreviation"`
	UpdateUserID int    `schema:"UpdateUserID"`
}
