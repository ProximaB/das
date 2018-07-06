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

const (
	// AccountStatusActivated sets a flag that account is activated and can function
	AccountStatusActivated = 1
	// AccountStatusUnverified is the status for most newly created accounts
	AccountStatusUnverified = 2
	// AccountStatusSuspended is the status for accounts that violates ToS or Privacy Policies
	AccountStatusSuspended = 3
	// AccountStatusLocked is the status for accounts that are locked due to security issues
	AccountStatusLocked = 4
)

// IAccountStatusRepository specifies the requirements
type IAccountStatusRepository interface {
	GetAccountStatus() ([]AccountStatus, error)
}

// AccountStatus defines the status that a DAS account could be. The status
// of an account can affect the authorization of some actions.
type AccountStatus struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}
