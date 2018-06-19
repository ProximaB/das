// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package viewmodel

import (
	"github.com/DancesportSoftware/das/businesslogic/reference"
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

func (create CreateCity) ToCityDataModel() referencebll.City {
	return referencebll.City{
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
