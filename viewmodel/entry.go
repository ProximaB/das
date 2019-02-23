package viewmodel

import "github.com/DancesportSoftware/das/businesslogic"

type SearchEntryForm struct {
	CompetitionID int `schema:"competitionId"`
	EventID       int `schema:"eventId"`
	FederationID  int `schema:"federationId"`
	DivisionID    int `schema:"divisionId"`
	AgeID         int `schema:"ageId"`
	ProficiencyID int `schema:"proficiencyId"`
	StyleID       int `schema:"styleId"`
	PartnershipID int `schema:"partnershipId,omitempty"`
	AthleteID     int `schema:"athleteId,omitempty"`
}

type AthleteCompetitionEntryViewModel struct {
	EntryID       int                  `json:"id"`
	CompetitionID int                  `json:"competitionId"`
	Athlete       AthleteTinyViewModel `json:"athlete"`
}

type CoupleCompetitionEntryViewModel struct {
	EntryID       int                      `json:"id"`
	CompetitionID int                      `json:"competitionId"`
	Couple        PartnershipTinyViewModel `json:"partnership"`
}

type AthleteEventEntryViewModel struct {
	EventID int                  `json:"eventId"`
	Athlete AthleteTinyViewModel `json:"athlete"`
}

type CoupleEventEntryViewModel struct {
	EventID  int                  `json:"eventId"`
	CoupleID int                  `json:"partnershipId"`
	Lead     AthleteTinyViewModel `json:"lead"`
	Follow   AthleteTinyViewModel `json:"follow"`
}

type CompetitionEntryListViewModel struct {
	Competition    CompetitionViewModel               `json:"competition"`
	AthleteEntries []AthleteCompetitionEntryViewModel `json:"athleteEntries"`
	CoupleEntries  []CoupleCompetitionEntryViewModel  `json:"partnershipEntries"`
}

func CompetitionEntriesToViewModel(entries businesslogic.CompetitionEntryList) CompetitionEntryListViewModel {
	view := CompetitionEntryListViewModel{}

	view.Competition = CompetitionDataModelToViewModel(entries.Competition, businesslogic.AccountTypeNoAuth)

	athletes := make([]AthleteCompetitionEntryViewModel, 0)
	for _, each := range entries.AthleteEntries {
		view := AthleteCompetitionEntryViewModel{
			EntryID:       each.ID,
			CompetitionID: each.Competition.ID,
			Athlete:       AthleteToTinyViewModel(each.Athlete),
		}
		athletes = append(athletes, view)
	}

	couples := make([]CoupleCompetitionEntryViewModel, 0)
	for _, each := range entries.CoupleEntries {
		view := CoupleCompetitionEntryViewModel{
			EntryID:       each.ID,
			CompetitionID: each.Competition.ID,
			Couple:        PartnershipToTinyViewModel(each.Couple),
		}
		couples = append(couples, view)
	}

	view.AthleteEntries = athletes
	view.CoupleEntries = couples
	return view
}

type EventEntryListViewModel struct {
	Event          EventViewModel               `json:"event"`
	AthleteEntries []AthleteEventEntryViewModel `json:"athleteEntries"`
	CoupleEntries  []CoupleEventEntryViewModel  `json:"partnershipEntries"`
}
