package viewmodel

type SearchEntryForm struct {
	CompetitionID int `schema:"competitionId"`
	EventID       int `schema:"eventId"`
	PartnershipID int `schema:"partnershipId,omitempty"`
	AthleteID     int `schema:"athleteId,omitempty"`
}
