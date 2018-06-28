// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package viewmodel

import "github.com/DancesportSoftware/das/businesslogic/reference"

type Proficiency struct {
	ProficiencyID int    `json:"id"`
	Name          string `json:"name"`
	Division      int    `json:"division"`
}

func ProficiencyDataModelToViewModel(dm referencebll.Proficiency) Proficiency {
	return Proficiency{
		ProficiencyID: dm.ID,
		Name:          dm.Name,
		Division:      dm.DivisionID,
	}
}
