// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import "time"

type PartnershipRequestBlacklistReason struct {
	ID              int
	Name            string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

type IPartnershipRequestBlacklistReasonRepository interface {
	GetPartnershipRequestBlacklistReasons() ([]PartnershipRequestBlacklistReason, error)
}
