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
