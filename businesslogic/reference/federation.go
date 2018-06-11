package referencebll

import (
	"errors"
	"time"
)

type Federation struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	YearFounded     int
	CountryID       int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

type SearchFederationCriteria struct {
	ID        int    `schema:"id"`
	Name      string `schema:"name"`
	CountryID int    `schema:"country"`
}

type IFederationRepository interface {
	CreateFederation(federation *Federation) error
	SearchFederation(criteria SearchFederationCriteria) ([]Federation, error)
	UpdateFederation(federation Federation) error
	DeleteFederation(federation Federation) error
}

func (federation Federation) GetDivisions(repo IDivisionRepository) ([]Division, error) {
	if repo == nil {
		return nil, errors.New("null IDivisionRepository")
	}
	return repo.SearchDivision(SearchDivisionCriteria{FederationID: federation.ID})
}
