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
