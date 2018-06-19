// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import "time"

const (
	AccountTypeUnauthorized  = -1
	ACCOUNT_TYPE_NOAUTH      = 0
	ACCOUNT_TYPE_ATHLETE     = 1
	AccountTypeAdjudicator   = 2
	AccountTypeScrutineer    = 3
	ACCOUNT_TYPE_ORGANIZER   = 4
	AccountTypeDeckCaptain   = 5
	AccountTypeEmcee         = 6
	AccountTypeAdministrator = 7
)

type AccountType struct {
	ID              int
	Name            string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

type IAccountTypeRepository interface {
	GetAccountTypes() ([]AccountType, error)
}
