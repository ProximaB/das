// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package referencebll

import (
	"time"
)

type Studio struct {
	ID              int
	Name            string
	Address         string
	CityID          int
	Website         string
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

type SearchStudioCriteria struct {
	ID        int    `schema:"id"`
	Name      string `schema:"name"`
	CityID    int    `schema:"city"`
	StateID   int    `schema:"state"`
	CountryID int    `schema:"country"`
}

type IStudioRepository interface {
	CreateStudio(studio *Studio) error
	SearchStudio(criteria SearchStudioCriteria) ([]Studio, error)
	DeleteStudio(studio Studio) error
	UpdateStudio(studio Studio) error
}
