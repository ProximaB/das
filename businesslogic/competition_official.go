// Dancesport Application System (DAS)
// Copyright (C) 2019 Yubing Hou
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

type CompetitionOfficial struct {
	ID                int
	CompetitionID     int
	OfficialAccountID int       // the ID for AccountRole
	OfficialRoleID    int       // the ID for AccountType
	EffectiveFrom     time.Time // have privileged access to competition data
	EffectiveUntil    time.Time
	AssignedBy        int // ID of an AccountRole object, must be an organizer. TODO: may use invitation instead of assignment
	CreateUserID      int
	DateTimeCreated   time.Time
	UpdateUserID      int
	DateTimeUpdated   time.Time
}

type SearchCompetitionOfficialCriteria struct {
}

type ICompetitionOfficialRepository interface {
	CreateCompetitionOfficial(official *CompetitionOfficial) error
	DeleteCompetitionOfficial(official CompetitionOfficial) error
	SearchCompetitionOfficial(criteria SearchCompetitionOfficialCriteria) ([]CompetitionOfficial, error)
	UpdateCompetitionOfficial(official CompetitionOfficial) error
}
