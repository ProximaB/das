// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package businesslogic

import "time"

type RoundEntry struct {
	RoundID         int
	CreateUserID    int
	DateTimeCreated time.Time
	UpdateUserID    int
	DateTimeUpdated time.Time
}

type PartnershipRoundEntry struct {
	ID                      int
	RoundEntry              RoundEntry
	PartnershipEventEntryID int
}

type SearchPartnershipRoundEntryCriteria struct {
}

type AdjudicatorRoundEntry struct {
	ID                 int
	AdjudicatorEntryID int
	RoundEntry         RoundEntry
}

type SearchAdjudicatorRoundEntryCriteria struct {
}
