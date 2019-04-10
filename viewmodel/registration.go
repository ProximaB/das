package viewmodel

// AthleteCompetitionRegistrationForm is the payload that should be submitted by athlete to sign up for a competition
// This form should only contain events that the couple would compete. If an existing registration
type AthleteCompetitionRegistrationForm struct {
	CompetitionID  int   `json:"competitionId"`
	PartnershipID  int   `json:"partnershipId"`
	AddedEvents    []int `json:"addedEvents"`
	DroppedEvents  []int `json:"droppedEvents"`
	Representation struct {
		CountryId int `json:"countryId,omitempty"`
		StateId   int `json:"stateId,omitempty"`
		SchoolId  int `json:"schoolId,omitempty"`
		StudioId  int `json:"studioId,omitempty"`
	} `json:"representation,omitempty"`
}
