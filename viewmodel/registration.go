package viewmodel

import "github.com/DancesportSoftware/das/businesslogic"

// AthleteCompetitionRegistrationForm is the payload that should be submitted by athlete to sign up for a competition
// This form should only contain events that the couple would compete. If an existing registration
type AthleteCompetitionRegistrationForm struct {
	PartnershipID      int   `json:"partnershipId"`
	CompetitionID      int   `json:"competitionId"`
	AddedEvents        []int `json:"addedEvents"`
	DroppedEvents      []int `json:"droppedEvents"`
	CountryRepresented int   `json:"countryId,omitempty"`
	StateRepresented   int   `json:"stateId,omitempty"`
	SchoolRepresented  int   `json:"schoolId,omitempty"`
	StudioRepresented  int   `json:"studioId,omitempty"`
}

func (form AthleteCompetitionRegistrationForm) EventRegistration() businesslogic.EventRegistrationForm {
	return businesslogic.EventRegistrationForm{
		CompetitionID:      form.CompetitionID,
		PartnershipID:      form.PartnershipID,
		EventsAdded:        form.AddedEvents,
		EventsDropped:      form.DroppedEvents,
		CountryRepresented: form.CountryRepresented,
		StateRepresented:   form.StateRepresented,
		SchoolRepresented:  form.SchoolRepresented,
		StudioRepresented:  form.StudioRepresented,
	}
}
