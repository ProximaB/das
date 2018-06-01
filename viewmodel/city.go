package viewmodel

import (
	"github.com/yubing24/das/businesslogic/reference"
)

type City struct {
	CityID int    `json:"id"`
	Name   string `json:"name"`
	State  int    `json:"state"`
}

type CreateCity struct {
	Name    string `schema:"name"`
	StateID int    `schema:"state"`
}

func (create CreateCity) ToCityDataModel() reference.City {
	return reference.City{
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
