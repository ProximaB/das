// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package referencebll

import (
	"time"
)

type Style struct {
	ID              int
	Name            string
	Description     string
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

type SearchStyleCriteria struct {
	StyleID int    `schema:"id"`
	Name    string `schema:"name"`
}

type IStyleRepository interface {
	CreateStyle(style *Style) error
	SearchStyle(criteria SearchStyleCriteria) ([]Style, error)
	UpdateStyle(style Style) error
	DeleteStyle(style Style) error
}
