// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import "time"

const (
	NewPartnershipRequestNotification = 1
)

// NotificationPreference stores the preference of how user would like to receive system-generated notification
type NotificationPreference struct {
}

type Notification struct {
	ID              int
	HasRead         bool
	DateTimeCreated time.Time
}
