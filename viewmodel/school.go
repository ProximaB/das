// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package viewmodel

import "github.com/DancesportSoftware/das/businesslogic/reference"

type School struct {
	SchoolID int    `json:"id"`
	Name     string `json:"name"`
	CityID   int    `json:"city"`
}

func SchoolDataModelToViewModel(school referencebll.School) School {
	return School{
		SchoolID: school.ID,
		Name:     school.Name,
		CityID:   school.CityID,
	}
}
