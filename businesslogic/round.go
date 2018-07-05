// Dancesport Application System (DAS)
// Copyright (C) 2018 Yubing Hou
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

// RoundOrder defines the order of the round, from lowest to the highest
type RoundOrder struct {
	ID              int
	Rank            int
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

// Round defines the round for each event
type Round struct {
	ID              int
	EventID         int
	Order           RoundOrder
	Entries         []EventEntry
	StartTime       time.Time
	EndTime         time.Time
	DateTimeCreated time.Time
	CreateUserID    int
	DateTimeUpdated time.Time
	UpdateUserID    int
}

// SearchRoundCriteria specifies the parameters that can be used to search Rounds in a Repository
type SearchRoundCriteria struct {
	CompetitionID int
	EventID       int
	RoundOrderID  int
}

// IRoundRepository specifies the interface that a Round Repository should implement
type IRoundRepository interface {
	CreateRound(round *Round) error
	DeleteRound(round Round) error
	SearchRound(criteria SearchRoundCriteria) ([]Round, error)
	UpdateRound(round Round) error
}
