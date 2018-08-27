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

package viewmodel

import "github.com/DancesportSoftware/das/businesslogic/reference"

type Proficiency struct {
	ProficiencyID int    `json:"id"`
	Name          string `json:"name"`
	Division      int    `json:"division"`
}

func ProficiencyDataModelToViewModel(dm reference.Proficiency) Proficiency {
	return Proficiency{
		ProficiencyID: dm.ID,
		Name:          dm.Name,
		Division:      dm.DivisionID,
	}
}
