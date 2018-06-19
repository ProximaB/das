// Copyright 2017, 2018 Yubing Hou. All rights reserved.
// Use of this source code is governed by GPL license
// that can be found in the LICENSE file

package viewmodel

type UpdateProvision struct {
	OrganizerID     string `json:"organizer"`
	AmountAllocated int    `json:"allocate"`
	Note            string `json:"note"`
}
type OrganizerProvisionSummary struct {
	OrganizerID int `json:"organizer"`
	Available   int `json:"available"`
	Hosted      int `json:"hosted"`
}
