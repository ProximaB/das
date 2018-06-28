// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package viewmodel

import "github.com/DancesportSoftware/das/businesslogic"

type CompetitionStatus struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CompetitionStatusDataModelToViewModel(status businesslogic.CompetitionStatus) CompetitionStatus {
	return CompetitionStatus{
		ID:   status.ID,
		Name: status.Name,
	}
}
