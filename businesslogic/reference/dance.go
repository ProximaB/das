// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package reference

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
