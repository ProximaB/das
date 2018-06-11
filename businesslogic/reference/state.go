package referencebll

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
