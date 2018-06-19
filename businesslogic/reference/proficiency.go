// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package referencebll

import (
	"time"
)

type Proficiency struct {
	ID              int
	Name            string
	Description     string
	DivisionID      int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTImeUpdated time.Time
}

type SearchProficiencyCriteria struct {
	ProficiencyID int `schema:"id"`
	DivisionID    int `schema:"division"`
}

type IProficiencyRepository interface {
	SearchProficiency(criteria SearchProficiencyCriteria) ([]Proficiency, error)
	CreateProficiency(proficiency *Proficiency) error
	UpdateProficiency(proficiency Proficiency) error
	DeleteProficiency(proficiency Proficiency) error
}
