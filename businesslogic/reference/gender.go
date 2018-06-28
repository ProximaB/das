// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package referencebll

import (
	"time"
)

const (
	GENDER_MALE   = 2
	GENDER_FEMALE = 1
)

type IGenderRepository interface {
	GetAllGenders() ([]Gender, error)
}

type Gender struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}
