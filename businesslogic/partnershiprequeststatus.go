// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import "time"

type PartnershipRequestStatus struct {
	ID              int
	Code            string
	Description     string
	DateTimeCreated time.Time
	DateTimeUpdated time.Time
}

type IPartnershipRequestStatusRepository interface {
	GetPartnershipRequestStatus() ([]PartnershipRequestStatus, error)
}
