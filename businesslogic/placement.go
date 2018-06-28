// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import "time"

type AdjudicatorRoundPlacement struct {
	ID                      int
	AdjudicatorRoundEntryID int
	PartnershipRoundEntryID int
	RoundDanceID            int
	Placement               int
	CreateUserID            int
	DateTimeCreated         time.Time
	UpdateUserID            int
	DateTimeUpdated         time.Time
}
