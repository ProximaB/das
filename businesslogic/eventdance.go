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

import (
	"time"
)

// EventDance represents the many-to-many relationship between competition eventdal and dance references.
type EventDance struct {
	ID              int
	EventID         int
	DanceID         int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type SearchEventDanceCriteria struct {
	EventDanceID  int
	CompetitionID int
	EventID       int
}

type IEventDanceRepository interface {
	SearchEventDance(criteria SearchEventDanceCriteria) ([]EventDance, error)
	CreateEventDance(eventDance *EventDance) error
	DeleteEventDance(eventDance EventDance) error
	UpdateEventDance(eventDance EventDance) error
}

func NewEventDance(event Event, danceID int) *EventDance {
	return &EventDance{
		EventID:         event.ID,
		DanceID:         danceID,
		CreateUserID:    event.CreateUserID,
		DateTimeCreated: time.Now(),
		UpdateUserID:    event.UpdateUserID,
		DateTimeUpdated: time.Now(),
	}
}
