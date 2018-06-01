package reference

import (
	"errors"
	"time"
)

type City struct {
	CityID          int
	Name            string
	StateID         int
	CreateUserID    *int // 2017-10-17 yubing24: use pointer so that if the value is nil, it will be ignored
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

func (city City) GetSchools(repo ISchoolRepository) ([]School, error) {
	if repo != nil {
		return repo.SearchSchool(&SearchSchoolCriteria{CityID: city.CityID})
	}
	return nil, errors.New("null ISchoolRepository")
}

func (city City) GetStudios(repo IStudioRepository) ([]Studio, error) {
	if repo != nil {
		return repo.SearchStudio(&SearchStudioCriteria{CityID: city.CityID})
	}
	return nil, errors.New("null IStudioRepository")
}

// Parameters for search City
type SearchCityCriteria struct {
	CityID  int    `schema:"id"`
	Name    string `schema:"name"`
	StateID int    `schema:"state"`
}

type ICityRepository interface {
	CreateCity(city *City) error
	SearchCity(criteria *SearchCityCriteria) ([]City, error)
	UpdateCity(city City) error
	DeleteCity(city City) error
}
