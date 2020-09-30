package viewmodel

import (
	"github.com/ProximaB/das/businesslogic"
	"time"
)

// Country is the view model of a Country object
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

// CreateCountry defines the form data for creating a country record
type CreateCountry struct {
	Name         string `json:"name" validate:"min=3"`
	Abbreviation string `json:"abbreviation" validate:"min=3,max=3"` // Abbreviation is the Olympic country code
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

type State struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
	CountryID    int    `json:"country"`
}

func StateDataModelToViewModel(dm businesslogic.State) State {
	return State{
		ID:           dm.ID,
		Name:         dm.Name,
		Abbreviation: dm.Abbreviation,
		CountryID:    dm.CountryID,
	}
}

type City struct {
	CityID int    `json:"id"`
	Name   string `json:"name"`
	State  int    `json:"state"`
}

type CreateCity struct {
	Name    string `schema:"name"`
	StateID int    `schema:"state"`
}

func (create CreateCity) ToCityDataModel() businesslogic.City {
	return businesslogic.City{
		Name:    create.Name,
		StateID: create.StateID,
	}
}

type UpdateCity struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	StateID int    `json:"state"`
}

type DeleteCity struct {
	ID int `json:"id"`
}

type Federation struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

type SearchDivisionViewModel struct {
	DivisionID   int    `schema:"id"`
	Name         string `schema:"name"`
	FederationID int    `schema:"federation"`
}

type DivisionViewModel struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Federation int    `json:"federation"`
}

type SearchAge struct {
	DivisionID int `schema:"division"`
	AgeID      int `schema:"id"`
}

type Age struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Division int    `json:"division"`
	Enforced bool   `json:"enforced"`
	Minimum  int    `json:"minimum"`
	Maximum  int    `json:"maximum"`
}

func AgeDataModelToViewModel(dm businesslogic.Age) Age {
	return Age{
		ID:       dm.ID,
		Name:     dm.Name,
		Division: dm.DivisionID,
		Enforced: dm.Enforced,
		Minimum:  dm.AgeMinimum,
		Maximum:  dm.AgeMaximum,
	}
}

type Proficiency struct {
	ProficiencyID int    `json:"id"`
	Name          string `json:"name"`
	Division      int    `json:"division"`
}

func ProficiencyDataModelToViewModel(dm businesslogic.Proficiency) Proficiency {
	return Proficiency{
		ProficiencyID: dm.ID,
		Name:          dm.Name,
		Division:      dm.DivisionID,
	}
}

type Style struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Dance struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	StyleID      int    `json:"style"`
	Abbreviation string `json:"abbreviation"`
}

type School struct {
	SchoolID int    `json:"id"`
	Name     string `json:"name"`
	CityID   int    `json:"city"`
}

func SchoolDataModelToViewModel(school businesslogic.School) School {
	return School{
		SchoolID: school.ID,
		Name:     school.Name,
		CityID:   school.CityID,
	}
}

type Studio struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	CityID  int    `json:"city"`
	Website string `json:"website"`
}

func StudioDataModelToViewModel(dm businesslogic.Studio) Studio {
	return Studio{
		ID:      dm.ID,
		Name:    dm.Name,
		Address: dm.Address,
		CityID:  dm.CityID,
		Website: dm.Website,
	}
}

type Gender struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GenderDataModelToViewModel(gender businesslogic.Gender) Gender {
	return Gender{
		ID:   gender.ID,
		Name: gender.Name,
	}
}
