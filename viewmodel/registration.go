package viewmodel

import "github.com/DancesportSoftware/das/businesslogic"

// AthleteCompetitionRegistrationForm is the payload that should be submitted by athlete to sign up for a competition
// This form should only contain events that the couple would compete. If an existing registration
type AthleteCompetitionRegistrationForm struct {
	PartnershipID  int   `json:"partnershipId"`
	CompetitionID  int   `json:"competitionId"`
	AddedEvents    []int `json:"addedEvents"`
	DroppedEvents  []int `json:"droppedEvents"`
	Representation struct {
		CountryId int `json:"countryId,omitempty"`
		StateId   int `json:"stateId,omitempty"`
		SchoolId  int `json:"schoolId,omitempty"`
		StudioId  int `json:"studioId,omitempty"`
	} `json:"representation,omitempty"`
}

func (form AthleteCompetitionRegistrationForm) EventRegistration() businesslogic.EventRegistrationForm {
	return businesslogic.EventRegistrationForm{
		CompetitionID:      form.CompetitionID,
		PartnershipID:      form.PartnershipID,
		EventsAdded:        form.AddedEvents,
		EventsDropped:      form.DroppedEvents,
		CountryRepresented: form.Representation.CountryId,
		StateRepresented:   form.Representation.StateId,
		SchoolRepresented:  form.Representation.SchoolId,
		StudioRepresented:  form.Representation.StudioId,
	}
}
