// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import "time"

type RoundEntry struct {
	RoundID int
}

type PartnershipRoundEntry struct {
	ID                      int
	RoundID                 int
	PartnershipEventEntryID int
	DateTimeCreated         time.Time
	DateTimeUpdated         time.Time
}

type AdjudicatorRoundEntry struct {
	ID                      int
	RoundID                 int
	AdjudicatorEventEntryID int
	CreateUserID            int
	DateTimeCreated         time.Time
	UpdateUserID            int
	DateTimeUpdated         time.Time
}
