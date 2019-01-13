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

// Age contains data for the age category requirement for events
// Age is associated with Division, which is associated with Federation
type Age struct {
	ID              int
	Name            string
	Description     string
	DivisionID      int
	Enforced        bool // if required, AgeMinimum and AgeMaximum must have non-zero value
	AgeMinimum      int
	AgeMaximum      int
	CreateUserID    *int
	DateTimeCreated time.Time
	UpdateUserID    *int
	DateTimeUpdated time.Time
}

// SearchAgeCriteria provides parameters when searching Age in IAgeRepository
type SearchAgeCriteria struct {
	DivisionID int `schema:"division"`
	AgeID      int `schema:"id"`
}

// IAgeRepository provides an interface for other businesslogic code to access Age data
type IAgeRepository interface {
	CreateAge(age *Age) error
	SearchAge(criteria SearchAgeCriteria) ([]Age, error)
	UpdateAge(age Age) error
	DeleteAge(age Age) error
}
