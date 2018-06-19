// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package referencebll

import (
	"time"
)

type School struct {
	ID              int
	Name            string
	CityID          int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

type SearchSchoolCriteria struct {
	ID      int    `schema:"id"`
	Name    string `schema:"name"`
	CityID  int    `schema:"city"`
	StateID int    `schema:"state"`
}

type ISchoolRepository interface {
	CreateSchool(school *School) error
	SearchSchool(criteria SearchSchoolCriteria) ([]School, error)
	UpdateSchool(school School) error
	DeleteSchool(school School) error
}
