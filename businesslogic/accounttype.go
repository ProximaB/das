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
	// AccountTypeUnauthorized is the constant for failing authorization
	AccountTypeUnauthorized = -1
	// AccountTypeNoAuth is the constant for No Authentication is required
	AccountTypeNoAuth = 0
	// AccountTypeAthlete is the constant for Athlete-level authorization requirement
	AccountTypeAthlete = 1
	// AccountTypeAdjudicator is the constant for Adjudicator-level authorization requirement
	AccountTypeAdjudicator = 2
	// AccountTypeScrutineer is the constant for Scrutineer-level authorization requirement
	AccountTypeScrutineer = 3
	// AccountTypeOrganizer is the constant for Organizer-level authorization requirement
	AccountTypeOrganizer = 4
	// AccountTypeDeckCaptain is the constant for Deck Captain-level authorization requirement
	AccountTypeDeckCaptain = 5
	// AccountTypeEmcee is the constant for Emcee-level authorization requirement
	AccountTypeEmcee = 6
	// AccountTypeAdministrator is the constant for Administrator-level authorization requirement
	AccountTypeAdministrator = 7
)

// AccountType defines the data to specify an account type. There are seven account types in DAS
type AccountType struct {
	ID              int
	Name            string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

// IAccountTypeRepository specifies the functiosn that need to be implemented for looking up account types in DAS
type IAccountTypeRepository interface {
	GetAccountTypes() ([]AccountType, error)
}
