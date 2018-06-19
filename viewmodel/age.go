// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package viewmodel

import (
	"github.com/DancesportSoftware/das/businesslogic/reference"
)

type SearchAge struct {
	DivisionID int `schema:"division"`
	AgeID      int `schema:"id"`
}

type Age struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Division int    `json:"division"`
	Enforced bool   `json:"enforced"`
	Minimum  int    `json:"minimum"`
	Maximum  int    `json:"maximum"`
}

func AgeDataModelToViewModel(dm referencebll.Age) Age {
	return Age{
		ID:       dm.ID,
		Name:     dm.Name,
		Division: dm.DivisionID,
		Enforced: dm.Enforced,
		Minimum:  dm.AgeMinimum,
		Maximum:  dm.AgeMaximum,
	}
}
