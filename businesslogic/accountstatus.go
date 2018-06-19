// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import (
	"time"
)

const (
	ACCOUNT_STATUS_ACTIVATED  = 1
	ACCOUNT_STATUS_UNVERIFIED = 2
	ACCOUNT_STATUS_SUSPENDED  = 3
	ACCOUNT_STATUS_LOCKED     = 4
)

// IAccountStatusRepository specifies the requirements
type IAccountStatusRepository interface {
	GetAccountStatus() ([]AccountStatus, error)
}

type AccountStatus struct {
	ID              int
	Name            string
	Abbreviation    string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}
