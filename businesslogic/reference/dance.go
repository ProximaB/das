// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package referencebll

import (
	"time"
)

// Dance is style-dependent. For example, Cha Cha of American Rhythm is different from Cha Cha of International Latin
type Dance struct {
	ID              int
	Name            string
	Description     string
	Abbreviation    string
	StyleID         int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

// SearchDanceCriteria specifies the parameters that can be used to to search dances in DAS
type SearchDanceCriteria struct {
	StyleID int    `schema:"style"`
	DanceID int    `schema:"id"`
	Name    string `schema:"name"`
}

// IDanceRepository specifies the interface that needs to be implemented to functions as a repository for dance
type IDanceRepository interface {
	CreateDance(dance *Dance) error
	SearchDance(criteria SearchDanceCriteria) ([]Dance, error)
	UpdateDance(dance Dance) error
	DeleteDance(dance Dance) error
}

// ByDanceID allows sort a slice of dances by their IDs
type ByDanceID []Dance

func (d ByDanceID) Len() int           { return len(d) }
func (d ByDanceID) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d ByDanceID) Less(i, j int) bool { return d[i].ID < d[j].ID }
