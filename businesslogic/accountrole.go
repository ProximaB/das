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

import (
	"time"
)

// AccountRole defines the role that an account can be associated with
type AccountRole struct {
	ID              int
	AccountID       int
	AccountTypeID   int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

// SearchAccountRoleCriteria specifies the parameters that can be used to search Account Roles in a repository
type SearchAccountRoleCriteria struct {
	AccountID     int
	AccountTypeID int
}

// IAccountRoleRepository defines the functions that an account role repository should implement
type IAccountRoleRepository interface {
	CreateAccountRole(role *AccountRole) error
	SearchAccountRole(criteria SearchAccountRoleCriteria) ([]AccountRole, error)
}
