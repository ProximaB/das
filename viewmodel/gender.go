// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package viewmodel

import (
	"github.com/DancesportSoftware/das/businesslogic/reference"
)

type Gender struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GenderDataModelToViewModel(gender referencebll.Gender) Gender {
	return Gender{
		ID:   gender.ID,
		Name: gender.Name,
	}
}
