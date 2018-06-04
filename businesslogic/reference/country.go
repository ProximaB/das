package reference

import (
	"errors"
	"time"
)

type Country struct {
	ID              int
	Name            string
	Abbreviation    string
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

func (country Country) GetStates(repo IStateRepository) ([]State, error) {
	if repo != nil {
		return repo.SearchState(SearchStateCriteria{CountryID: country.ID})

	}
	return nil, errors.New("null IStateRepository")
}

func (country Country) GetFederations(repo IFederationRepository) ([]Federation, error) {
	if repo != nil {
		return repo.SearchFederation(SearchFederationCriteria{CountryID: country.ID})
	}
	return nil, errors.New("null IFederationRepository")
}

type SearchCountryCriteria struct {
	CountryID    int    `schema:"id"`
	Name         string `schema:"name"`
	Abbreviation string `schema:"abbreviation"`
}

type ICountryRepository interface {
	CreateCountry(country *Country) error
	SearchCountry(criteria SearchCountryCriteria) ([]Country, error)
	DeleteCountry(country Country) error
	UpdateCountry(country Country) error
}
