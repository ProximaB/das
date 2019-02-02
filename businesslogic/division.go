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

package businesslogic

import (
	"time"
)

type Division struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	FederationID    int
	Note            string
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

type SearchDivisionCriteria struct {
	ID           int    `schema:"id"`
	Name         string `schema:"name"`
	FederationID int    `schema:"federation"`
}

type IDivisionRepository interface {
	CreateDivision(division *Division) error
	SearchDivision(criteria SearchDivisionCriteria) ([]Division, error)
	UpdateDivision(division Division) error
	DeleteDivision(division Division) error
}
