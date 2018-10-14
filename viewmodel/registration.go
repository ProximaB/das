package viewmodel

import "github.com/DancesportSoftware/das/businesslogic"

// SubmitCompetitionRegistrationForm is the payload that should be submitted by athlete to sign up for a competition
// This form should only contain events that the couple would compete. If an existing registration
type SubmitCompetitionRegistrationForm struct {
	PartnershipID     int   `json:"partnershipId"`
	CompetitionID     int   `json:"competitionId"`
	EventIds          []int `json:"events"`
	StateRepresented  int   `json:"repStateId"`
	SchoolRepresented int   `json:"repSchoolId"`
	StudioRepresented int   `json:"repStudioId"`
}

func (form SubmitCompetitionRegistrationForm) EventRegistration() businesslogic.EventRegistration {
	return businesslogic.EventRegistration{
		CompetitionID: form.CompetitionID,
		PartnershipID: form.PartnershipID,
	}
}
