// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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

// AccountTypeRepository specifies the functiosn that need to be implemented for looking up account types in DAS
type IAccountTypeRepository interface {
	GetAccountTypes() ([]AccountType, error)
}
