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
