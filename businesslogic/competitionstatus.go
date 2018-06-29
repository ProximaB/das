// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import "time"

const (
	CompetitionStatusPreRegistration    = 1
	CompetitionStatusOpenRegistration   = 2
	CompetitionStatusClosedRegistration = 3
	CompetitionStatusInProgress         = 4
	CompetitionStatusProcessing         = 5
	CompetitionStatusClosed             = 6
	CompetitionStatusCancelled          = 7
)

// CompetitionStatus defines the data that is required to label the status of a Competition
type CompetitionStatus struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

// ICompetitionStatusRepository defines the function that a CompetitionStatusRepository should implement
type ICompetitionStatusRepository interface {
	GetCompetitionAllStatus() ([]CompetitionStatus, error)
}
