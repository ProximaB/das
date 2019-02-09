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
