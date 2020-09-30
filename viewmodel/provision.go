package viewmodel

import (
	"github.com/ProximaB/das/businesslogic"
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
	RoleID      int    `json:"roleId" validate:"min=2,max=6"` // CAUTION! hard coded role ID here!
	Description string `json:"description" validate:"min=20"`
}

type RoleApplicationAdminView struct {
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

// RespondRoleApplication specifies the payload for responding a role application
type RespondRoleApplication struct {
	ApplicationID int `json:"applicationId" validate:"min=1"`
	Response      int `json:"responseId" validate:"min=1,max=2"`
}
