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
	OrganizerID string `json:"uuid"`
	Name        string `json:"name"`
	Available   int    `json:"available"`
	Hosted      int    `json:"hosted"`
}

func (summary *OrganizerProvisionSummary) Summarize(provision businesslogic.OrganizerProvision) {
	summary.OrganizerID = provision.Organizer.UID
	summary.Available = provision.Available
	summary.Name = provision.Organizer.FullName()
	summary.Hosted = provision.Hosted
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
	Description       string    `json:"description"`
	Status            int       `json:"status"`
	DateTimeSubmitted time.Time `json:"created"`
	DateTimeResponded time.Time `json:"responded"`
}

// SearchRoleApplicationCriteria specifies the search criteria for role application
type SearchRoleApplicationCriteriaViewModel struct {
	ID             int  `schema:"id"`
	AccountID      int  `schema:"applicant"`
	AppliedRoleID  int  `schema:"appliedRole"`
	StatusID       int  `schema:"statusId"`
	ApprovalUserID int  `schema:"approvedBy"`
	Responded      bool `schema:"responded"`
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
	ApplicationID int `json:"applicationId"`
	Response      int `json:"responseId"`
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
