// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

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
