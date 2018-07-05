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

// Placement defines the minimal data for a dance placement: Round + Dance + Adjudicator + Partnership + Placement together
// defines a unique placement
type Placement struct {
	ID                        int
	AdjudicatorRoundEntryID   int
	PartnershipRoundEntryID   int
	RoundDanceID              int
	PreliminaryRoundIndicator bool
	Placement                 int
	CreateUserID              int
	DateTimeCreated           time.Time
	UpdateUserID              int
	DateTimeUpdated           time.Time
}

// SearchPlacementCriteria specifies the parameters that can be used to search Placement in a repository
type SearchPlacementCriteria struct {
	CompetitionID int
	EventID       int
	PartnershipID int
}

// IPlacementRepository specifies the functions that a Placement Repository should implement
type IPlacementRepository interface {
	CreatePlacement(placement *Placement) error
	DeletePlacement(placement Placement) error
	SearchPlacement(criteria SearchPlacementCriteria) ([]Placement, error)
	UpdatePlacement(placement Placement) error
}
