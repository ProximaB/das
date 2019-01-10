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

import (
	"github.com/DancesportSoftware/das/businesslogic"
)

type Studio struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	CityID  int    `json:"city"`
	Website string `json:"website"`
}

func StudioDataModelToViewModel(dm businesslogic.Studio) Studio {
	return Studio{
		ID:      dm.ID,
		Name:    dm.Name,
		Address: dm.Address,
		CityID:  dm.CityID,
		Website: dm.Website,
	}
}
