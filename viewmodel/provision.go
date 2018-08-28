// Dancesport Application System (DAS)
// Copyright (C) 2017, 2018 Yubing Hou
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package viewmodel

import (
	"errors"
	"github.com/DancesportSoftware/das/businesslogic"
	"time"
)

// UpdateProvision specifies the payload that admin user need to submit to update organizer's provision
type UpdateProvision struct {
	OrganizerID     string `json:"organizer"`
	AmountAllocated int    `json:"allocate"`
	Note            string `json:"note"`
}

// OrganizerProvisionSummary specifies payload for Organizer Provision
type OrganizerProvisionSummary struct {
	OrganizerID int `json:"organizer"`
	Available   int `json:"available"`
	Hosted      int `json:"hosted"`
}

// SubmitRoleApplication is the payload for role application submission
type SubmitRoleApplication struct {
	AppliedRoleID int    `json:"role"`
	Description   string `json:"description"`
}

type RoleApplication struct {
	ID                int       `json:"id"`
	ApplicantName     string    `json:"applicant"`
	RoleApplied       int       `json:"role"`
	Description       string    `json:"description "`
	Status            int       `json:"status"`
	DateTimeSubmitted time.Time `json:"created"`
	DateTimeResponded time.Time `json:"responded"`
}

// Validate validates SubmitRoleApplication and check if data sanitized
func (dto SubmitRoleApplication) Validate() error {
	if dto.AppliedRoleID < businesslogic.AccountTypeAdjudicator || dto.AppliedRoleID > businesslogic.AccountTypeEmcee {
		return errors.New("this role does not exist")
	}
	if len(dto.Description) < 20 {
		return errors.New("insufficient description")
	}
	return nil
}

// RespondRoleApplication specifies the payload for responding a role application
type RespondRoleApplication struct {
	ApplicationID int `json:"application"`
	Response      int `json:"response"`
}

// Validate validates RespondRoleApplication and check if data sanitized
func (dto RespondRoleApplication) Validate() error {
	if dto.ApplicationID == 0 {
		return errors.New("application does not exist")
	}
	if !(dto.Response == businesslogic.RoleApplicationStatusApproved || dto.Response == businesslogic.RoleApplicationStatusDenied) {
		return errors.New("invalid response to the application")
	}
	return nil
}
